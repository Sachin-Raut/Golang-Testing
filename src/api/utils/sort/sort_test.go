package sort

import (
	"fmt"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	//initialization
	elements := []int{9, 5, 8, 4, 0}
	fmt.Println(elements)

	/*
		1.if we remove "keepWorking = false", then our test will run forever
		2.so let make sure that the test will run only for 5 seconds
		3.we will use channels for that
	*/

	timeoutChan := make(chan bool, 1)

	defer close(timeoutChan)

	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(5 * time.Second)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		t.Error("BubbleSort took more than 5 seconds")
		return
	}

	//execution
	// BubbleSort(elements)
	// fmt.Println(elements)

	//validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}
