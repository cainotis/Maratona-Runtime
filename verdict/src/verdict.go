package verdict

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/maratona-run-time/Maratona-Runtime/comparator/src"
	"github.com/maratona-run-time/Maratona-Runtime/executor/src"
)

func Verdict(timeout float32, executablePath string, inputFilesFolder string, outputFilesFolder string, result chan<- string) {
	res := executor.Execute(executablePath, inputFilesFolder, timeout)

	for _, executionResult := range res {
		_, fileName := path.Split(executionResult[0])
		testResult := executionResult[1]
		programOutput := executionResult[2]
		switch testResult {
		case "TLE":
			result <- "TLE"
			return
		case "RTE":
			result <- "RTE"
			return
		case "OK":
			testName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			outputFileName := filepath.Join(outputFilesFolder, testName+".out")
			expectedData, err := ioutil.ReadFile(outputFileName)
			if err != nil {
				fmt.Println(err)
			}
			expectedOutput := string(expectedData)
			if !comparator.Compare(expectedOutput, programOutput) {
				result <- "WA"
				return
			}
		}
	}
	result <- "AC"
	return
}