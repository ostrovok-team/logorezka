package main

import "logorezka/core"
import "logorezka/inputs"
import "logorezka/outputs"

func main() {
	core.ConfigParse()
	inputs.InputRun(outputs.AllOutputs())
}
