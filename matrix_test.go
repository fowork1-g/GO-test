package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareMatrixConfPositiv(t *testing.T) {
	testDefaultUnicFlag := DEFAULT_FLAG_UNIC
	testUnicFlagTrue := true
	testUnicFlagFalse := false
	defaultNumFrom := 0

	testTable := []struct {
		data     MatrixConf
		expected *MatrixConf
	}{
		{
			data:     MatrixConf{x: 5, y: 5},
			expected: &MatrixConf{x: 5, y: 5, numFrom: defaultNumFrom, numTo: DEFAULT_NUM_TO, shouldUnic: &testDefaultUnicFlag},
		}, {
			data:     MatrixConf{x: 5, y: 5, numTo: 100},
			expected: &MatrixConf{x: 5, y: 5, numFrom: defaultNumFrom, numTo: 100, shouldUnic: &testDefaultUnicFlag},
		}, {
			data:     MatrixConf{x: 5, y: 5, numFrom: -100, numTo: 10},
			expected: &MatrixConf{x: 5, y: 5, numFrom: -100, numTo: 10, shouldUnic: &testDefaultUnicFlag},
		}, {
			data:     MatrixConf{x: 5, y: 5, numFrom: 100, numTo: -100},
			expected: &MatrixConf{x: 5, y: 5, numFrom: -100, numTo: 100, shouldUnic: &testDefaultUnicFlag},
		}, {
			data:     MatrixConf{x: 5, y: 5, shouldUnic: &testUnicFlagTrue},
			expected: &MatrixConf{x: 5, y: 5, numFrom: defaultNumFrom, numTo: DEFAULT_NUM_TO, shouldUnic: &testUnicFlagTrue},
		}, {
			data:     MatrixConf{x: 5, y: 5, shouldUnic: &testUnicFlagFalse},
			expected: &MatrixConf{x: 5, y: 5, numFrom: defaultNumFrom, numTo: DEFAULT_NUM_TO, shouldUnic: &testUnicFlagFalse},
		}, {
			data:     MatrixConf{x: 5, y: 5, numTo: 10, shouldUnic: &testUnicFlagFalse},
			expected: &MatrixConf{x: 5, y: 5, numFrom: defaultNumFrom, numTo: 10, shouldUnic: &testUnicFlagFalse},
		},
	}

	for _, testCase := range testTable {
		result, err := PrepareMatrixConf(&testCase.data)

		assert.Nil(t, err)
		assert.Equal(t, testCase.expected, result)
	}
}

func TestPrepareMatrixConfNegative(t *testing.T) {
	testTable := []struct {
		data MatrixConf
	}{
		{
			data: MatrixConf{x: 599, y: 599},
		}, {
			data: MatrixConf{x: 5, y: 5, numTo: 10},
		}, {
			data: MatrixConf{x: -5, y: 5},
		}, {
			data: MatrixConf{x: 5, y: -5},
		}, {
			data: MatrixConf{x: 0, y: 5},
		}, {
			data: MatrixConf{x: 5, y: 0},
		},
	}

	for _, testCase := range testTable {
		_, err := PrepareMatrixConf(&testCase.data)
		assert.NotNil(t, err)
	}
}

func TestMakeNumberGenerator(t *testing.T) {
	testTable := []struct {
		data     MatrixConf
		mem      map[int]bool
		expected int
	}{
		{
			data:     MatrixConf{numFrom: 0, numTo: 1},
			expected: 0,
		}, {
			data:     MatrixConf{numFrom: -1, numTo: 0},
			expected: -1,
		}, {
			data:     MatrixConf{numFrom: 1, numTo: 2},
			expected: 1,
		}, {
			data:     MatrixConf{numFrom: 1, numTo: 3},
			mem:      map[int]bool{1: true},
			expected: 2,
		},
	}

	for _, testCase := range testTable {
		gen := makeNumberGenerator(testCase.data, testCase.mem)

		assert.Equal(t, reflect.Func, reflect.TypeOf(gen).Kind())

		result := gen()

		assert.Equal(t, testCase.expected, result)
	}
}

func TestMakeMatrix(t *testing.T) {
	testUnicFlagTrue := true
	testUnicFlagFalse := false

	testTable := []struct {
		data         MatrixConf
		expected     [][]int
		isValuesDiff bool
	}{
		{
			data: MatrixConf{x: 1, y: 2, numTo: 1, shouldUnic: &testUnicFlagFalse},
			expected: [][]int{
				{0, 0},
			},
		}, {
			data: MatrixConf{x: 1, y: 2, numFrom: 99, numTo: 100, shouldUnic: &testUnicFlagFalse},
			expected: [][]int{
				{99, 99},
			},
		}, {
			data: MatrixConf{x: 1, y: 2, numFrom: 1, numTo: 3, shouldUnic: &testUnicFlagTrue},
			expected: [][]int{
				{1, 2},
			},
			isValuesDiff: true,
		},
	}

	for _, testCase := range testTable {
		result := MakeMatrix(testCase.data)

		//assert.Nil(t, err)
		if testCase.isValuesDiff {
			assert.NotEqual(t, result[0][0], result[0][1])
		} else {
			assert.Equal(t, testCase.expected, result)
		}
	}
}
