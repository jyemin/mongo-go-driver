package core

import "time"
import "github.com/craiggwilson/mongo-go-driver/core/msg"

var defaultCodec = msg.NewWireProtocolCodec

var defaultEndpoint = Endpoint("localhost:27017")

var defaultEndpointDialer = DialEndpoint

var defaultHeartbeatInterval = time.Duration(10) * time.Second

var defaultServerOptionsFactory = func(endpoint Endpoint) ServerOptions {
	return ServerOptions{
		ConnectionOptions: ConnectionOptions{
			Endpoint: endpoint,
		},
		HeartbeatInterval: defaultHeartbeatInterval,
	}
}
