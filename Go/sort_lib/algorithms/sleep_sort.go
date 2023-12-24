package algorithms

import (
	"sync"
	"time"
)

func SleepSort(numbers *[]int) {
	output := make(chan int)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, num := range *numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			time.Sleep(time.Duration(num) * time.
				Second)
			mutex.Lock()
			output <- num
			mutex.Unlock()
		}(num)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	var result []int
	for num := range output {
		result = append(result, num)
	}
	*numbers = result
}
