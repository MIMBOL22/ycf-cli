package internal

import "sync"

func SmartFor[T any](list []*T, f func(*T)) {
	var wg sync.WaitGroup
	for _, el := range list {
		wg.Add(1)
		go func(el *T) {
			defer wg.Done()

			f(el)
		}(el)
	}
	wg.Wait()
}
