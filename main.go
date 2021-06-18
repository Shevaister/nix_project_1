package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	//"os"
	"strconv"
)

func flux(a string) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + a) 
	if err != nil {
		fmt.Println("read1",err)
		return
	}
	defer resp.Body.Close()
	
	n, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read2", err)
		return
	}
	ioutil.WriteFile("results/" + a + ".txt", n, 0777) 
}

func main() {
	for i := 1; i <= 100; i++ {
		go flux(strconv.Itoa(i))
	}
	var input string
	fmt.Scanln(&input)
}