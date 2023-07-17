package configs

import "fmt"

type GRPC struct {
	IP   string `json:"ip" yaml:"ip"`
	Port int    `json:"port" yaml:"port"`
}

func (g GRPC) Addr() string {
	return fmt.Sprintf("%s:%d", g.IP, g.Port)
}
