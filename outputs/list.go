package outputs

import "logorezka/outputs/elasticsearch"
import "logorezka/outputs/file"

func AllOutputs () *[]Writer {
	return &[]Writer {
		file.Output(0),
		elasticsearch.Output(0),
	}
}
