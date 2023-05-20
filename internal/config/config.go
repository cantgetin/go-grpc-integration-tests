package config

type Config struct {
	AppName  string `env:"APP_NAME" envDefault:"test-app"`
	PgDSN    string `env:"PG_DSN" envDefault:"postgresql://localhost:5468/db?user=db&password=db"`
	GRPCAddr string `env:"GRPC_ADDR" envDefault:":9017"`
}
