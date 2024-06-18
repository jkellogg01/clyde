package request

import (
	"net/http"
)

type Request struct {
	Method string
	Secure bool // this denotes http vs. https
	User   string
	Pass   string
    Host string
    Port string
    Path string
    // etc etc etc... sticking with the barest version for now
}

func Parse(s string) (*Request, error) {
	r := new(Request)
	return r, nil
}

func (r Request) Execute() (*http.Response, error)
