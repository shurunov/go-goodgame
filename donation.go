package goodgame

import "fmt"

type DonationsPaginate struct {
	Links      LinksPaginate `json:"_links,omitempty"`
	Embedded   Donations     `json:"_embedded,omitempty"`
	PageCount  int           `json:"page_count,omitempty"`
	PageSize   int           `json:"page_size,omitempty"`
	TotalItems int           `json:"total_items,omitempty"`
	Page       int           `json:"page,omitempty"`
}

type Donations struct {
	Donations []Donation `json:"donations,omitempty"`
}

type Donation struct {
	ID       string      `json:"id,omitempty"`
	Username string      `json:"username,omitempty"`
	Amount   string      `json:"amount,omitempty"`
	PaidDate string      `json:"paid_date,omitempty"`
	Comment  string      `json:"comment,omitempty"`
	Links    LinksSingle `json:"_links,omitempty"`
}

type DonationOptions struct {
	Page          int    `url:"page,omitempty"`
	FromTimestamp string `url:"from_timestamp,omitempty"`
}

type DonationRequest struct {
	client *Client
}

// Получение списка донатов
func (r DonationRequest) All(channel string, options *DonationOptions) (*DonationsPaginate, error) {
	p := fmt.Sprintf("channel/%s/donations", channel)
	path, err := r.client.path(p, options)
	if err != nil {
		return nil, err
	}

	donations := &DonationsPaginate{}
	_, err = r.client.Get(path, donations)
	if err != nil {
		return nil, err
	}

	return donations, nil
}
