package main

import (
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"sync"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
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
	configMu sync.RWMutex
)

var (
	rooms map[string]Room
	roomsMu sync.RWMutex
)

// Room descript the client info
type Room struct {
	Name   string `yaml:"name"`
	Num    string `yaml:"num"`
	Neibor string `yaml:"neibor"`
}

func parseYaml() {
	user, _ := user.Current()
	filename := path.Join(user.HomeDir, ".nhweb", "rooms.yml")
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Warning: Can't open rooms.yml file: %v", err)
		rooms = make(map[string]Room) // Initialize empty map
		return
	}
	
	roomsMu.Lock()
	defer roomsMu.Unlock()
	err = yaml.Unmarshal(yamlFile, &rooms)
	if err != nil {
		log.Printf("Warning: Failed to parse rooms.yml: %v", err)
		rooms = make(map[string]Room) // Initialize empty map
	}
}

func init() {
	parseYaml()
	user, _ := user.Current()
	configPath := path.Join(user.HomeDir, ".nhweb", "config.toml")
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		log.Println("Write Config File with following fields:")
		log.Println("Port Pass SessionKey IPS")
		log.Fatalln("Read Config file error,Quit.")
	}
}

func main() {
	r := router()
	log.Printf("server is running on %s...", config.Port)
	
	server := &http.Server{
		Addr:    config.Port,
		Handler: r,
	}
	
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
