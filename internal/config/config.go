package config

import (
	"os"
	"time"
)

type Config struct {
	Keybase *KBConfig
	Kodi    *KodiConfig
}

type KBConfig struct {
	User     string
	PaperKey string
	Location string
	Home     string
}

type KodiConfig struct {
	Address        string
	ConnectTimeout time.Duration
}

func New() *Config {
	return &Config{
		Keybase: &KBConfig{
			Location: "keybase",
		},
		Kodi: &KodiConfig{
			ConnectTimeout: 2,
		},
	}
}

func NewFromEnv() *Config {
	c := New()
	c.Keybase.User = os.Getenv("KB_USER")
	c.Keybase.PaperKey = os.Getenv("KB_PAPERKEY")
	c.Keybase.Home = os.Getenv("KB_HOME")
	c.Kodi.Address = os.Getenv("KODI_ADDRESS")
	return c
}
