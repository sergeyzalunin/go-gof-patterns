package helper

import (
	"io/ioutil"
	"os"
)

// GetPrintOutput is a callback function that can catch stdout stream and return 
func GetPrintOutput(testingFunc func()) string {
	rescueStdout := os.Stdout
	reaader, writer, _ := os.Pipe()
	os.Stdout = writer

	testingFunc()

	writer.Close()
	out, _ := ioutil.ReadAll(reaader)
	os.Stdout = rescueStdout
	return string(out)
}