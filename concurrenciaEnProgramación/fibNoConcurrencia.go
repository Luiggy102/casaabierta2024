// funcion Fibonacci sin concurrencia
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	var nums int
	if len(os.Args) > 1 {
		numsInt, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		nums = numsInt
	} else {
		nums = 10
	}
	fmt.Println("Cantidad de números:", nums)
	var arr []int
	// llenar arreglo
	for i := 0; i < nums; i++ {
		arr = append(arr, 38+rand.Int()%5)
		// arr[i] = 38 + rand.Int()%5
	}
	fmt.Println(arr)     // mostrar arreglo
	inicio := time.Now() // tiempo

	// calcular y mostrar
	fmt.Println("número\tfibonacci\ttiempo")
	for i := 0; i < nums; i++ {
		fibNum := Fibonacci(arr[i])
		fmt.Printf("%d\t%d\t%v\n", arr[i], fibNum, time.Since(inicio))
	}
	fmt.Println("Tomó:", time.Since(inicio))
}
