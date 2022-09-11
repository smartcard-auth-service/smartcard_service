package config

var (
	Cfg *GlobalConfig
)

type GlobalConfig struct {
	MONGO_URI              string `json:"MONGO_URI" default:"mongodb://localhost:27017"`
	GRPC_LISTEN_PORT       string `json:"GRPC_LISTEN_PORT" default:"8082"`
	TLS_CLIENT_CRT         string `json:"TLS_CLIENT_CRT" default:"/home/dmitry/go/smartcard_service/pkg/tls/tls_client/client.crt"`
	TLS_CLIENT_KEY         string `json:"TLS_CLIENT_KEY" default:"/home/dmitry/go/smartcard_service/pkg/tls/tls_client/client.key"`
	TLS_SERVER_CRT         string `json:"TLS_SERVER_CRT" default:"/home/dmitry/go/smartcard_service/pkg/tls/tls_server/server.crt"`
	TLS_SERVER_KEY         string `json:"TLS_SERVER_KEY" default:"/home/dmitry/go/smartcard_service/pkg/tls/tls_server/server.key"`
	TLS_URL                string `json:"TLS_URL" default:"https://localhost:1443"`
	TLS_SERVER_LISTEN_PORT string `json:"TLS_SERVER_LISTEN_PORT" default:"1443"`
}

func InitGlobalConfig() {
	Cfg = &GlobalConfig{
		MONGO_URI:        "mongodb://localhost:27017",
		GRPC_LISTEN_PORT: ":8082",
	}
}
