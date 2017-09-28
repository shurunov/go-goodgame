package goodgame

type LinksPaginate struct {
	Self  Href `json:"self,omitempty"`
	First Href `json:"first,omitempty"`
	Last  Href `json:"last,omitempty"`
	Next  Href `json:"next,omitempty"`
}

type LinksSingle struct {
	Self Href `json:"self,omitempty"`
}

type Href struct {
	Href string `json:"href,omitempty"`
}
