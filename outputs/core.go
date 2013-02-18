package outputs

import "logorezka/core"

type Writer interface {
	Write(event *core.Event)
}

func Write(outputs *[]Writer, event *core.Event) {
	for _, val := range(*outputs) {
		val.Write(event)
	}
}
