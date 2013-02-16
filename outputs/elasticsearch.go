package outputs

import (
	//	el_core "github.com/mattbaird/elastigo/core"
	"logorezka/core"
)

func OutputToElasticsearch(event *core.Event) {
	//	index := "logstash-%{+YYYY.MM.dd}"
	//	client.index(index, type, event.to_hash)
	//	response, _ := el_core.Index(true, "twitter", "tweet", "1", NewTweet("blah", "Search is cool"))
}
