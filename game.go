package goodgame

import "fmt"

type GamesPaginate struct {
	Links      LinksPaginate `json:"_links,omitempty"`
	Embedded   Games         `json:"_embedded,omitempty"`
	PageCount  int           `json:"page_count,omitempty"`
	PageSize   int           `json:"page_size,omitempty"`
	TotalItems int           `json:"total_items,omitempty"`
	Page       int           `json:"page,omitempty"`
}

type Games struct {
	Game []Game `json:"games,omitempty"`
}

type Game struct {
	ID     string      `json:"id,omitempty"`
	Title  string      `json:"title,omitempty"`
	Url    string      `json:"url,omitempty"` // TODO: создать отдельный c анмаршалингом
	Short  string      `json:"short,omitempty"`
	Poster string      `json:"poster,omitempty"`
	Links  LinksSingle `json:"_links,omitempty"`
}

type GameOptions struct {
	Page   int    `url:"page,omitempty"`
	Filter string `url:"filter,omitempty"`
}

type GameRequest struct {
	client *Client
}

// Получение информации об зарегистрированных в сервисе играх
func (r GameRequest) All(options *GameOptions) (*GamesPaginate, error) {
	path, err := r.client.path("games", options)
	if err != nil {
		return nil, err
	}

	games := &GamesPaginate{}
	_, err = r.client.Get(path, games)
	if err != nil {
		return nil, err
	}

	return games, nil
}

// Получение информации по игре, зная ее url
func (r GameRequest) Get(channel string) (*Game, error) {
	p := fmt.Sprintf("games/%s", channel)
	path, err := r.client.path(p, nil)
	if err != nil {
		return nil, err
	}

	game := &Game{}
	_, err = r.client.Get(path, game)
	if err != nil {
		return nil, err
	}

	return game, nil
}
