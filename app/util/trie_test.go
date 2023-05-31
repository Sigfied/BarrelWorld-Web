package util

import (
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T) {
	array := []string{
		"apple",
		"banana",
		"cherry",
		"grape",
		"orange",
		"pineapple",
		"dadada",
		"中文",
		"引文",
		"大大撒打发发中文发开孔处",
	}

	substring := "中文"

	result := FuzzySearch(substring, array)

	fmt.Println(result) // Output: [banana, orange]
}
