package main

// api => https://dummyjson.com/quotes/:id

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

// fill struct pointer from byte data
func unmarshalQuoteModel(data []byte, quote *QuoteModel) (err error) {
	err = json.Unmarshal(data, &quote)
	return err
}

// struct response body
type QuoteModel struct {
	ID     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

// go routine
func getQuotes(inc *int, ch chan string) {
	time.Sleep(time.Second * 2) // make delay
	fmt.Printf("[SYSTEM] Retrieving quotes %d...\n", *inc)
	time.Sleep(time.Second * 2) // make delay

	_, resp, err := fasthttp.Get(nil, fmt.Sprintf("https://dummyjson.com/quotes/%d", *inc)) //get response

	if err != nil {
		log.Fatalln(err)
	}

	var quoteModel QuoteModel
	unmarshalQuoteModel(resp, &quoteModel) // unmarshal byte response to quoteModel struct

	*inc += 1
	ch <- fmt.Sprintf("[QUOTES] \"%s\", %s", quoteModel.Quote, quoteModel.Author) // assign data into channel
}

func main() {
	channel := make(chan string)
	var inc int = 1

	for { // infinite loop
		go getQuotes(&inc, channel)
		fmt.Println(<-channel) // retrieve data from channel
	}
}
