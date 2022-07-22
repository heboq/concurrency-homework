package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan struct{}, pool)
	wg.Add(int(n))
	for i := int64(0); i < n; i++ {
		go func(i int64) {
			defer wg.Done()
			sem <- struct{}{}
			user := getOne(i)
			mu.Lock()
			res = append(res, user)
			mu.Unlock()
			<-sem
		}(i)
	}
	wg.Wait()
	return
}
