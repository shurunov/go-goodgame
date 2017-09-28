package goodgame

import "fmt"

type StreamsPaginate struct {
	Links      LinksPaginate `json:"_links,omitempty"`
	Embedded   Streams       `json:"_embedded,omitempty"`
	PageCount  int           `json:"page_count,omitempty"`
	PageSize   int           `json:"page_size,omitempty"`
	TotalItems int           `json:"total_items,omitempty"`
	Page       int           `json:"page,omitempty"`
}

type Streams struct {
	Streams []Stream `json:"streams,omitempty"`
}

type Channel struct {
	ID          int      `json:"id,omitempty"`
	Key         string   `json:"key,omitempty"`
	Premium     string   `json:"premium,omitempty"`
	Title       string   `json:"title,omitempty"`
	MaxViewers  int      `json:"max_viewers,omitempty"`
	PlayerType  string   `json:"player_type,omitempty"`
	GGPlayerSrc string   `json:"gg_player_src,omitempty"`
	Embed       string   `json:"embed,omitempty"`
	Img         string   `json:"img,omitempty"`
	Thumb       string   `json:"thumb,omitempty"`
	Description string   `json:"description,omitempty"`
	Adult       bool     `json:"adult,omitempty"`
	Games       []SGames `json:"games,omitempty"`
	Url         string   `json:"url,omitempty"` // TODO: создать отдельный c анмаршалингом
}

type SGames struct {
	Title string `json:"title,omitempty"`
	Url   string `json:"url,omitempty"`
}

type Stream struct {
	RequestKey       string      `json:"request_key,omitempty"`
	ID               int         `json:"id,omitempty"`
	Key              string      `json:"key,omitempty"`
	IsBroadcast      bool        `json:"is_broadcast,omitempty"`
	BroadcastStarted int         `json:"broadcast_started,omitempty"`
	BroadcastEnd     int         `json:"broadcast_end,omitempty"`
	Url              string      `json:"url,omitempty"`
	Status           string      `json:"status,omitempty"`
	Viewers          string      `json:"viewers,omitempty"`
	PlayerViewers    string      `json:"player_viewers,omitempty"`
	UserInChat       string      `json:"users_in_chat,omitempty"`
	Channel          Channel     `json:"channel,omitempty"`
	Links            LinksSingle `json:"_links,omitempty"`
}

type StreamRequest struct {
	client *Client
}

type StreamOptions struct {
	IDs    string `url:"ids,omitempty"`
	Page   int    `url:"page,omitempty"`
	OnlyGG bool   `url:"only_gg,omitempty"`
	Adult  bool   `url:"adult,omitempty"`
	Game   string `url:"game,omitempty"`
	Hidden bool   `url:"hidden,omitempty"`
}

// Получение информации о трансляциях
func (r StreamRequest) All(options *StreamOptions) (*StreamsPaginate, error) {
	path, err := r.client.path("streams", options)
	if err != nil {
		return nil, err
	}

	streams := &StreamsPaginate{}
	_, err = r.client.Get(path, streams)
	if err != nil {
		return nil, err
	}

	return streams, nil
}

// Получение информации о конкретном стриме
func (r StreamRequest) Get(channel string) (*Stream, error) {
	p := fmt.Sprintf("streams/%s", channel)
	path, err := r.client.path(p, nil)
	if err != nil {
		return nil, err
	}

	stream := &Stream{}
	_, err = r.client.Get(path, stream)
	if err != nil {
		return nil, err
	}

	return stream, nil
}
