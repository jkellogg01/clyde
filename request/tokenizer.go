package request

import (
	"errors"
	"log"
)

type tokenType int

const (
	text tokenType = iota
	dot
	space
	colon
	equal
	f_slash
	at_sign
	question
	ampersand
)

var (
	ErrInputTerminated     = errors.New("end of input")
	ErrUnexpectedCharacter = errors.New("unexpected character")
)

type token struct {
	tipe tokenType
	text string
}

type tokenizer struct {
	start  int
	curr   int
	body   string
	tokens []token
}

func (t *tokenizer) tokenize() error {
	t.tokens = make([]token, 0)
	for {
		t.start = t.curr
		if t.isAtEnd() {
			return nil
		}
		var tkn token
		c := t.advance()
		switch c {
		case ' ':
			tkn = t.emit(space)
		case '.':
			tkn = t.emit(dot)
		case ':':
			tkn = t.emit(colon)
		case '=':
			tkn = t.emit(equal)
		case '/':
			tkn = t.emit(f_slash)
		case '@':
			tkn = t.emit(at_sign)
		case '?':
			tkn = t.emit(question)
		case '&':
			tkn = t.emit(ampersand)
		}
        // i hate this so much
        emptyToken := token{}
		if tkn != emptyToken {
			t.tokens = append(t.tokens, tkn)
		}
	}
}

func (t *tokenizer) advance() byte {
	t.curr++
	if t.isAtEnd() {
		return 0
	}
	return t.body[t.curr]
}

func (t tokenizer) emit(tipe tokenType) token {
	content := t.body[t.start:t.curr]
	log.Printf("emitting token: body[%d:%d] = %s", t.start, t.curr, content)
	return token{
		tipe: tipe,
		text: content,
	}
}

func (t tokenizer) isAtEnd() bool {
	return t.curr >= len(t.body)
}
