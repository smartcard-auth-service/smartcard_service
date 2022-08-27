package config

var (
	Cfg *GlobalConfig
)

type GlobalConfig struct {
	MONGO_URI        string `json:"MONGO_URI" default:"mongodb://localhost:27017"`
	GRPC_LISTEN_PORT string `json:"GRPC_LISTEN_PORT" default:"8082"`
}
