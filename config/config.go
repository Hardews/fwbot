/**
 * @Author: Hardews
 * @Date: 2023/5/13 23:52
 * @Description:
**/

package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

var Config CONFIG

type CONFIG struct {
	Mysql `yaml:"mysql"`
	Gpt   `yaml:"gpt"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	DbName   string `yaml:"db-name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Gpt struct {
	ApiKey       string `yaml:"api-key"`
	Organization string `yaml:"organization"`
	Model        string `yaml:"model"`
	Role         string `yaml:"role"`
}

func SetConfig() {
	file, err := os.Open("./config/config.yaml")
	if err != nil {
		log.Fatalln("can not open config file,err:", err)
	}

	configByte, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("can not read config file,err:", err)
	}

	yaml.Unmarshal(configByte, &Config)
}
