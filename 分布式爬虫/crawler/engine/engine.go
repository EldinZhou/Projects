package engine

import (
	"Projects/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requestsQ []Request
	for _, r := range seeds{
		requestsQ = append(requestsQ, r)
	}

	for len(requestsQ) > 0 {
		r := requestsQ[0]
		requestsQ = requestsQ[1:]

		log.Printf("Fetching %s", r.Url)
		//Fetch
		body, err := fetcher.Fetch(r.Url)
		//fmt.Println(string(body))
		if err != nil {
			//忽略掉拿不到的
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		//Parser
		ParseResult := r.ParserFunc(body)
		requestsQ = append(requestsQ, ParseResult.Requests...)//[0] [1] [2] ...

		for _, item := range ParseResult.Items {
			log.Printf("Got item %v", item) //%v: 不转义， %s这个方法不好
		}
	}
}
