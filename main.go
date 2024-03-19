package main

import (
	"fmt"
	"time"
)

const (
	testItems = 10
)

func main() {

	pivot := time.Now()

	testNativeTimeline := NativeTimeline{
		Items: []NativeObject{},
	}

	for i := 0; i < testItems; i++ {
		testNativeTimeline.Items = append(testNativeTimeline.Items, NativeObject{
			ID:        fmt.Sprintf("id-%d", i),
			Content:   fmt.Sprintf("content-%d", i),
			Timestamp: pivot.Add(time.Duration(i) * time.Hour),
		})
	}

	for i := 0; i < testItems; i++ {
		testNativeTimeline.Items = append(testNativeTimeline.Items, NativeObject{
			ID:        fmt.Sprintf("id-%d", testItems+i),
			Content:   fmt.Sprintf("content-%d", testItems+i),
			Timestamp: pivot.Add(time.Duration(i)*time.Hour + time.Hour*30),
		})
	}

	stp, err := RenderTimeline(testNativeTimeline, 60*60*6) // 6 hours
	if err != nil {
		fmt.Println(err)
	}

	jsonPrint(stp)
}
