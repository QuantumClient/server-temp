package main

import (
	"errors"
	"flag"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"quantumclient.org/backend/v2/models"
	"quantumclient.org/backend/v2/server"
)


func main() {

	path, err := parseFlags()
	if err != nil {
		println(err.Error())
	}
	config, _ := parseConfig(path)
	setUpLogging(config)
	svr := server.NewServer(config)

	svr.Init()

	svr.Start()
}

func parseFlags() (string, error) {

	path := flag.String("config", "./config.yml", "path to config file")

	flag.Parse()

	s, err := os.Stat(*path)

	if err != nil || s.IsDir() {
		return "/config.yml", errors.New("config file not found")
	}

	return *path, nil
}

func parseConfig(configFile string) (*models.Config, error) {
	var config models.Config
	file, err := os.Open(configFile)
	if err != nil {
		return generateConfig(), err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	d.KnownFields(true)

	if err = d.Decode(&config); err != nil {
		return generateConfig(), errors.New("config file is not valid")
	}
	return &config, nil
}

func generateConfig() *models.Config {
	println("Generating config file")
	config := models.NewConfig()
	b, _ := yaml.Marshal(config)
	ioutil.WriteFile("./config.yml", b, 0644)
	return &config
}

func setUpLogging(c *models.Config)  {

	// Log to error file
	errPath, err := checkFilePath(c.Log.ErrorFile, "./logs/error.log", false)
	if err != nil {
		slog.Fatal(err)
	}
	errorHandler := handler.MustFileHandler(errPath, c.Log.Json)
	errorHandler.Levels = slog.Levels{slog.PanicLevel, slog.FatalLevel, slog.ErrorLevel, slog.WarnLevel}

	// Log to info file
	infoPath, err := checkFilePath(c.Log.InfoFile, "./logs/info.log", false)
	if err != nil {
		slog.Fatal(err)
	}

	infoHandler := handler.MustFileHandler(infoPath, c.Log.Json)
	infoHandler.Levels = slog.Levels{slog.InfoLevel, slog.NoticeLevel, slog.DebugLevel, slog.TraceLevel}

	// Add handlers
	slog.AddHandlers(errorHandler, infoHandler)

	go slog.FlushDaemon()

}


func checkFilePath(path, fallBack string, existCheck bool) (string, error) {
	if s, err := os.Stat(path); existCheck && (err != nil || s.IsDir()) {
		return fallBack, errors.New("config file not found")
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fallBack, errors.New("config file not found")
	}
	return absPath, nil
}