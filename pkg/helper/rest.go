package helper

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
	Type string `json:"type,omitempty"`
}

type Response[T any] struct {
	Data  *T     `json:"data"`
	Links []Link `json:"_links"`
}

func AddHateoas(rel, href, linkType string) Link {
	return Link{
		Rel:  rel,
		Href: href,
		Type: linkType,
	}
}
