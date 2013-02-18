package inputs

import (
	"bytes"
	zmq "github.com/alecthomas/gozmq"
	msgpack "github.com/msgpack/msgpack-go"
	"logorezka/core"
	"logorezka/outputs"
	"reflect"
)

func msgFieldAsString(obj *reflect.Value, field_name string) string {
	field_obj := obj.MapIndex(reflect.ValueOf(field_name))
	return string(field_obj.Interface().([]byte))
}

func msgFieldAsBytes(obj *reflect.Value, field_name string) []byte {
	field_obj := obj.MapIndex(reflect.ValueOf(field_name))
	return field_obj.Interface().([]byte)
}

func InputRun(outputs_list *[]outputs.Writer) {
	context, err := zmq.NewContext()
	if err != nil {
		panic(err.Error())
	}

	socket, err := context.NewSocket(zmq.PULL)
	if err != nil {
		panic(err.Error())
	}

	err = socket.Bind(core.ConfInputZmqBind)
	if err != nil {
		panic(err.Error())
	}

	for {
		msg, err := socket.Recv(0)
		if err != nil {
			panic(err.Error())
		}

		buffer := bytes.NewBuffer(msg)
		obj, _, err := msgpack.Unpack(buffer)
		if err != nil {
			panic(err.Error())
		}

		event := core.Event{
			Source_host: msgFieldAsString(&obj, "@source_host"),
			Source_path: msgFieldAsString(&obj, "@source_path"),
			Message:     msgFieldAsString(&obj, "@message"),
			Type:        msgFieldAsString(&obj, "@type"),
			Timestamp:   msgFieldAsString(&obj, "@timestamp"),
		}

		outputs.Write(outputs_list, &event)
	}
}
