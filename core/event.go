package core

import (
	"encoding/json"
	"fmt"
)

func (event *Event)MarshalJSON() ([]byte, error) {
	dict := make(map[string]string)
	dict["@source_host"] = event.Source_host
	dict["@source_path"] = event.Source_path
	dict["@message"] = event.Message
	dict["@type"] = event.Type
	dict["@timestamp"] = event.Timestamp
	res, err := json.Marshal(dict)
	fmt.Println(string(res))
	return res, err
}

type Event struct {
	Source_host string
	Source_path string
	Message     string
	Type        string
	Timestamp   string
}
