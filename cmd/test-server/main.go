package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/thisisdevelopment/helium-route-updater/pkg/api/helium/service/iot_config"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

type TestServer struct {
	iot_config.UnimplementedRouteServer
	Oui          uint64
	KeyPair      *helium_crypto.KeyPair
	Routes       map[string]*iot_config.RouteV1
	RouteDevices map[string]map[string]*iot_config.EuiPairV1
}

func (t *TestServer) List(ctx context.Context, req *iot_config.RouteListReqV1) (*iot_config.RouteListResV1, error) {
	if req.Oui != t.Oui {
		return nil, fmt.Errorf("invalid oui")
	}

	var routes []*iot_config.RouteV1
	for _, route := range t.Routes {
		routes = append(routes, route)
	}
	return &iot_config.RouteListResV1{
		Routes: routes,
	}, nil
}

// Get Route for an Org (auth delegate_keys/owner/admin)
func (t *TestServer) Get(ctx context.Context, req *iot_config.RouteGetReqV1) (*iot_config.RouteV1, error) {
	res, found := t.Routes[req.Id]
	if !found {
		return nil, fmt.Errorf("route not found")
	}
	return res, nil
}

// Create Route for an Org (auth delegate_keys/owner/admin)
func (t *TestServer) Create(ctx context.Context, req *iot_config.RouteCreateReqV1) (*iot_config.RouteV1, error) {
	if req.Oui != t.Oui {
		return nil, fmt.Errorf("invalid oui")
	}
	id := uuid.New()
	route := req.Route
	route.Oui = t.Oui
	route.Id = id.String()

	t.Routes[route.Id] = route
	return route, nil
}

// Update Route for an Org (auth delegate_keys/owner/admin)
func (t *TestServer) Update(ctx context.Context, req *iot_config.RouteUpdateReqV1) (*iot_config.RouteV1, error) {
	route, found := t.Routes[req.Route.Id]
	if !found {
		return nil, fmt.Errorf("route not found")
	}
	if req.Route.NetId > 0 {
		route.NetId = req.Route.NetId
	}
	if req.Route.Server != nil {
		route.Server = req.Route.Server
	}
	if req.Route.MaxCopies > 0 {
		route.MaxCopies = req.Route.MaxCopies
	}
	route.Active = req.Route.Active
	route.Locked = req.Route.Locked

	return route, nil
}

// Delete Route for an Org (auth delegate_keys/owner/admin)
func (t *TestServer) Delete(ctx context.Context, req *iot_config.RouteDeleteReqV1) (*iot_config.RouteV1, error) {
	res, found := t.Routes[req.Id]
	if !found {
		return nil, fmt.Errorf("route not found")
	}
	delete(t.Routes, req.Id)
	delete(t.RouteDevices, req.Id)
	return res, nil
}

func (t *TestServer) GetEuis(req *iot_config.RouteGetEuisReqV1, srv iot_config.Route_GetEuisServer) error {
	_, found := t.Routes[req.RouteId]
	if !found {
		return fmt.Errorf("route not found")
	}

	for _, dev := range t.RouteDevices[req.RouteId] {
		err := srv.Send(dev)
		if err != nil {
			return err
		}
	}

	return nil
}
func (t *TestServer) UpdateEuis(srv iot_config.Route_UpdateEuisServer) error {
	ctx := srv.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		if !helium_crypto.VerifyRequest(req, t.KeyPair) {
			log.Println("Invalid signature!!")
			return fmt.Errorf("invalid signature")
		}

		_, found := t.Routes[req.EuiPair.RouteId]
		if !found {
			return fmt.Errorf("route not found")
		}

		devEui := strconv.FormatUint(req.EuiPair.DevEui, 16)
		_, exists := t.RouteDevices[req.EuiPair.RouteId][devEui]

		switch req.Action {
		case iot_config.ActionV1_add:
			if exists {
				return fmt.Errorf("route already exists for deveui: %s", devEui)
			}
			fmt.Printf("Added euipair: %+v\n", req.EuiPair)
			t.RouteDevices[req.EuiPair.RouteId][devEui] = req.EuiPair
			break
		case iot_config.ActionV1_remove:
			if !exists {
				return fmt.Errorf("route does not exist for deveui: %s", devEui)
			}
			fmt.Printf("removed euipair: %+v\n", req.EuiPair)
			delete(t.RouteDevices[req.EuiPair.RouteId], devEui)
			break
		default:
			panic("not implemented")
		}
	}

	return srv.SendAndClose(&iot_config.RouteEuisResV1{})
}

// Delete all EUIs for a Route (auth delegate_keys/owner/admin)
func (t *TestServer) DeleteEuis(ctx context.Context, req *iot_config.RouteDeleteEuisReqV1) (*iot_config.RouteEuisResV1, error) {
	_, found := t.Routes[req.RouteId]
	if !found {
		return nil, fmt.Errorf("route not found")
	}
	delete(t.RouteDevices, req.RouteId)

	return &iot_config.RouteEuisResV1{}, nil
}

func main() {
	//TODO: params
	//TODO: auth

	listen := ":8000"
	var oui uint64 = 0
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on:", listen)

	testServer := &TestServer{
		Oui:          oui,
		KeyPair:      helium_crypto.KeyPairFromString("13ooSU22gnnV6DPEqp8dGzXfmVc4r6AeK7JZWb6fs934cTAu7fu"),
		Routes:       make(map[string]*iot_config.RouteV1),
		RouteDevices: make(map[string]map[string]*iot_config.EuiPairV1),
	}
	testServer.Routes["test"] = &iot_config.RouteV1{
		Id:        "test",
		NetId:     53,
		Oui:       oui,
		Server:    nil,
		MaxCopies: 5,
		Active:    true,
		Locked:    false,
	}

	testServer.RouteDevices["test"] = make(map[string]*iot_config.EuiPairV1)

	grpcServer.RegisterService(&iot_config.Route_ServiceDesc, testServer)
	grpcServer.Serve(lis)
}
