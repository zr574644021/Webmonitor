package main

import (
	"fmt"
	"time"
	"net/http"
)

func  main()  {
	//a,err := http.Get("http://www.ccsdf.wer")
	//if err != nil {
	//	fmt.Printf("%s",err)
	//	return
	//}
	//fmt.Println(a.Status)
	http.Post()
	for i:=0;i<10;i++ {
		go func (j int) {
			fmt.Println("****",j)
		}(i)
	}
	time.Sleep(3*time.Second)


	//a := func (i int) int {
	//	return i * 2
	//}
	//
	//fmt.Println(a(1))
	//fmt.Println(func (i int) int {
	//	return i*2
	//}(1))
}

////func main() {
////	get := make(chan float64)
////	go func() {
////		time1 := time.Now()
////		_, err := http.Get("http://www.ccsx.cn")
////		if err != nil {
////			fmt.Println("error")
////			return
////		}
////		time2 := time.Now()
////		get <- time2.Sub(time1).Seconds()
////	}()
////	select {
////	case time_ := <-get:
////		fmt.Println(time_)
////	case <-time.After(3 * time.Second):
////		fmt.Println("3 second")
////	}
////
////}
