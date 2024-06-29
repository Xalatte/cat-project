package main

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestPrintFiles(t *testing.T) {
	tCases := []struct {
		name           string
		reader         io.Reader
		writer         io.Writer
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "Empty file",
			reader:         strings.NewReader(""),
			writer:         &bytes.Buffer{},
			expectedOutput: "",
			expectedError:  nil,
		},
		{
			name:           "nil reader",
			reader:         nil,
			writer:         &bytes.Buffer{},
			expectedOutput: "",
			expectedError:  ReaderIsNilErr,
		},
		{
			name:           "nil writer",
			reader:         nil,
			writer:         &bytes.Buffer{},
			expectedOutput: "",
			expectedError:  WriterIsNilErr,
		},
	}
	for _, tCase := range tCases {
		t.Run(tCase.name, func(t *testing.T) {
			err := printFiles(tCase.reader, tCase.writer)
			if tCase.expectedError != nil {
				if !errors.Is(err, tCase.expectedError) {
					t.Fail()
				}
			}

			if tCase.writer != nil {
				actualOutput := tCase.writer.(*bytes.Buffer).String()
				if actualOutput != tCase.expectedOutput {
					t.Fail()
				}
			}
		})
	}
}
