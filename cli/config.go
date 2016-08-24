package cli

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type Config struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Server   string `json:"server" yaml:"server"`
}

func (c Config) Json() string {
	j, _ := json.Marshal(c)
	return string(j)
}

func (c Config) Yaml() string {
	j, _ := yaml.Marshal(c)
	return string(j)
}

func Configure(c *cli.Context) error {
	if c.NumFlags() != 3 || len(c.String(FlCBUsername.Name)) == 0 || len(c.String(FlCBPassword.Name)) == 0 || len(c.String(FlCBServer.Name)) == 0 {
		return cli.NewExitError(fmt.Sprintf("You need to specify all the parameters. See '%s configure --help'.", c.App.Name), 1)
	}

	err := WriteConfigToFile(c.String(FlCBServer.Name), c.String(FlCBUsername.Name), c.String(FlCBPassword.Name))
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("[WriteConfigToFile] %s", err.Error()), 1)
	}
	return nil
}

func GetCurrentUser() *user.User {
	currentUser, err := user.Current()
	if err != nil {
		log.Printf("[GetCurrentUser] %s", err.Error())
		os.Exit(1)
	}
	return currentUser
}

func WriteConfigToFile(server string, username string, password string) error {
	currentUser := GetCurrentUser()
	log.Printf("[WriteConfigToFile] current user: %s", currentUser.Username)

	hdcDir := currentUser.HomeDir + "/" + hdc_dir
	configFile := hdcDir + "/" + config_file

	if _, err := os.Stat(hdcDir); os.IsNotExist(err) {
		log.Printf("[WriteCredentialsToFile] create dir: %s", hdcDir)
		err = os.MkdirAll(hdcDir, 0744)
		if err != nil {
			return err
		}
	} else {
		log.Printf("[WriteConfigToFile] dir already exists: %s", hdcDir)
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Printf("[WriteCredentialsToFile] create file: %s", configFile)
		if _, err := os.Create(configFile); err != nil {
			return err
		}
	} else {
		log.Printf("[WriteConfigToFile] file already exists: %s", configFile)
	}

	f, err := os.OpenFile(configFile, os.O_WRONLY, 0744)
	if err != nil {
		return err
	}

	log.Printf("[WriteConfigToFile] writing credentials to file: %s", configFile)
	confJson := Config{Server: server, Username: username, Password: password}.Yaml()
	if _, err = f.WriteString(confJson); err != nil {
		return err
	}
	f.Close()

	return nil
}

func ReadConfig() (*Config, error) {
	currentUser := GetCurrentUser()
	log.Printf("[ReadConfig] current user: %s", currentUser.Username)

	hdcDir := currentUser.HomeDir + "/" + hdc_dir
	configFile := hdcDir + "/" + config_file

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, err
	}
	log.Printf("[ReadConfig] found config file: %s", configFile)

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
