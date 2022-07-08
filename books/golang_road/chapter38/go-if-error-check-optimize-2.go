package chapter38

import (
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// 虽然使用check/handle风格优化了copyFile代码，可读性提高了。但是panic/recover处理性能降低了约90%
func CopyFile(src, dst string) (err error) {
	var r, w *os.File

	defer func() {
		if r != nil {
			r.Close()
		}
		if w != nil {
			w.Close()
		}
		if e := recover(); e != nil {
			if w != nil {
				os.Remove(dst)
			}
			err = fmt.Errorf("copy %s %s: %v", src, dst, err)
		}
	}()

	r, err = os.Open(src)
	check(err)

	w, err = os.Create(dst)
	check(err)

	_, err = io.Copy(w, r)
	check(err)

	return nil
}
