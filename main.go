package main

import(
	"fmt"
	"net/http"
	"io"
	"os"
	"strconv"
)

func flux(a string) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + a) 
	if err != nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	fmt.Println()
}

func main() {
	for i := 1; i <= 100; i++ {
		go flux(strconv.Itoa(i))
	}
	var input string
	fmt.Scanln(&input)
}