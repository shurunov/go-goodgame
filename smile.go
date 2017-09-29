package goodgame

import "strconv"

type SmilesPaginate struct {
	Links      LinksPaginate `json:"_links,omitempty"`
	Embedded   Smiles        `json:"_embedded,omitempty"`
	PageCount  int           `json:"page_count,omitempty"`
	PageSize   int           `json:"page_size,omitempty"`
	TotalItems int           `json:"total_items,omitempty"`
	Page       int           `json:"page,omitempty"`
}

type Smiles struct {
	Smile []Smile `json:"smiles,omitempty"`
}

type Smile struct {
	Key       string      `json:"key,omitempty"`
	DonateLvl int         `json:"donate_lvl,omitempty"`
	IsPremium bool        `json:"is_premium,omitempty"`
	IsPaid    bool        `json:"is_paid,omitempty"`
	ChannelID string      `json:"channel_id,omitempty"`
	Urls      SmileUrls   `json:"urls,omitempty"`
	Links     LinksSingle `json:"_links,omitempty"`
}

type SmileUrls struct {
	Img string `json:"img,omitempty"` // TODO: создать отдельный c анмаршалингом
	Big string `json:"big,omitempty"` // TODO: создать отдельный c анмаршалингом
	Gif string `json:"gif,omitempty"` // TODO: создать отдельный c анмаршалингом
}

type SmileRequest struct {
	client *Client
}

type SmileOptions struct {
	ChannelID int `url:"channel_id,omitempty"`
	Page      int `url:"page,omitempty"`
}

// Получение информации о смайлах
func (r SmileRequest) All(options *SmileOptions) (*SmilesPaginate, error) {
	p := "smiles"
	if options.ChannelID > 0 {
		p += "/" + strconv.Itoa(options.ChannelID)
	}
	path, err := r.client.path(p, options)
	if err != nil {
		return nil, err
	}

	smiles := &SmilesPaginate{}
	_, err = r.client.Get(path, smiles)
	if err != nil {
		return nil, err
	}

	return smiles, nil
}
