package env

import "os"

const (
	HttpPrefix = "AOC2024_HTTP_PREFIX"
	HttpPort   = "AOC2024_HTTP_PORT"
	GrpcPort   = "AOC2024_GRPC_PORT"
)

func GetString(
	key string,
) string {
	return os.Getenv(key)
}

func GetStringOrDefault(
	key string,
	def string,
) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
