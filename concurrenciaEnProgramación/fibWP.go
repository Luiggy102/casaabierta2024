// concurrencia con go usando workers
package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"sync"
	"time"
)

type Operation struct {
	worker    int
	num       string
	fibonacci int
	took      time.Duration
}

func (o Operation) String() string {
	return fmt.Sprintf("Worker id %d\tNum: %s\t\tFib: %d\t\ttiempo: %v", o.worker, o.num, o.fibonacci, o.took)
}

func newOperation(workerID int, num string, fib int, took time.Duration) *Operation {
	return &Operation{workerID, num, fib, took}
}

func worker(id int, jobs <-chan string, results chan<- Operation, cache *Memory) {
	for j := range jobs {
		jInt, err := strconv.Atoi(j)
		if err != nil {
			continue
		}
		// result := Fibonacci(jInt)
		// results <- *newOperation(id, j, result)

		workerStart := time.Now()
		result, err := cache.Get(jInt)
		if err != nil {
			continue
		}
		results <- *newOperation(id, j, result.(int), time.Since(workerStart))
	}
}

func main() {
	// analisis de redimiento (cpu)
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// ...
	start := time.Now()
	cache := &Memory{GetFibonacci, map[int]struct {
		fib interface{}
		err error
	}{}, sync.Mutex{}}

	const numTask = 1_000_000
	var tasks [numTask]string

	for i := 0; i < numTask; i++ {
		tasks[i] = strconv.Itoa(38 + rand.Int()%3)
	}

	// tasks := os.Args[1:]
	// numTask := len(os.Args[1:])
	jobs := make(chan string, numTask)
	results := make(chan Operation, numTask)

	// numWorkersStr := os.Args[1]
	// numWorkers, err := strconv.Atoi(numWorkersStr)
	// if err != nil {
	// 	panic("Valor inválido número de workers")
	// }

	numWorkers := runtime.NumCPU()
	runtime.GOMAXPROCS(numWorkers) // numero de cpus que se usarán en el runtime

	fmt.Println("Num CPU:", numWorkers)
	fmt.Println("Cantidad de números:", len(tasks))

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results, cache)
	}

	for j := 0; j < len(tasks); j++ {
		jobs <- tasks[j]
	}
	close(jobs)

	for i := 1; i <= numTask; i++ {
		fmt.Println(i, "\t", <-results)
	}
	fmt.Println("Tomó:", time.Since(start))
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// / sistema de cache
func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

type Memory struct {
	function func(int) (interface{}, error)
	cache    map[int]struct {
		fib interface{}
		err error
	}
	lock sync.Mutex
}

func (m *Memory) Get(n int) (interface{}, error) {
	m.lock.Lock()
	res, ok := m.cache[n]
	m.lock.Unlock()
	if !ok {
		m.lock.Lock()
		res.fib, res.err = m.function(n)
		m.cache[n] = res
		m.lock.Unlock()
	}
	return res.fib, res.err
}
