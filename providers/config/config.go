package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

//Config for application
type Config struct {
	URL           string
	WebPath       string
	TemplatesPath string
	Theme         string
}

func (c Config) String() string {
	return fmt.Sprintf("%s", c.URL)
}

// LoadConfig return a toml configuration
func LoadConfig() *Config {
	conf := Config{}
	content, err := ioutil.ReadFile("./auth.conf")
	if err != nil {
		log.Printf("No configuration file found. %s", err)
		os.Exit(-1)
	}
	configData := string(content)
	if _, err := toml.Decode(configData, &conf); err != nil {
		log.Println("File does not contain a valid configuration. Using default values.")
	}
	conf.TemplatesPath = path.Join(conf.WebPath, "templates")
	return &conf
}
