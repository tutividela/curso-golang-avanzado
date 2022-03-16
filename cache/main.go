package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock *sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
		lock: &sync.Mutex{},
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exist := m.cache[key]
	m.lock.Unlock()
	if !exist {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		
		m.cache[key] = result
		m.lock.Unlock()
	}
	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	var wg sync.WaitGroup
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}

	for _, v := range fibo {
		wg.Add(1)
		go func (index int)  {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%d ,%s,%d \n",index,time.Since(start),value)
		}(v)
	}
	wg.Wait()
}