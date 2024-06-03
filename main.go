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
)

type Query struct {
	URL     string
	Request func(string) (*http.Response, error)
}

func main() {
	devMode := flag.Bool("v", false, "verbose")
	flag.Parse()
	if *devMode {
		log.SetLevel(log.DebugLevel)
		log.Debug("dev mode enabled, you will see debug logging")
	}

	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a query in the format 'METHOD URL'\n> ")
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
	log.Debug("server responded!", "duration", time.Since(queryTime))
	mustPrintHTTPResponse(resp)
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

func mustPrintHTTPResponse(r *http.Response) {
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
	fmt.Printf("Status: %s\n", r.Status)
	fmt.Println("Headers:")
	for k, v := range r.Header {
		fmt.Printf("\t%s: %s\n", k, strings.Join(v, ", "))
	}
	fmt.Printf("Response body:\n%s\n", body)
}
