package goodgame

import "fmt"

type Player struct {
	ChannelID      string      `json:"channel_id,omitempty"`
	ChannelKey     string      `json:"channel_key,omitempty"`
	ChannelTitle   string      `json:"channel_title,omitempty"`
	ChannelStatus  string      `json:"channel_status,omitempty"`
	ChannelPoster  string      `json:"channel_poster,omitempty"`
	ChannelPremium bool        `json:"channel_premium,omitempty"`
	ChannelStart   string      `json:"channel_start,omitempty"`
	Viewers        string      `json:"viewers,omitempty"`
	StreamerID     string      `json:"streamer_id,omitempty"`
	StreamerName   string      `json:"streamer_name,omitempty"`
	StreamerAvatar string      `json:"streamer_avatar,omitempty"`
	PremiumOnly    bool        `json:"premium_only,omitempty"`
	Adult          int         `json:"adult,omitempty"`
	Hidden         int         `json:"hidden,omitempty"`
	GACode         string      `json:"ga_code,omitempty"`
	Broadcast      Broadcast   `json:"broadcast,omitempty"`
	Links          LinksSingle `json:"_links,omitempty"`
}

type Broadcast struct {
	Title       string `json:"broadcast_title,omitempty"`
	Start       string `json:"broadcast_start,omitempty"`
	Games       string `json:"broadcast_games,omitempty"`
	Description string `json:"broadcast_description,omitempty"`
	Logo        string `json:"broadcast_logo,omitempty"`
}

type PlayerRequest struct {
	client *Client
}

// Получение информации о плеере
// src - идентификатор плеера
func (r PlayerRequest) Get(src int) (*Player, error) {
	p := fmt.Sprintf("player/%d", src)
	path, err := r.client.path(p, nil)
	if err != nil {
		return nil, err
	}

	player := &Player{}
	_, err = r.client.Get(path, player)
	if err != nil {
		return nil, err
	}

	return player, nil
}
