package models

import (
	"crypto/rand"
	"encoding/base64"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`

	} `yaml:"database"`
	Log struct {
		Json bool `yaml:"json"`
		InfoFile string `yaml:"info_file"`
		ErrorFile string `yaml:"error_file"`
	} `yaml:"log"`
	Auth struct {
		AccessTokenSecret string `yaml:"access_token_secret"`
		RefreshTokenSecret string `yaml:"refresh_token_secret"`
	} `yaml:"auth"`
	Discord struct {
		Token string `yaml:"token"`
		Prefix string `yaml:"prefix"`
		Guild string `yaml:"guild_id"`
		LogChannel string `yaml:"log_channel_id"`
		MuteRole string `yaml:"mute_role_id"`
	} `yaml:"discord"`

}

func NewConfig() Config {
	cryptoRand(128)
	c := Config{
		Server: struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		}{
			Host: "localhost",
			Port: "8080",
		},
		Database: struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Name     string `yaml:"name"`
		}{
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "",
			Name:     "db",
		},
		Log: struct {
			Json bool `yaml:"json"`
			InfoFile string `yaml:"info_file"`
			ErrorFile string `yaml:"error_file"`
		}{
			Json: false,
			InfoFile: "info.log",
			ErrorFile: "error.log",
		},
		Auth: struct {
			AccessTokenSecret string `yaml:"access_token_secret"`
			RefreshTokenSecret string `yaml:"refresh_token_secret"`
		}{
			AccessTokenSecret: cryptoRand(256),
			RefreshTokenSecret: cryptoRand(128),
		},
		Discord: struct {
			Token  string `yaml:"token"`
			Prefix string `yaml:"prefix"`
			Guild string `yaml:"guild_id"`
			LogChannel string `yaml:"log_channel_id"`
			MuteRole string `yaml:"mute_role_id"`
		}{
			Token:  "",
			Prefix: "!",
			Guild: "",
			LogChannel: "",
			MuteRole: "",
		},
	}
	return c
}

// generate random bytes
func cryptoRand(l int) string {
	b := make([]byte, l)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
