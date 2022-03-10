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
