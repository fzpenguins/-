package constants

import "time"

//rpc
const (
	MuxConnection  = 10
	RPCTimeout     = 30 * time.Second
	ConnectTimeout = 500 * time.Millisecond
)

const (
	MaxIdleConns    = 10
	MaxConnections  = 1000
	ConnMaxLifetime = 10 * time.Second
)

//milvus
const (
	CollectionName = "Images"
)
