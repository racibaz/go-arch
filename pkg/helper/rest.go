package helper

import (
	"reflect"
	"strings"
)

type Link struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Type   string `json:"type,omitempty"`
	Schema string `json:"schema,omitempty"`
}

type Response[T any] struct {
	Data  *T     `json:"data"`
	Links []Link `json:"_links"`
}

func AddHateoas(rel, href, linkType, schema string) Link {
	return Link{
		Rel:    rel,
		Href:   href,
		Type:   linkType,
		Schema: schema,
	}
}

type FieldSchema struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Required  bool     `json:"required,omitempty"`
	MaxLength int      `json:"maxLength,omitempty"`
	Values    []string `json:"values,omitempty"`
	Rules     string   `json:"rules,omitempty"`
}

/*
type ActionSchema struct {
	Name        string        `json:"name"`
	Method      string        `json:"method"`
	Endpoint    string        `json:"endpoint"`
	ContentType string        `json:"contentType"`
	Fields      []FieldSchema `json:"fields"`
}
*/

func BuildSchemaFromStruct(v any) []FieldSchema {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fields := make([]FieldSchema, 0)

	for i := 0; i < t.NumField(); i++ {

		f := t.Field(i)

		field := FieldSchema{
			Name:     f.Tag.Get("json"),
			Type:     f.Type.String(),
			Required: !strings.Contains(f.Tag.Get("json"), "omitempty"),
			Rules:    f.Tag.Get("validate"),
		}

		fields = append(fields, field)
	}

	return fields
}
