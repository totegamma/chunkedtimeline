package main

import (
	"time"
)

type STP struct {
	Timeline  Timeline                `json:"timeline"`
	Iterators map[string]Iterator     `json:"iterators"`
	Chunks    map[string]Chunk        `json:"chunks"`
	Datas     map[string]NativeObject `json:"datas"`
}

type Timeline struct {
	IteratorTemplate string    `json:"iterator_template"`
	ChunkTemplate    string    `json:"chunk_template"`
	ChunkSize        int       `json:"chunksize"`
	FirstItem        time.Time `json:"firstitem"`
	LastItem         time.Time `json:"lastitem"`
}

type Chunk struct {
	Key   string      `json:"key"`
	Items []ChunkItem `json:"items"`
}

type Iterator struct {
	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

type ChunkItem struct {
	ObjectRef string    `json:"objectref"`
	TimeStamp time.Time `json:"timestamp"`
}

type NativeObject struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type NativeTimeline struct {
	Items []NativeObject `json:"items"`
}
