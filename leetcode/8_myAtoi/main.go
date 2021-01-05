package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//8. 字符串转换整数 (atoi)
//思路

func main() {
	names := []string{"xixixi", "hahaha", "Mark"}
	for _, name := range names {
		go func() {
			fmt.Printf("%s ", name)
		}()
		time.Sleep(time.Millisecond)
	}
}

func anotherStore(a atomic.Value)  {
	a.Store([]int{1,2,3})
	b := a.Load().([]int)
	b[0] = 3
}

type AI interface {
	Add()
	Des()
}

//type BI interface {
//	Des()
//	AI
//}

type A struct {
	A int
}

func (a *A) Add() {
	a.A = a.A + 1
}

func (a A) Des()  {

}

//func AddValue(delta int32) {
//	for  {
//		v := atomic.LoadInt32(&value)
//		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
//			break
//		}
//	}
//}