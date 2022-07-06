package chapter21

import (
	"net/http"
	"time"
)

// 正常创建
func newPrinter() {

}

// 在函数内创建
func funcA() {
	funcB := func() {

	}

	funcB()
}

// 作为类型
type HandlerFunc func(response http.Response, r *http.Request)

// 存储到变量中
func A() {
	a := func() {}
	a()
}

// 作为参数传入函数
func AfterFunc(d time.Duration, f func()) {
}

// 作为返回值从函数返回
func C() func() {
	return func() {

	}
}

// 总结: Go语言中函数是一等公民，函数可以像普通整型那样创建和使用，还可以放入数组、切片或map中，可以像其他类型变量一样赋值
// 赋值给interface{}，甚至可以建立元素为函数的channel
