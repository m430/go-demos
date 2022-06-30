package chapter19

import (
	"fmt"
	"testing"
	"time"
)

// 1. 小心迭代变量的重用
func TestForRange(t *testing.T) {
	var m = []int{1, 2, 3, 4, 5}
	for i, v := range m {
		// 这里goroutine引用了i, v，当开启goroutine的时候，主线程会继续执行，共享访问的变量最终打印出来就是主线程循环
		// 执行解除的值
		go func() {
			time.Sleep(time.Second * 1)
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 5)
}

// 修正上面的例子
func TestOkForRange(t *testing.T) {
	var m = []int{1, 2, 3, 4, 5}
	for i, v := range m {
		// 这次循环每次开启的goroutine访问的变量是函数参数值拷贝的变量，会按照预期打印。
		// 不过打印顺序是随机的，是由goroutine的调度决定的
		go func(i, v int) {
			time.Sleep(time.Second * 1)
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 5)
}

// 2. 注意参与迭代的是range表达式的副本
func TestRangeExpression(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("arrayRangeExpression result:")
	fmt.Println("a = ", a)

	// 这里的a是上面a的副本，并不是a本身
	for i, v := range a {
		if i == 0 {
			// 这里改变的确实是a本身的值
			a[1] = 12
			a[2] = 13
		}
		// 由于v是a副本迭代值，所以是a刚开始的值
		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

// 解决方法1：如果希望循环改变数组值的话，可以使用指针
func TestRangeExpressionByPoint(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("arrayPointRangeExpression result:")
	fmt.Println("a = ", a)

	// 这里的遍历的是a数组的指针的副本，指针的副本依然指向数组a本身
	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		// 所以v的值是迭代的&a，所以指向的依然是原来a数组
		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

// 解决方法2：如果希望循环改变数组值的话，可以使用切片
func TestRangeExpressionBySlice(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("sliceRangeExpression result:")
	fmt.Println("a = ", a)

	// 这里a[:]相当于把数组a转换成对应的切片了，切片内部是一个结构体由(*T,len, cap)组成，*T指向切片对应的数组指针，
	// 所以用切片变量就是操作了数组的指针
	for i, v := range a[:] {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

// 测试切片和数组append
func TestSliceLenRangeExpression(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	fmt.Println("TestSliceLenRangeExpression result:")
	fmt.Println("a = ", a)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}
		r = append(r, v)
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

// 总结：range表达式复制行为会造成一定性能损耗，简一使用数组指针或切片

// range string demo1
func TestRangeString1(t *testing.T) {
	// string内部结构为struct{*byte, len} 并且是不可变啊的
	var s = "中国人"
	// range string的时候，循环的单位是rune，而不是一个byte，返回的i为迭代字符码点的第一字节位置
	for i, v := range s {
		fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	}
}

// range string demo2
func TestRangeString2(t *testing.T) {
	var s1 = []byte{0xe4, 0xb8, 0xad, 0xe5, 0x9b, 0xbd, 0xe4, 0xba, 0xba}

	for _, v := range s1 {
		fmt.Printf("0x%x ", v)
	}

	fmt.Println()

	s1[3] = 0xd0
	s1[4] = 0xd6
	s1[5] = 0xb9

	for i, v := range string(s1) {
		fmt.Printf("%d %x\n", i, v)
	}
}

// range map
func TestRangeMap1(t *testing.T) {
	var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "tony")
		}
		counter++
		fmt.Println(k, v)
	}

	// 可能是3也可能是2
	fmt.Println("counter is ", counter)
}

func TestRangeMap2(t *testing.T) {
	var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			m["lucy"] = 24
		}
		counter++
		fmt.Println(k, v)
	}

	// 可能是4也可能是3
	fmt.Println("counter is ", counter)
}

// 总结1：for range遍历map无法保证顺序
// 总结2：遍历中对map修改，对后续迭代过程也是不确定的

// range channel
func TestRangeChannel1(t *testing.T) {
	var c = make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func TestRangeChannel2(t *testing.T) {
	var c chan int

	// nil channel会一直堵塞在这里
	for v := range c {
		fmt.Println(v)
	}
}

// 总结1：for range channel也会被复制，range的时候会一直阻塞在channel上，直到channel关闭
// 总结2：当遍历一个nil channel的时候会一直阻塞，直到程序陷入deadLock状态，并抛出panic
