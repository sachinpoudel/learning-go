package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)


type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true"` //this is struct tags	
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == ""{
		flags := flag.String("config", "", "path to the configuration file") //flags are used for command line arguments simple mean is -config <path>
		flag.Parse()

		configPath = *flags

		if configPath == ""{
			log.Fatal("config file is not set")


		}
	}
	if _, err := os.Stat(configPath);
	     os.IsNotExist(err){
		log.Fatalf("config does not exits: %s")
	}
	var cfg Config

	err:= cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file : %s", err.Error())

	}
	return &cfg
}