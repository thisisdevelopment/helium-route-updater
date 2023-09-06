package lns

import (
	"context"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	chirpstack "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/chirpstack/chirpstack/api/go/v4/common"
	"github.com/chirpstack/chirpstack/api/go/v4/meta"
	"github.com/redis/go-redis/v9"
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"log"
	"strconv"
	"strings"
)

type APIToken string

func (a APIToken) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", a),
	}, nil
}

func (a APIToken) RequireTransportSecurity() bool {
	return false
}

type ChirpstackClient struct {
	BaseClient
	deviceClient        chirpstack.DeviceServiceClient
	deviceProfileClient chirpstack.DeviceProfileServiceClient
	appClient           chirpstack.ApplicationServiceClient
	tenantClient        chirpstack.TenantServiceClient
	tenantId            string
}

func (c *ChirpstackClient) parseDevEui(deviceId string) ([]byte, bool) {
	decoded, err := hex.DecodeString(deviceId)
	if err != nil || len(decoded) != 8 {
		return nil, false
	}
	return decoded, true
}

func (c *ChirpstackClient) Listen(ch chan<- types.DeviceEvent, syncCh chan<- bool) {
	opt, err := redis.ParseURL(c.config.Listen)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	ctx := context.Background()

	streamMetaKey := "stream:meta"
	streamRequestKey := "api:stream:request"
	streamGwFrameKey := "gw:stream:frame"
	lastMetaId := "$"
	lastRequestId := "$"
	lastGwFrameId := "$"

	for {
		resp, err := rdb.XRead(ctx, &redis.XReadArgs{
			Streams: []string{streamMetaKey, streamRequestKey, streamGwFrameKey, lastMetaId, lastRequestId, lastGwFrameId},
			Count:   10,
			Block:   0,
		}).Result()
		if err != nil {
			log.Fatal(err)
		}

		for _, streamResp := range resp {
			for _, msg := range streamResp.Messages {
				if streamResp.Stream == streamRequestKey {
					lastRequestId = msg.ID
					if b, ok := msg.Values["request"].(string); ok {
						var pl chirpstack.RequestLog
						if err := proto.Unmarshal([]byte(b), &pl); err != nil {
							log.Fatal(err)
						}

						if pl.Service == "api.DeviceService" {
							if pl.Method != "Delete" && pl.Method != "Create" && pl.Method != "Update" {
								continue
							}

							if _, ok := c.parseDevEui(pl.Metadata["dev_eui"]); !ok {

								fmt.Printf("[lns][warn][%s] unable to process request for invalid device: %s\n", msg.ID, pl.Metadata["dev_eui"])
								continue
							}

							if pl.Method == "Delete" {
								//TODO: as we don't have the join eui here, we cannot remove it from our helium route
								//      instead we schedule a full sync
								fmt.Printf("[lns][%s] device deleted: %s\n", msg.ID, pl.Metadata["dev_eui"])
								syncCh <- true
							} else if pl.Method == "Create" {
								fmt.Printf("[lns][%s] device created %s\n", msg.ID, pl.Metadata["dev_eui"])
								ch <- types.DeviceEvent{
									Update: []*types.Device{c.GetDevice(pl.Metadata["dev_eui"])},
									Delete: []*types.Device{},
								}
							} else if pl.Method == "Update" {
								fmt.Printf("[lns][%s] device updated %s\n", msg.ID, pl.Metadata["dev_eui"])
								ch <- types.DeviceEvent{
									Update: []*types.Device{c.GetDevice(pl.Metadata["dev_eui"])},
									Delete: []*types.Device{},
								}
							}
						}
					}
				} else if streamResp.Stream == streamMetaKey {
					lastMetaId = msg.ID
					if b, ok := msg.Values["down"].(string); ok {
						var pl meta.DownlinkMeta
						if err := proto.Unmarshal([]byte(b), &pl); err != nil {
							log.Fatal(err)
						}

						if pl.MessageType == common.MType_JOIN_ACCEPT {
							fmt.Printf("[lns][%s] device joined %s\n", msg.ID, pl.DevEui)
							ch <- types.DeviceEvent{
								Update: []*types.Device{c.GetDevice(pl.DevEui)},
								Delete: []*types.Device{},
							}
						}
					}
				} else if streamResp.Stream == streamGwFrameKey {
					lastGwFrameId = msg.ID
					if b, ok := msg.Values["up"].(string); ok {
						if c.config.AutoRoaming {
							var pl chirpstack.UplinkFrameLog
							if err := proto.Unmarshal([]byte(b), &pl); err != nil {
								log.Fatal(err)
							}
							if pl.MType == common.MType_JOIN_REQUEST || pl.MType == common.MType_REJOIN_REQUEST {
								region, ok := pl.RxInfo[0].Metadata["region_common_name"]
								if ok {
									regionId, ok := common.Region_value[region]
									if ok {
										c.AutoRoaming(pl.DevEui, common.Region(regionId))
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (c *ChirpstackClient) GetApplicationIds() []string {
	offset := 0
	limit := 100
	count := limit
	res := []string{}

	for offset < count {
		apps, err := c.appClient.List(context.Background(), &chirpstack.ListApplicationsRequest{
			Limit:    uint32(limit),
			Offset:   uint32(offset),
			TenantId: c.tenantId,
		})
		if err != nil {
			panic(err)
		}
		for _, app := range apps.Result {
			res = append(res, app.Id)
		}
		count = int(apps.TotalCount)
		offset = offset + limit
	}

	return res
}

func (c *ChirpstackClient) GetDevice(deviceId string) *types.Device {
	d, err := c.deviceClient.Get(context.Background(), &chirpstack.GetDeviceRequest{DevEui: deviceId})
	if err != nil {
		panic(err)
	}

	activation, err := c.deviceClient.GetActivation(context.Background(), &chirpstack.GetDeviceActivationRequest{DevEui: deviceId})
	if err != nil {
		panic(err)
	}

	devEui, _ := strconv.ParseUint(d.Device.DevEui, 16, 64)
	joinEui, _ := strconv.ParseUint(d.Device.JoinEui, 16, 64)
	devAddr := uint64(0)
	sKey := []byte{}
	if activation.DeviceActivation != nil {
		devAddr, _ = strconv.ParseUint(activation.DeviceActivation.DevAddr, 16, 32)
		sKey, _ = hex.DecodeString(activation.DeviceActivation.NwkSEncKey)
	}

	maxCopies, err := strconv.ParseUint(d.Device.Variables["max_copies"], 10, 32)
	if err != nil {
		maxCopies = 0
	} else if maxCopies > 15 {
		maxCopies = 15
	}

	return &types.Device{
		DevEui:     devEui,
		JoinEui:    joinEui,
		DevAddr:    uint32(devAddr),
		SessionKey: sKey,
		MaxCopies:  uint32(maxCopies),
	}
}

func (c *ChirpstackClient) AutoRoaming(deviceId string, region common.Region) {
	d, err := c.deviceClient.Get(context.Background(), &chirpstack.GetDeviceRequest{DevEui: deviceId})
	if err != nil {
		panic(err)
	}

	profile, err := c.deviceProfileClient.Get(context.Background(), &chirpstack.GetDeviceProfileRequest{Id: d.Device.DeviceProfileId})
	if err != nil {
		panic(err)
	}

	if profile.DeviceProfile.Region != region {
		fmt.Printf("[lns] autoroaming %s; updating region from %s to %s\n", deviceId, profile.DeviceProfile.Region.String(), region.String())
		profile.DeviceProfile.Region = region
		_, err := c.deviceProfileClient.Update(context.Background(), &chirpstack.UpdateDeviceProfileRequest{DeviceProfile: profile.DeviceProfile})
		if err != nil {
			panic(err)
		}
	}
}

func (c *ChirpstackClient) GetDevices() []*types.Device {
	devices := []*types.Device{}
	for _, appId := range c.GetApplicationIds() {
		offset := 0
		limit := 100
		count := limit
		for offset < count {
			resp, err := c.deviceClient.List(context.Background(), &chirpstack.ListDevicesRequest{
				Limit:         uint32(limit),
				Offset:        uint32(offset),
				ApplicationId: appId,
			})
			if err != nil {
				panic(err)
			}

			for _, dev := range resp.Result {
				devices = append(devices, c.GetDevice(dev.DevEui))
			}

			count = int(resp.TotalCount)
			offset = offset + limit
		}
	}
	return devices
}

func NewChirpstackClient(client BaseClient) *ChirpstackClient {
	authParts := strings.Split(client.config.ApiAuth, ":")
	server := client.config.ApiEndpoint
	serverCredentials := credentials.NewTLS(&tls.Config{})
	if strings.HasPrefix(server, "http://") {
		server = server[7:]
		serverCredentials = insecure.NewCredentials()
	} else if strings.HasPrefix(server, "https://") {
		server = server[8:]
	}

	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(APIToken(authParts[1])),
		grpc.WithTransportCredentials(serverCredentials),
	}

	conn, err := grpc.Dial(server, dialOpts...)
	if err != nil {
		panic(err)
	}

	deviceClient := chirpstack.NewDeviceServiceClient(conn)
	deviceProfileClient := chirpstack.NewDeviceProfileServiceClient(conn)
	appClient := chirpstack.NewApplicationServiceClient(conn)
	tenantClient := chirpstack.NewTenantServiceClient(conn)

	return &ChirpstackClient{BaseClient: client, deviceClient: deviceClient, deviceProfileClient: deviceProfileClient, appClient: appClient, tenantClient: tenantClient, tenantId: authParts[0]}
}
