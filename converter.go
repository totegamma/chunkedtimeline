package main

import (
	"container/heap"
	"fmt"
	"sort"
	"time"
)

func RenderTimeline(native NativeTimeline, chunkSize int) (STP, error) {

	datas := map[string]NativeObject{}
	chunks := map[string]Chunk{}
	iterators := map[string]Iterator{}

	firstTs := time.Time{}
	lastTs := time.Time{}

	keyHeap := &IntHeap{}
	heap.Init(keyHeap)

	buckets := make(map[string][]NativeObject)
	for _, item := range native.Items {

		datas[item.ID] = item
		bucket := item.Timestamp.Unix() / int64(chunkSize)
		bucketKey := fmt.Sprintf("%d", bucket)
		heap.Push(keyHeap, bucket)
		if _, ok := buckets[bucketKey]; !ok {
			buckets[bucketKey] = []NativeObject{}
		}
		buckets[bucketKey] = append(buckets[bucketKey], item)

		if firstTs.IsZero() || item.Timestamp.Before(firstTs) {
			firstTs = item.Timestamp
		}
		if lastTs.IsZero() || item.Timestamp.After(lastTs) {
			lastTs = item.Timestamp
		}
	}

	for key, value := range buckets {

		sort.Slice(value, func(i, j int) bool {
			return value[i].Timestamp.Before(value[j].Timestamp)
		})

		items := []ChunkItem{}

		for _, item := range value {
			items = append(items, ChunkItem{
				ObjectRef: item.ID,
				TimeStamp: item.Timestamp,
			})
		}

		chunks[key] = Chunk{
			Key:   key,
			Items: items,
		}
	}

	firstChunk := firstTs.Unix() / int64(chunkSize)
	lastChunk := lastTs.Unix() / int64(chunkSize)
	numOfChunks := int(lastChunk - firstChunk + 1)

	for i := 0; i < numOfChunks; i++ {
		chunkItr := firstChunk + int64(i)
		chunkID := fmt.Sprintf("%d", chunkItr)

		nextKey := ""
		prevKey := ""

		nextIndex := -1
		for j := 0; j < keyHeap.Len()-1; j++ { // TODO: should be binary search
			if (*keyHeap)[j] <= chunkItr && chunkItr < (*keyHeap)[j+1] {
				nextIndex = j
				break
			}
		}
		if (*keyHeap)[keyHeap.Len()-1] == chunkItr {
			nextIndex = keyHeap.Len() - 1
		}

		if nextIndex != -1 {
			nextKey = fmt.Sprintf("%d", (*keyHeap)[nextIndex])
		}

		prevIndex := -1
		for j := 1; j < keyHeap.Len(); j++ {
			if (*keyHeap)[j-1] < chunkItr && chunkItr <= (*keyHeap)[j] {
				prevIndex = j
				break
			}
		}
		if (*keyHeap)[0] == chunkItr {
			prevIndex = 0
		}

		if prevIndex != -1 {
			prevKey = fmt.Sprintf("%d", (*keyHeap)[prevIndex])
		}

		iterators[chunkID] = Iterator{
			Next: nextKey,
			Prev: prevKey,
		}
	}

	return STP{
		Timeline: Timeline{
			IteratorTemplate: "iterator_template",
			ChunkTemplate:    "chunk_template",
			ChunkSize:        chunkSize,
			FirstItem:        firstTs,
			LastItem:         lastTs,
		},
		Iterators: iterators,
		Chunks:    chunks,
		Datas:     datas,
	}, nil
}

func QueryTimeline(url string) ([]NativeObject, error) {
	return nil, nil
}
