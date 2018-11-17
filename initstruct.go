package main

import "sync"
import "time"
import "fmt"

type SyncMap struct {
        lock *sync.RWMutex
        hm map[int]string
}

func NewSyncMap() *SyncMap {
        return &SyncMap{lock: new(sync.RWMutex), hm: make(map[int]string)}
}

func (m *SyncMap) Put (k int, v string) {
        m.lock.Lock()
        defer m.lock.Unlock()
        m.hm[k] = v
}

func main() {
  sm := NewSyncMap()
  for j:=0;j<5;j++ {
    sm.Put(j, "no")
  }
  time.Sleep(time.Second)
  sm2 := NewSyncMap()
  for i:=0;i<10;i++{
    go sm2.Put(i, "go")
  }
  time.Sleep(time.Second)
  fmt.Println(sm)
  fmt.Println(sm2)
}
