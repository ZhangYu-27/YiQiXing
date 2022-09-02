package tool

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Mysql *Mysql `yaml:"mysql"`
}

type Mysql struct {
	Port     string `yaml:"port"`
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Db_name  string `yaml:"dbname"`
}

var config Config

func (this Config) initYaml() {
	config = Config{}
	yamlFile, err := ioutil.ReadFile("/www/wwwroot/config.yaml")
	fmt.Println(yamlFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("config.log: %#v\n", config.Mysql)

}
func (this Config) GetMysqlEnv() *Mysql {
	config.initYaml()
	return config.Mysql
}
