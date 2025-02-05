// Description: 互斥锁

package main

import "sync"

var mu sync.Mutex

func main() {
	mu.Lock()
	defer mu.Unlock()
}