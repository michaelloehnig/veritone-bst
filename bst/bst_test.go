package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeepestNodes(t *testing.T) {
	testTable := []struct {
		input []int

		expectedValue []int
		expectedDepth int
	}{
		{input: []int{12, 11, 90, 82, 7, 9}, expectedValue: []int{9}, expectedDepth: 3},
		{input: []int{12, 11, 90, 82, 7, 9, 6}, expectedValue: []int{6, 9}, expectedDepth: 3},
		{input: []int{12, 11, 90, 82, 85, 7, 75, 9, 6}, expectedValue: []int{6, 9, 75, 85}, expectedDepth: 3},
		{input: []int{12}, expectedValue: []int{12}, expectedDepth: 0},
	}

	for testNumber, test := range testTable {
		tree := FromSlice(test.input)
		actualValue, actualDepth, err := tree.DeepestNodes()
		if err != nil {
			t.Errorf("unexptected error %v", err)
		}
		if !assert.Equal(t, test.expectedDepth, actualDepth, "testcase #%v failed expected depth to be equal", testNumber) {
			continue
		}
		if !assert.Equal(t, test.expectedValue, actualValue, "testcase #%v failed expected values to be equal", testNumber) {
			continue
		}
	}
}

func TestFind(t *testing.T) {
	testTable := []struct {
		inputSlice []int
		find       int

		expected bool
	}{
		{inputSlice: []int{12, 11, 90, 82, 7, 9}, find: 9, expected: true},
		{inputSlice: []int{12, 11, 90, 82, 7, 9}, find: 12, expected: true},
		{inputSlice: []int{12, 11, 90, 82, 7, 9}, find: 11, expected: true},
		{inputSlice: []int{12, 11, 90, 82, 7, 9}, find: 90, expected: true},
		{inputSlice: []int{12, 11, 90, 82, 7, 9}, find: 99, expected: false},
	}

	for testNumber, test := range testTable {
		tree := FromSlice(test.inputSlice)
		_, actual := tree.Find(test.find)
		if !assert.Equal(t, test.expected, actual, "testcase #%v failded expected to be equal", testNumber) {
			t.Fatal()
		}
	}
}

func TestFromSlice(t *testing.T) {
	testTable := []struct {
		inputSlice []int

		expected []int
	}{
		{inputSlice: []int{12, 11, 90, 82, 7, 9}, expected: []int{12, 11, 90, 7, 82, 9}},
	}

	for testNumber, test := range testTable {
		tree := FromSlice(test.inputSlice)
		slice := tree.ToSlice()
		if !assert.Equal(t, test.expected, slice, "testcase #%v failded expected to be equal", testNumber) {
			t.Fatal()
		}
	}
}

func TestDelete(t *testing.T) {
	testTable := []struct {
		inputSlice []int
		toDelete   int

		expected []int
	}{
		{inputSlice: []int{12, 11, 90, 82, 7, 9}, toDelete: 9, expected: []int{12, 11, 90, 7, 82}}, // no children were affected
		{inputSlice: []int{12, 11, 90, 82, 7, 9}, toDelete: 90, expected: []int{12, 11, 82, 7, 9}}, // ensure we keep 82
		{inputSlice: []int{12, 11, 90, 82, 7, 9, 6, 8, 10}, toDelete: 9, expected: []int{12, 11, 90, 7, 82, 6, 10, 8}},
		{inputSlice: []int{20, 11, 7, 15, 13, 19, 12, 14, 18}, toDelete: 15, expected: []int{20, 11, 7, 19, 18, 13, 12, 14}},
		{inputSlice: []int{20, 11, 7, 15, 13, 19, 12, 14, 18}, toDelete: 19, expected: []int{20, 11, 7, 15, 13, 18, 12, 14}},
	}

	for testNumber, test := range testTable {
		tree := FromSlice(test.inputSlice)
		tree.Delete(test.toDelete)
		actual := tree.ToSlice()
		if !assert.Equal(t, test.expected, actual, "testcase #%v failded expected to be equal", testNumber) {
			t.Fatal()
		}
	}
}
