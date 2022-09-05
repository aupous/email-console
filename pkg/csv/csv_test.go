package csv

import (
	"email-console/pkg/array"
	"os"
	"testing"
)

var (
	testPath  = "test.csv"
	testLines = [][]string{
		{"1", "2"},
		{"3", "4"},
	}
)

func TestWriteToFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		lines    [][]string
		expected string
	}{
		{
			name:     "write success",
			path:     testPath,
			lines:    testLines,
			expected: "1,2\n3,4\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteToFile(tt.path, tt.lines)
			defer os.Remove(tt.path)
			if err != nil {
				t.Errorf("expect success but got error %s", err.Error())
			}
			fileData, _ := os.ReadFile(tt.path)
			actual := string(fileData)
			if actual != tt.expected {
				t.Errorf("expect file data is %s but got %s", tt.expected, actual)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		lines [][]string
	}{
		{
			name:  "write success",
			path:  testPath,
			lines: testLines,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteToFile(tt.path, tt.lines)
			defer os.Remove(tt.path)
			if err != nil {
				t.Error("error when write lines to file")
				return
			}
			actualLines, err := ReadFile(tt.path)
			equal := len(tt.lines) == len(actualLines)
			for ind, line := range tt.lines {
				if !array.StringArrayEqual(line, actualLines[ind]) {
					equal = false
				}
			}
			if !equal {
				t.Errorf("expected read data to be %v but got %v", tt.lines, actualLines)
			}
		})
	}
}
