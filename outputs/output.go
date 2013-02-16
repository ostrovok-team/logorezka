package outputs

import "logorezka/core"

func OutputEvent(event *core.Event) {
	OutputToFile(event)
	//	OutputToElasticsearch(event)
}
