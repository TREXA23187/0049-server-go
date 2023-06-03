package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User bool
}

// Parse Parsing Command Parameters
func Parse() Option {
	db := sys_flag.Bool("db", false, "Initialize Database")
	user := sys_flag.Bool("user", false, "Init Admin")

	// Parse command line parameters and write them into the registered flag
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWebStop whether stop the web project after executing commands
func IsWebStop(option Option) (res bool) {
	maps := structs.Map(&option)
	for _, value := range maps {
		switch v := value.(type) {
		case string:
			if v != "" {
				res = true
			}
		case bool:
			if v == true {
				res = true
			}
		}
	}

	return res
}

// SwitchOption Execute different functions according to commands
func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
		return
	}

	if option.User {
		CreateUser()
		return
	}

	sys_flag.Usage()
}
