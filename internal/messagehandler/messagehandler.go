package messagehandler

import (
	"errors"
	"strings"

	"github.com/szeist/home-bot/internal/config"
	"github.com/szeist/home-bot/internal/events"
	"github.com/szeist/home-bot/internal/facebook"
	"github.com/szeist/home-bot/internal/kodi"
)

type MessageHandler struct {
	Message  chan *events.Message
	Response chan *events.Response
	kodi     *kodi.Kodi
}

func New(cfg *config.Config) *MessageHandler {
	return &MessageHandler{
		Message:  make(chan *events.Message),
		Response: make(chan *events.Response),
		kodi:     kodi.New(cfg.Kodi.Address, cfg.Kodi.ConnectTimeout),
	}
}

func (m *MessageHandler) Start() {
	for {
		msg := <-m.Message
		textResponse, err := m.handleMessage(msg.Text)
		m.Response <- &events.Response{
			Text:   textResponse,
			Sender: msg.Sender,
			Error:  err,
		}
	}
}

func (m *MessageHandler) handleMessage(cmd string) (string, error) {
	if strings.HasPrefix(cmd, "play ") {
		mediaURL := strings.TrimPrefix(cmd, "play ")
		if !strings.HasPrefix(mediaURL, "https://www.facebook.com/") {
			return "", errors.New("Video source not supported")
		}
		videoURL, err := facebook.GetVideoURL(mediaURL)
		if err != nil {
			return "", err
		}
		err = m.kodi.PlayURL(videoURL)
		if err != nil {
			return "", err
		}
		return "Playing video: " + mediaURL, nil
	}
	return "Could not unserstand: " + cmd, nil
}
