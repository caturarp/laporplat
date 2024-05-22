package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

// Environment utility

func loadEnvStr(key string, result *string) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return
	}

	*result = s
}

// Configuration

func defaultConfig() config {
	return config{
		Listen:   defaultListenConfig(),
		DBConfig: defaultPgConfig(),
	}
}

type listenConfig struct {
	Host string `yaml:"host" json:"host"`
	Port uint   `yaml:"port" json:"port"`
}

func (l listenConfig) Address() string { return fmt.Sprintf("%s:%d", l.Host, l.Port) }
func defaultListenConfig() listenConfig {
	return listenConfig{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

func (l *listenConfig) loadFromEnv() {
	loadEnvStr("LISTEN_HOST", &l.Host)
	loadEnvUint("LISTEN_PORT", &l.Port)
}

type pgConfig struct {
	Host    string `yaml:"host" json:"host"`
	Port    uint   `yaml:"port" json:"port"`
	Pass    string `yaml:"pass" json:"pass"`
	User    string `yaml:"user" json:"user"`
	DBName  string `yaml:"db_name" json:"name"`
	SslMode string `yaml:"ssl_mode" json:"ssl_mode"`
}

func (p pgConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d config=%s sslmode=%s", p.Host, p.Port, p.DBName, p.SslMode)
}

func defaultPgConfig() pgConfig {
	return pgConfig{
		Host:    "localhost",
		Port:    5432,
		DBName:  "todo",
		SslMode: "disable",
	}
}

func (p *pgConfig) loadFromEnv() {
	loadEnvStr("DB_HOST", &p.Host)
	loadEnvStr("DB_USER", &p.User)
	loadEnvStr("DB_PASS", &p.Pass)
	loadEnvUint("DB_PORT", &p.Port)
	loadEnvStr("DB_NAME", &p.DBName)
	loadEnvStr("DB_SSL", &p.SslMode)
}

type config struct {
	Listen   listenConfig `yaml:"listen" json:"listen"`
	DBConfig pgConfig     `yaml:"db" json:"db"`
}

func (c *config) loadFromEnv() {
	c.Listen.loadFromEnv()
	c.DBConfig.loadFromEnv()
}

func loadEnvUint(key string, result *uint) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return
	}

	*result = uint(n) // will clamp the negative value
}

func loadConfigFromReader(r io.Reader, c *config) error {
	return yaml.NewDecoder(r).Decode(c)
}

func loadConfigFromFile(fn string, c *config) error {
	_, err := os.Stat(fn)
	if err != nil {
		return err
	}

	f, err := os.Open(fn)
	if err != nil {
		return err
	}

	defer f.Close()

	return loadConfigFromReader(f, c)
}
