package helper

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
	Type string `json:"type,omitempty"`
}

type HateoasResponse[T any] struct {
	Data  T      `json:"data"`
	Links []Link `json:"links"`
}
