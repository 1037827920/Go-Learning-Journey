// Description: 让一个结构体可以被并发访问

package main

import "sync"

type Visted struct {
	mu sync.Mutex
	visited map[string]int
}

func (v *Visted) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {
	
}