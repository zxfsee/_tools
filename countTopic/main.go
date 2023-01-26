package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

type Wiki struct {
	Parse struct {
		Title  string `json:"title"`
		Pageid int    `json:"pageid"`
		Text   string `json:"text"`
	} `json:"parse"`
}

func countTopic(topic string) int {
	wikipedia := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&page=%s&prop=text&formatversion=2&format=json", url.QueryEscape(topic))

	resp, err := http.Get(wikipedia)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	var data Wiki
	for i := 0; scanner.Scan() && i < 1; i++ {
		body := scanner.Bytes()
		err := json.Unmarshal(body, &data)
		if err != nil {
			panic(err.Error())
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	r := regexp.MustCompile(topic)
	matches := r.FindAllStringIndex(data.Parse.Text, -1)

	return len(matches)
}

func main() {
	topic := "Pet door"
	fmt.Println(countTopic(topic))
}