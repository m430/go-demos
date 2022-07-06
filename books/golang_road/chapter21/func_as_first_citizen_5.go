package main

import "fmt"

// 函数式编程
type IntSliceFunctor interface {
	Fmap(fn func(int) int) IntSliceFunctor
}

type intSliceFunctorImpl struct {
	ints []int
}

func (isf intSliceFunctorImpl) Fmap(fn func(int) int) IntSliceFunctor {
	newInts := make([]int, len(isf.ints))
	for i, elt := range isf.ints {
		retInt := fn(elt)
		newInts[i] = retInt
	}
	return intSliceFunctorImpl{ints: newInts}
}

func NewIntSliceFunctor(slice []int) IntSliceFunctor {
	return intSliceFunctorImpl{ints: slice}
}

func main() {
	intSlice := []int{1, 2, 3, 4}
	fmt.Printf("init a functor from int size: %#v\n", intSlice)
	f := NewIntSliceFunctor(intSlice)
	fmt.Printf("original functor: %+v\n", f)

	mapperFunc1 := func(i int) int {
		return i + 10
	}

	mapped1 := f.Fmap(mapperFunc1)
	fmt.Printf("mapped functor1: %+v\n", mapped1)

	mapperFunc2 := func(i int) int {
		return i * 3
	}

	mapped2 := mapped1.Fmap(mapperFunc2)
	fmt.Printf("mapped functor2: %+v\n", mapped2)
	fmt.Printf("original functor: %+v\n", f) // 原函子没有改变
	fmt.Printf("composite functor: %+v\n", f.Fmap(mapperFunc1).Fmap(mapperFunc2))
}

// 总结：上面例子实现了函数柯里化，使用了闭包的特性。闭包是在函数内部定义的匿名函数，并允许该匿名函数访问定义它的外部函数作用域
