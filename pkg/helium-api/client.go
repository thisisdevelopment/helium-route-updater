package helium_api

import (
	"context"
	"crypto/tls"
	"github.com/thisisdevelopment/helium-route-updater/pkg/api/helium/service/iot_config"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strings"
	"time"
)

type Client struct {
	Keypair *helium_crypto.KeyPair
	conn    *grpc.ClientConn
}

func NewClient(server string, keypair *helium_crypto.KeyPair) *Client {
	serverCredentials := credentials.NewTLS(&tls.Config{})
	if strings.HasPrefix(server, "http://") {
		server = server[7:]
		serverCredentials = insecure.NewCredentials()
	} else if strings.HasPrefix(server, "https://") {
		server = server[8:]
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		server,
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithTransportCredentials(serverCredentials),
	)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return &Client{
		Keypair: keypair,
		conn:    conn,
	}
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) NewRouteClient() iot_config.RouteClient {
	return iot_config.NewRouteClient(c.conn)
}

func (c *Client) NewGatewayClient() iot_config.GatewayClient {
	return iot_config.NewGatewayClient(c.conn)
}

func (c *Client) NewOrgClient() iot_config.OrgClient {
	return iot_config.NewOrgClient(c.conn)
}

func (c *Client) NewAdminClient() iot_config.AdminClient {
	return iot_config.NewAdminClient(c.conn)
}
