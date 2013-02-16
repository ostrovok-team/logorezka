package core

import (
	"flag"
	"os"
)

var ConfOutputFilesDir string
var ConfInputZmqBind string

func ConfigParse() {
	flag.StringVar(&ConfOutputFilesDir, "output-files-dir", os.TempDir(), "path to write output log files to")
	flag.StringVar(&ConfInputZmqBind, "input-zmq-bind", "tcp://0.0.0.0:2120", "address zmq to bind on")
	flag.Parse()
}
