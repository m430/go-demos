package main

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

type result struct {
	value string
}

func first(servers ...*httptest.Server) (result, error) {
	c := make(chan result, len(servers))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	queryFunc := func(i int, server *httptest.Server) {
		url := server.URL
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Printf("query goroutine-%d: http NewRequests error: %s\n", i, err)
			return
		}
		req = req.WithContext(ctx)
		log.Printf("query goroutine-%d: send requests...\n", i)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("query goroutine-%d: get return error: %s\n", i, err)
			return
		}
		log.Printf("query goroutine-%d: get response\n", i)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		c <- result{
			value: string(body),
		}
		return
	}

	for i, server := range servers {
		go queryFunc(i, server)
	}
	select {
	case r := <-c:
		return r, nil
	case <-time.After(500 * time.Millisecond):
		return result{}, errors.New("timeout")
	}
}

func fakeWeatherServer(name string, interval int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s receive a http request\n", name)
		time.Sleep(time.Duration(interval) * time.Millisecond)
		w.Write([]byte(name + ":ok"))
	}))
}

func main() {
	result, err := first(fakeWeatherServer("open-weather-1", 200), fakeWeatherServer("open-weather-2", 1000), fakeWeatherServer("open-weather-3", 600))
	if err != nil {
		log.Println("invoke first error: ", err)
		return
	}
	log.Println(result)
	time.Sleep(10 * time.Second)
}

// 利用context.WithCancel创建了一个可以取消的context.Context，当其中一个goroutine返回结果后，其余的goroutine的http请求会被取消掉
