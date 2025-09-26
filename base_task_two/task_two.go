package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 指针
// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值

func addTen(num *int) {
	*num += 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func multiplyByTwo(nums *[]int) {
	for index, _ := range *nums {
		(*nums)[index] *= 2
	}
}

// Goroutine
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func GoroutinePrint() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i += 2 {
			fmt.Printf("Odd: %d\n", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i < 10; i += 2 {
			fmt.Printf("Even: %d\n", i)
		}
	}()
	wg.Wait()
	println("goroutine end")
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 任务1
func taskOne(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	time.Sleep(1 * time.Second)
	end := time.Now()
	fmt.Printf("Task One took %v\n", end.Sub(start))
}

// 任务2
func taskTwo(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	fmt.Printf("Task Two took %v\n", end.Sub(start))
}

// 任务3
func taskThree(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	time.Sleep(3 * time.Second)
	end := time.Now()
	fmt.Printf("Task Three took %v\n", end.Sub(start))
}

// 面向对象
// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
// 实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Aera() float64
	Perimeter() float64
}

// 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Aera() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Aera() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 问题2

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	employeeID int
	person     Person
}

func (em Employee) PrintInfo() {

	fmt.Printf("Employee ID: %d, Name: %s, Age: %d\n", em.employeeID, em.person.Name, em.person.Age)
}

// Channel
func generateData(ch *chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 100; i++ {
		*ch <- i
		// time.Sleep(1 * time.Second)
	}
	close(*ch)
}

func readData(ch *chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		num, ok := <-*ch
		if ok {
			fmt.Printf("num is %d\n", num)
		} else {
			break
		}
	}

}

// 锁机制
var mutex sync.Mutex

func dealMutex(num *int, wg *sync.WaitGroup) {

	defer wg.Done()
	mutex.Lock()
	for i := 0; i < 1000; i++ {

		*num++
	}
	mutex.Unlock()
}

func dealAtomic(num *int32, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 0; i < 1000; i++ {

		atomic.AddInt32(num, 1)
	}
}

func main() {
	// 指针
	// value := 5
	// addTen(&value)
	// fmt.Println("Modified value:", value)

	// nums := []int{1, 2, 3, 4, 5}
	// multiplyByTwo(&nums)
	// fmt.Println("Modified slice:", nums)

	// Goroutine
	// GoroutinePrint()
	// 记住初始化函数作为slice的对象时需要加上参数
	// funcList := []func(wg *sync.WaitGroup){taskOne, taskTwo, taskThree}
	// var wg sync.WaitGroup

	// for _, f := range funcList {
	// 	wg.Add(1)
	// 	go f(&wg)
	// }
	// wg.Wait()
	// println("All tasks completed")

	// 面向对象

	// rect := Rectangle{Width: 5, Height: 10}
	// circ := Circle{Radius: 7}

	// var s Shape
	// s = rect
	// fmt.Printf("Rectangle Area: %v, Perimeter: %v\n", s.Aera(), s.Perimeter())
	// s = circ
	// fmt.Printf("Circle Area: %v, Perimeter: %v\n", s.Aera(), s.Perimeter())

	// user := Employee{employeeID: 12345, person: Person{Name: "Alice", Age: 30}}

	// user.PrintInfo()

	// Channel
	// ch := make(chan int, 100)
	// var wg sync.WaitGroup

	// wg.Add(2)
	// // 产生数据
	// go generateData(&ch, &wg)
	// time.Sleep(1 * time.Second)
	// // 读取数据
	// go readData(&ch, &wg)
	// wg.Wait()

	// fmt.Printf(" finish data ")

	// 锁机制
	var wg sync.WaitGroup
	var num int32
	num = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go dealAtomic(&num, &wg)
	}

	wg.Wait() //
	fmt.Printf("final num %d\n", num)
}
