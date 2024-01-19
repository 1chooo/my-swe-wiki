/*
* Future
 */
package csp_test

import (
	"fmt"
	"testing"
)

type Function func(string) (string, error)

type Future interface {
	SuccessCallback() error
	FailCallback() error
	Execute(Function) (bool, chan struct{})
}

type AccountCache struct {
	Name string
}

func (a *AccountCache) SuccessCallback() error {
	fmt.Println("It's success~")
	return nil
}

func (a *AccountCache) FailCallback() error {
	fmt.Println("It's fail~")
	return nil
}

func (a *AccountCache) Execute(f Function) (bool, chan struct{}) {
	done := make(chan struct{})
	go func(a *AccountCache) {
		_, err := f(a.Name)
		if err != nil {
			_ = a.FailCallback()
		} else {
			_ = a.SuccessCallback()
		}
		done <- struct{}{}
	}(a)
	return true, done
}

func NewAccountCache(name string) *AccountCache {
	return &AccountCache{
		name,
	}
}

func TestFuture1(t *testing.T) {
	var future Future = NewAccountCache("Tom")
	updateFunc := func(name string) (string, error) {
		fmt.Println("cache update:", name)
		return name, nil
	}
	_, done := future.Execute(updateFunc)
	defer func() {
		<-done
	}()
}

func TestFuture2(t *testing.T) {
	// var future Future = NewAccountCache("Tom")
	future := NewAccountCache("Tom")
	updateFunc := func(name string) (string, error) {
		fmt.Println("cache update:", name)
		return name, nil
	}
	_, done := future.Execute(updateFunc)
	defer func() {
		<-done
	}()
	// do something
}

/*
這裡有一個技巧：為什麼使用struct 類型作為channel 的通知？
很多開源代碼都是使用這種方式來作為信號通知機制，主要是因為空struct 在Go 中佔的內存是最少的。
*/
