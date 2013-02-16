package core

import (
	"bytes"
	zmq "github.com/alecthomas/gozmq"
	msgpack "github.com/msgpack/msgpack-go"
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

func InputRun() {
	context, err := zmq.NewContext()
	if err != nil {
		panic(err.Error())
	}

	socket, err := context.NewSocket(zmq.PULL)
	if err != nil {
		panic(err.Error())
	}

	err = socket.Bind(ConfInputZmqBind)
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

		event := Event{
			_source_host: msgFieldAsString(&obj, "@source_host"),
			_source_path: msgFieldAsString(&obj, "@source_path"),
			_message:     msgFieldAsBytes(&obj, "@message"),
		}

		OutputEvent(&event)
	}
}
