package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var file string

func find(topic string, docs []string, genFilesNum int) int {
	var found int
	for _, doc := range docs {
		items, err := read(doc, genFilesNum)
		if err != nil {
			continue
		}
		for _, item := range items {
			if strings.Contains(item.Description, topic) {
				found++
			}
		}
	}
	return found
}

// goroutine is equal to docs num
func findConcurrent(topic string, docs []string, genFilesNum int) int {
	var found int64
	var wg sync.WaitGroup
	for _, doc := range docs {
		wg.Add(1)
		go func(doc string) {
			items, err := read(doc, genFilesNum)
			if err != nil {
				return
			}
			lfound := 0
			for _, item := range items {
				if strings.Contains(item.Description, topic) {
					lfound++
				}
			}
			atomic.AddInt64(&found, int64(lfound))
			wg.Done()

		}(doc)
	}
	wg.Wait()
	return int(found)
}

// goroutine is equal to CPU.NUM
func findConcurrentEqualToCpu(topic string, docs []string, genFilesNum int, goroutines int) int {
	var found int64
	ch := make(chan string, len(docs))
	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			var lfound int
			for doc := range ch {
				items, err := read(string(doc), genFilesNum)
				if err != nil {
					fmt.Println("contiinue")
					continue
				}
				for _, item := range items {
					if strings.Contains(item.Description, topic) {
						lfound++
					}
				}
			}
			atomic.AddInt64(&found, int64(lfound))
			wg.Done()
		}()
	}
	wg.Wait()
	return int(found)
}

type item struct {
	Description string
}
type document struct {
	Items []item
}

func read(doc string, genFilesNum int) ([]item, error) {
	filesStr := genFileInfo(genFilesNum)
	time.Sleep(time.Millisecond) // Simulate blocking disk read.
	var d document
	if err := json.Unmarshal([]byte(filesStr), &d); err != nil {
		return nil, err
	}
	return d.Items, nil
}

func genFileInfo(n int) string {
	files := make([]item, 0, n)
	for i := 0; i < n; i++ {
		itemElem := item{
			Description: fmt.Sprintf("%d", i),
		}
		files = append(files, itemElem)
	}
	doc := document{
		Items: files,
	}
	filesStr, err := json.Marshal(doc)
	if err != nil {
		return ""
	}
	return string(filesStr)
}

//My goal in implementing the concurrent version was to control the number of Goroutines that are used to process the unknown number of documents. A pooling pattern where a channel is used to feed the pool of Goroutines was my choice.

func TestSucceed(t *testing.T) {
	topic := "4"
	docsNums := 2
	genFilesNum := 20
	docs := make([]string, docsNums)
	findCnt := find(topic, docs, genFilesNum)
	if findCnt != 4 {
		t.Errorf("findCnt should be 4")
	}
	findCnt = find(topic, docs, 10)
	if findCnt != 2 {
		t.Errorf("findCnt should be 2")
	}

	findCnt = findConcurrent(topic, docs, 10)
	if findCnt != 2 {
		t.Errorf("findCnt should be 2;real=%d", findCnt)
	}
	findCnt = findConcurrent(topic, docs, 30)
	if findCnt != 6 {
		t.Errorf("findCnt should be 6;real=%d", findCnt)
	}
	findCnt = findConcurrentEqualToCpu(topic, docs, 30, 4)
	if findCnt != 6 {
		t.Errorf("findCnt should be 6;real=%d", findCnt)
	}
}
func BenchmarkSequential(b *testing.B) {
	genFilesNum := 10000
	docsNums := 200
	docs := make([]string, docsNums)
	for i := 0; i < b.N; i++ {
		find("4", docs, genFilesNum)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	genFilesNum := 10000
	docsNums := 200
	docs := make([]string, docsNums)
	for i := 0; i < b.N; i++ {
		findConcurrent("4", docs, genFilesNum)
	}
}
func BenchmarkConcurrentEqualToGpu(b *testing.B) {
	genFilesNum := 10000
	docsNums := 200
	docs := make([]string, docsNums)
	for i := 0; i < b.N; i++ {
		findConcurrentEqualToCpu("4", docs, genFilesNum, runtime.NumCPU())
	}
}
