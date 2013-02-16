package core

import (
	"flag"
)

var ConfOutputFilesDir string
var ConfInputZmqBind string

func ConfigParse() {
	flag.StringVar(&ConfOutputFilesDir, "output-files-dir", "/tmp", "path to write output log files to")
	flag.StringVar(&ConfInputZmqBind, "input-zmq-bind", "tcp://0.0.0.0:2120", "address zmq to bind on")
	flag.Parse()
}
