package request

import (
	"log"
)

type Request struct {
	Method string
	Secure bool // this denotes http vs. https
	User   string
	Pass   string
	Host   string
	Port   string
	Path   string
	// etc etc etc... sticking with the barest version for now
}

func Parse(s string) (*Request, error) {
	r := new(Request)
	t := tokenizer{
		body: s,
	}
	err := t.tokenize()
	if err != nil {
		return nil, err
	}
	log.Print(t.tokens)
	return r, nil
}
