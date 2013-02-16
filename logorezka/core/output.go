package core

import (
	"bufio"
	"os"
	"path"
	"strings"
)

var NEW_LINE = []byte{0x0A}
var PATH_SEP = string(os.PathSeparator)

var writers = make(map[string]*bufio.Writer)

func OutputEvent(event *Event) {
	output_path := strings.Join([]string{
		ConfOutputFilesDir,
		MakeFilenameSafe(event._source_host),
		MakePathSafe(event._source_path)}, PATH_SEP)

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

	_, err := writer.Write(event._message)
	if err != nil {
		panic(err.Error())
	}

	_, err = writer.Write(NEW_LINE)
	if err != nil {
		panic(err.Error())
	}

	err = writer.Flush()
	if err != nil {
		panic(err.Error())
	}
}
