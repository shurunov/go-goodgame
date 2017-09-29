package goodgame

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client
	url    *url.URL
	Stream *StreamRequest
	Game   *GameRequest
	Smile  *SmileRequest
}

const apiUrl = "http://api2.goodgame.ru/v2/"

// Создание нового клиента для работы с API сервиса
func NewClient() *Client {
	c := &Client{client: &http.Client{}}

	c.url, _ = url.Parse(apiUrl)
	c.Stream = &StreamRequest{client: c}
	c.Game = &GameRequest{client: c}
	c.Smile = &SmileRequest{client: c}

	return c
}

// Передается адрес запроса и структура, которая
// заполнится данными из JSON пришедшего ответа
func (c Client) Get(path string, r interface{}) (*http.Response, error) {
	p, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	u := c.url.ResolveReference(p)
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/vnd.goodgame.v2+json")

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(r); err != nil {
		return nil, err
	}

	return response, nil
}

// Получение полного адреса для отправки запросов
// В options попадает структура для настрйки запроса
func (c Client) path(path string, options interface{}) (string, error) {
	if options != nil {
		v, err := query.Values(options)
		if err != nil {
			return "", err
		}

		path += "?" + v.Encode()
	}

	return path, nil
}
