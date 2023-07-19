package main

import (
	"log"
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestModelComputerClub(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name       string
		inputFile  string
		outputFile string
	}{
		// the table itself
		{"BasicCase", "testcases/basicCase.txt", "testcases/basicCaseOutput.txt"},
		{"ClientUnknown", "testcases/ClientUnknown.txt", "testcases/ClientUnknownOutput.txt"},
		{"PlaceIsBusy", "testcases/PlaceIsBusy.txt", "testcases/PlaceIsBusyOutput.txt"},
		{"TimeFormatError", "testcases/TimeFormatError.txt", "testcases/TimeFormatErrorOutput.txt"},
		{"TooMuchArgumentsError", "testcases/TooMuchArgumetsError.txt", "testcases/TooMuchAggumentsErrorOutput.txt"},
		{"YouShallNotPass", "testcases/YouShallNotPass.txt", "testcases/YouShallNotPassOutput.txt"},
		{"NotEnoughArguments", "testcases/NotEnoghArguments.txt", "testcases/NotEnoughArgumentsOutput.txt"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ModelComputerClub(tt.inputFile)

			f, err := os.ReadFile(tt.outputFile)
			if err != nil {
				log.Fatal(err)
			}
			want := string(f)

			if runtime.GOOS == "windows" {
				ans = strings.TrimRight(ans, "\r\n")
				want = strings.TrimRight(want, "\r\n")
			} else {
				ans = strings.TrimRight(ans, "\n")
				want = strings.TrimRight(want, "\n")
			}

			if strings.Compare(ans, want) != 0 {
				t.Errorf("got:\n%s\n\nwant:\n%s\n", ans, want)
			}
		})
	}
}
