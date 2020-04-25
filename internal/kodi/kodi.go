package kodi

import (
	"time"

	"github.com/pdf/kodirpc"
)

type Kodi struct {
	address string
	config  *kodirpc.Config
}

type kodiPlayItem struct {
	File string `json:"file"`
}

type kodiPlayParams struct {
	Item kodiPlayItem `json:"item"`
}

func New(address string, connectTimeout time.Duration) *Kodi {
	k := &Kodi{
		address: address,
		config:  kodirpc.NewConfig(),
	}
	k.config.ConnectTimeout = connectTimeout
	return k
}

func (k *Kodi) PlayURL(url string) error {
	kodiClient, err := kodirpc.NewClient(k.address, k.config)
	if err != nil {
		return err
	}

	params := &kodiPlayParams{
		Item: kodiPlayItem{
			File: url,
		},
	}
	_, err = kodiClient.Call("Player.Open", params)
	if err != nil {
		return err
	}
	return kodiClient.Close()
}
