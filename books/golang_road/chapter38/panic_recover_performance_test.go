package chapter38

import (
	"errors"
	"testing"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func FooWithoutDefer() error {
	return errors.New("foo demo error")
}

func FooWithDefer() (err error) {
	defer func() {
		err = errors.New("foo demo error")
	}()
	return
}

func FooWithPanicAndRecover() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("foowithpanic demo error")
		}
	}()

	check(FooWithoutDefer())
	return nil
}

func FooWithoutPanicAndRecover() error {
	return FooWithDefer()
}

func BenchmarkFooWithoutPanicAndRecover(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FooWithoutPanicAndRecover()
	}
}

func BenchmarkFooWithPanicAndRecover(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FooWithPanicAndRecover()
	}
}
