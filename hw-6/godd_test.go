package main

import (
	"bytes"
	"io"
	"testing"
)

var (
	testdata  = []byte("1234567890")
	TestCases = []struct {
		testCaseID     string // Name of testcase
		limit          int64  // Number of bytes to copy
		offset         int64
		source         []byte // Offset in the source file
		expectedResult string // Expected result data
		expectedError  error  // Expected error
	}{
		{
			testCaseID:     "Test 1: copy all ",
			limit:          0,
			offset:         0,
			source:         testdata,
			expectedResult: "1234567890",
			expectedError:  nil,
		},
		{
			testCaseID:     "Test 2: copy half from begining",
			limit:          5,
			offset:         0,
			source:         testdata,
			expectedResult: "12345",
			expectedError:  nil,
		},
		{
			testCaseID:     "Test 3: copy half from middle",
			limit:          5,
			offset:         5,
			source:         testdata,
			expectedResult: "67890",
			expectedError:  nil,
		},
		{
			testCaseID:     "Test 4: limit is greater than size of source file",
			limit:          15,
			offset:         0,
			source:         testdata,
			expectedResult: "1234567890",
			expectedError:  nil,
		},
		{
			testCaseID:     "Test 5: offset is greater than size of source file",
			limit:          0,
			offset:         15,
			source:         testdata,
			expectedResult: "",
			expectedError:  nil,
		},
		{
			testCaseID:     "Test 6: Negative offset",
			limit:          0,
			offset:         -5,
			source:         testdata,
			expectedResult: "67890",
			expectedError:  nil,
		},
		{
			testCaseID:     "Test 7: Negative offset and some limit",
			limit:          2,
			offset:         -5,
			source:         testdata,
			expectedResult: "67",
			expectedError:  nil,
		},
	}
)

func TestGodd(t *testing.T) {

	for _, test := range TestCases {

		reader := bytes.NewReader(test.source)
		writer := bytes.NewBuffer([]byte{})

		err := copier(reader, writer, int64(len(test.source)), test.offset, test.limit)

		if err != nil {
			if err != io.EOF {
				t.Error("Failed test case", test.testCaseID, err)
			}
		}

		if result := writer.String(); result != test.expectedResult {
			t.Error("Copied data does not equals the expected result", test.testCaseID, result, test.expectedResult)
		}
	}

}
