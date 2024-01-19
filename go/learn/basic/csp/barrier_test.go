/*
* Barrier
 */
package csp_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

type barrierResp struct {
	Err    error
	Resp   string
	Status int
}

// 建造請求
// 宣告成 只寫channel
func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}

	client := http.Client{
		Timeout: time.Duration(2 * time.Microsecond), //故意觸發timeout
		// Timeout: time.Duration(2 * time.Second),
	}

	resp, err := client.Get(url)
	if resp != nil {
		res.Status = resp.StatusCode
	}
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}

// 合併結果
func barrier(endpoints ...string) {
	requestNumber := len(endpoints)

	in := make(chan barrierResp, requestNumber)
	response := make([]barrierResp, requestNumber)

	defer close(in)

	for _, endpoints := range endpoints {
		go makeRequest(in, endpoints)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err, resp.Status)
			hasError = true
		}
		response[i] = resp
	}
	if !hasError {
		for _, resp := range response {
			fmt.Println(resp.Status)
		}
	}
}

func TestBarrierURL(t *testing.T) {
	barrier([]string{"https://www.google.com/", "https://hub.docker.com/", "https://hackmd.io/"}...)
}

/*
Barrier 模式也可以使用errgroup 擴展庫來實現，這樣更加簡單明了。
這個包有點類似於sync.WaitGroup，但是區別是當其中一個任務發生錯誤時，可以返回該錯誤。
而這也滿足我們Barrier 模式的需求。
*/
func barrierErrGroup(endpoints ...string) {
	var g errgroup.Group
	var mu sync.Mutex

	response := make([]barrierResp, len(endpoints))

	for i, endpoint := range endpoints {
		i, endpoint := i, endpoint // create locals for closure below
		g.Go(func() error {
			res := barrierResp{}
			resp, err := http.Get(endpoint)
			if err != nil {
				return err
			}

			byt, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return err
			}

			res.Resp = string(byt)
			mu.Lock()
			response[i] = res
			mu.Unlock()
			return err
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("ERR:", err)
	}
	for _, resp := range response {
		fmt.Println("resp.Status:", resp.Status)
	}
}

func TestBarrierErrGroup(t *testing.T) {
	barrierErrGroup([]string{"https://www.google88.com/", "https://hub.docker.com/", "https://hackmd.io/"}...)
}
