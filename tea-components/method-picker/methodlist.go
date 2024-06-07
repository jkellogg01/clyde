package methodpicker

import "net/http"

type methodList struct {
	methods []method
	index   int
}

func NewMethodList() *methodList {
	return &methodList{
		methods: []method{Get, Put, Post, Patch, Delete},
	}
}

func (l *methodList) Advance() {
	l.index += 1
	if l.index >= len(l.methods) {
		l.index = 0
	}
}

func (l methodList) String() string {
	return string(l.methods[l.index])
}

// TODO: deepen the method type to associate some styling with each method
// or maybe this method will just assert that styling instead. who's to say.
func (l methodList) View() string {
    return string(l.methods[l.index])
}

type method string

const (
	Get     method = http.MethodGet
	Put     method = http.MethodPut
	Head    method = http.MethodHead
	Post    method = http.MethodPost
	Patch   method = http.MethodPatch
	Trace   method = http.MethodTrace
	Delete  method = http.MethodDelete
	Connect method = http.MethodConnect
	Options method = http.MethodOptions
)
