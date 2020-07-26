package bubblesort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	//initialization
	elements := []int{9, 5, 8, 4, 0}

	fmt.Println(elements)

	//execution
	BubbleSort(elements)
	fmt.Println(elements)

	//validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}
