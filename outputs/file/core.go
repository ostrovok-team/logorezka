package file

import (
	"bufio"
	"logorezka/core"
	"os"
	"path"
	"strings"
)

var new_line = []byte{0x0A}
var path_sep = string(os.PathSeparator)

var writers = make(map[string]*bufio.Writer)

// fake type to specify interface :(
type Output int;

func (Output)Write(event *core.Event) {
	output_path := strings.Join([]string{
		core.ConfOutputFilesDir,
		MakeFilenameSafe(event.Source_host),
		MakePathSafe(event.Source_path)}, path_sep)

	writer, ok := writers[output_path]
	if !ok {
		err := os.MkdirAll(path.Dir(output_path), 0733)
		if err != nil {
			panic(err.Error())
		}

		file, err := os.OpenFile(output_path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0622)
		if err != nil {
			panic(err.Error())
		}

		writer = bufio.NewWriter(file)
		writers[output_path] = writer
	}

	_, err := writer.Write([]byte(event.Message))
	if err != nil {
		panic(err.Error())
	}

	_, err = writer.Write(new_line)
	if err != nil {
		panic(err.Error())
	}

	err = writer.Flush()
	if err != nil {
		panic(err.Error())
	}
}
