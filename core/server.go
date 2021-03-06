package core

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/craiggwilson/mongo-go-driver/core/util"
)

// Server represents a server.
type Server interface {
	// Connection gets a connection to the server.
	Connection() (Connection, error)
}

// ServerDesc is a description of a server.
type ServerDesc struct {
	endpoint Endpoint

	averageRTT         time.Duration
	electionID         bson.ObjectId
	lastError          error
	lastWriteTimestamp time.Time
	maxBatchCount      uint16
	maxDocumentSize    uint32
	maxMessageSize     uint32
	canonicalEndpoint  Endpoint
	members            []Endpoint
	serverType         ServerType
	setName            string
	setVersion         uint32
	tags               []bson.D
	wireVersion        util.Range
	version            util.Version
}

// AverageRTT is the average round trip time for the server.
func (d *ServerDesc) AverageRTT() time.Duration {
	return d.averageRTT
}

// Endpoint is the endpoint of the server.
func (d *ServerDesc) Endpoint() Endpoint {
	return d.endpoint
}

// LastError holds the last error that occured during a heartbeat
// which causes this server to be unusable. Will be nil when the
// server is available.
func (d *ServerDesc) LastError() error {
	return d.lastError
}

// Type is the type of the server.
func (d *ServerDesc) Type() ServerType {
	return d.serverType
}

// Tags are the tags used to select this server during server selection.
func (d *ServerDesc) Tags() []bson.D {
	// TODO: make a copy of the slice?
	return d.tags
}

// Version is the version of the server.
func (d *ServerDesc) Version() util.Version {
	return d.version
}

// ServerType represents a type of server.
type ServerType uint32

// ServerType constants.
const (
	UnknownServerType ServerType = 0
	Standalone        ServerType = 1
	RSMember          ServerType = 2
	RSPrimary         ServerType = 4 + RSMember
	RSSecondary       ServerType = 8 + RSMember
	RSArbiter         ServerType = 16 + RSMember
	RSGhost           ServerType = 32 + RSMember
	Mongos            ServerType = 256
)
