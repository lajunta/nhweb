package main

import (
	"log"
	"net/http"
	"os/user"
	"path"

	"github.com/BurntSushi/toml"
)

// Configuration means app config
type Configuration struct {
	Port        string
	Pass        string
	SessionKey  string
	CommandPath string
	IPS         []string
}

var (
	config = Configuration{}
)

func init() {
	user, _ := user.Current()
	configPath := path.Join(user.HomeDir, ".nhweb", "config.toml")
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		log.Println("Write Config File with following fields:")
		log.Println("Port Pass CommandPath SessionKey IPS")
		log.Fatalln("Read Config file error,Quit.")
	}
}

func main() {
	r := router()
	log.Println("server is running...")
	http.ListenAndServe(config.Port, r)
}
