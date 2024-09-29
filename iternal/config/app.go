package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type AppCfg struct {
	Addr         string
	DB           DBCfg
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DBCfg struct {
	Host     string
	Port     uint16
	DB       string
	User     string
	Password string
}

func Load() (AppCfg, error) {

	if err := godotenv.Load(".env"); err != nil {
		return AppCfg{}, err
	}

	readTimeoutInSec, err := strconv.Atoi(os.Getenv("readTimeoutInSec"))
	if err != nil {
		return AppCfg{}, err
	}

	writeTimeoutInSec, err := strconv.Atoi(os.Getenv("writeTimeoutInSec"))
	if err != nil {
		return AppCfg{}, err
	}

	port, err := strconv.Atoi(os.Getenv("db.port"))
	if err != nil {
		return AppCfg{}, err
	}

	return AppCfg{
		Addr:         os.Getenv("addr"),
		ReadTimeout:  time.Duration(readTimeoutInSec) * time.Second,
		WriteTimeout: time.Duration(writeTimeoutInSec) * time.Second,
		DB: DBCfg{
			Host:     os.Getenv("db.host"),
			Port:     uint16(port),
			DB:       os.Getenv("db.db"),
			User:     os.Getenv("db.user"),
			Password: os.Getenv("db.password"),
		},
	}, nil
}
