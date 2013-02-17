package outputs

import (
	el_core "github.com/mattbaird/elastigo/core"
	"logorezka/core"
	"time"
)

func OutputToElasticsearch(event *core.Event) {
	index := "logorezka-" + time.Now().Format("2006.01.02")
	_, err := el_core.Index(true, index, "default", "", event)
	if (err != nil) {
		panic(err.Error())
	}
}
