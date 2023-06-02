package configs

import "strconv"

type Mysql struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Config   string `yaml:"config"` // Advanced configurations such as: charset
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}

// Dsn is the function to get the dsn of the mysql database.
func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
}
