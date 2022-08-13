package batch

import (
	"sync"
	"time"
)

type User struct {
	ID int64
}

func getOne(id int64) User {
	time.Sleep(time.Millisecond * 100)
	return User{ID: id}
}

func GetBatch(n int64, pool int64) []User {
	var wg sync.WaitGroup
	var mu sync.Mutex
	res := make([]User, n)
	sem := make(chan struct{}, pool)
	wg.Add(int(n))
	for i := int64(0); i < n; i++ {
		go func(i int64) {
			defer wg.Done()
			sem <- struct{}{}
			user := getOne(i)
			mu.Lock()
			// // res = append(res, user)
			res[i] = user
			mu.Unlock()
			<-sem
		}(i)
	}
	wg.Wait()
	return res
}
