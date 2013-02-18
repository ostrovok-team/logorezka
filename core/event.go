package core

import (
	"encoding/json"
)

func (event *Event)MarshalJSON() ([]byte, error) {
	dict := make(map[string]string)
	dict["@source_host"] = event.Source_host
	dict["@source_path"] = event.Source_path
	dict["@message"] = event.Message
	return json.Marshal(dict)
}

type Event struct {
	Source_host string
	Source_path string
	Message     string
}
