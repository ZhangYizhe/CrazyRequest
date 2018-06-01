package main

import (
	"net/http"
	"sync"
	"fmt"
	"io/ioutil"
	"bytes"
)


func main()  {

	wg := sync.WaitGroup{}
	for a := 0; a < 1; a++ {
		wg.Add(1)
		go requset(&wg,a)
	}
	wg.Wait()


}

//错误检查
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func requset(w *sync.WaitGroup,a int)  {
	
	resp, err := http.Get("https://testyizhe.tutuapp.com/ios/member")
	if err != nil {
		// handle error
		checkErr(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		checkErr(err)
	}

	test:=bytes.Count(body,nil)-1

	fmt.Println("============= "+fmt.Sprintf("%d", a)  +" =============="+fmt.Sprintf("%d", test) )

	w.Done()
}


