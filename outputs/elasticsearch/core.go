package elasticsearch

import (
	el_core "github.com/mattbaird/elastigo/core"
	"logorezka/core"
	"time"
)

// fake type to specify interface :(
type Output int;

func (Output)Write(event *core.Event) {
	index := "logstash-" + time.Now().Format("2006.01.02")
	_, err := el_core.Index(true, index, "default", "", event)
	if (err != nil) {
		panic(err.Error())
	}
}
