package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Mode string

const (
	ModeDevelop Mode = "develop"
	ModeServe   Mode = "serve"

	ConfigName = "das"
	ConfigPath = "."
	ConfigType = "yaml"
)

type Config struct {
	Name     string   `mapstructure:"name"`
	Version  string   `mapstructure:"version"`
	Http     Http     `mapstructure:"http"`
	Database Database `mapstructure:"database"`
	Mode     Mode     `mapstructure:"mode"`
}

func (c *Config) Load(confFile string) error {
	//_, err := toml.DecodeFile(confFile, c)
	//return err
	return nil
}

type Http struct {
	Port int `mapstructure:"port"`
}

type Database struct {
	Dialect string `mapstructure:"dialect"`
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
	User    string `mapstructure:"user"`
	Pwd     string `mapstructure:"pwd"`
	Name    string `mapstructure:"name"`
}

var defaultConf = Config{
	Name:    "DAS",
	Version: "1.0.0",
	Http: Http{
		Port: 8080,
	},
	Database: Database{
		Dialect: "mysql",
		Address: "127.0.0.1",
		Port:    3306,
		User:    "root",
		Pwd:     "123456",
		Name:    "das",
	},
	Mode: ModeServe,
}

var globalConf = defaultConf

func GetConfig() *Config {
	return &globalConf
}

func GetDefaultConfig() *Config {
	return &defaultConf
}

func init() {
	viper.SetConfigName(ConfigName)
	viper.AddConfigPath(ConfigPath)
	viper.SetConfigType(ConfigType)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error config file: %s \n", err))
	} else {
		err = viper.Unmarshal(&globalConf)
		if err != nil {
			panic(fmt.Errorf("Fatal error when reading %s config, unable to decode into struct, %v", ConfigName, err))
		}
	}
}
