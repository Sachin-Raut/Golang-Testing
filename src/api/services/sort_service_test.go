package services

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	//initialisation
	elements := []int{7, 9, 2, 5, 0, 6}
	fmt.Println(elements)

	//execution
	Sort(elements)
	fmt.Println(elements)

	//validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}
