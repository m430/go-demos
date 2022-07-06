package chapter22

import (
	"errors"
	"os"
	"sync"
)

// 如果没有defer的话，开发人员需要特别关注资源的释放
func writeToFile(fname string, data []byte, mu *sync.Mutex) error {
	mu.Lock()
	f, err := os.OpenFile(fname, os.O_RDWR, 0666)
	if err != nil {
		mu.Unlock()
		return err
	}

	_, err = f.Seek(0, 2)
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}

	err = f.Close()
	if err != nil {
		mu.Unlock()
		return err
	}

	mu.Unlock()
	return nil
}

// 用defer函数后，开发人员不用小心翼翼的在每个错误处理分支检查是否遗漏了资源释放
func writeToFileByDefer(fname string, data []byte, mu *sync.Mutex) error {
	mu.Lock()
	defer mu.Unlock()
	f, err := os.OpenFile(fname, os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Seek(0, 2)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return f.Sync()
}

// 常见用法1：拦截panic
func makeSlice(n int) []byte {
	defer func() {
		// 尝试恢复panic
		if recover() != nil {
			panic(errors.New("出发一个新的panic"))
		}
	}()
	return make([]byte, n)
}

//
