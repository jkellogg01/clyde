package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/bubbles/textinput"
)

type Query struct {
	URL     string
	Request func(string) (*http.Response, error)
}

func main() {
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()
	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

    input := textinput.New()
	r := bufio.NewReader(os.Stdin)
	fmt.Print(input.View())
	rawQuery, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read input", "error", err)
	}
	// this removes the newline from the input we took from stdin
	rawQuery = rawQuery[:len(rawQuery)-1]
	method, userUrl, didCut := strings.Cut(rawQuery, " ")
	if !didCut {
		log.Fatalf("Malformed user input: %s", rawQuery)
	}
	query, err := newQuery(method, userUrl)
	if err != nil {
		log.Fatal(err)
	}
	queryTime := time.Now()
	log.Debug("sending request", "url", userUrl)
	resp, err := query.execute()
	if err != nil {
		log.Fatal(err)
	}
    respTime := time.Since(queryTime)
	log.Debug("server responded!", "duration", respTime)
	mustPrintHTTPResponse(resp, respTime)
}

func newQuery(method, strurl string) (Query, error) {
	methodFunc, err := getMethodFunc(method)
	if err != nil {
		return Query{}, err
	}
	return Query{
		Request: methodFunc,
		URL:     strurl,
	}, nil
}

func getMethodFunc(method string) (func(string) (*http.Response, error), error) {
	upperMethod := strings.ToUpper(method)
	log.Debug("dispatching text method to http request function", "method", upperMethod)
	switch upperMethod {
	case http.MethodGet:
		return http.Get, nil
	default:
		return nil, fmt.Errorf("http method '%s' is not currently supported", upperMethod)
	}
}

func (q Query) execute() (*http.Response, error) {
	return q.Request(q.URL)
}

func mustPrintHTTPResponse(r *http.Response, respTime time.Duration) {
	decoder := json.NewDecoder(r.Body)
	var data interface{}
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
    var status string
    switch r.StatusCode / 100 {
    case 2:
        status = styleStatusGreen.Render(r.Status)
    case 1, 3:
        status = styleStatusYellow.Render(r.Status)
    default:
        status = styleStatusRed.Render(r.Status)
    }
	fmt.Printf("Responded with %s in %v\n", status, respTime)
	fmt.Println("Headers:")
	for k, v := range r.Header {
		fmt.Printf("\t%s: %s\n", k, strings.Join(v, ", "))
	}
	fmt.Printf("Response body:\n%s\n", body)
}
