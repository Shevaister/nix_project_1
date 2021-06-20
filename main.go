package main

import(
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Parser() {
	data := make([]map[string]interface{}, 0)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=7") 
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
	
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/parse")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	for _, value := range data {
		go flux1stGrade(db, value)
	}
	
	var input string
	fmt.Scanln(&input)
}

func flux1stGrade(db *sql.DB, data map[string]interface{}) {
	_, err := db.Exec("INSERT INTO posts VALUES(?, ?, ?, ?)", data["userId"], data["id"], data["title"], data["body"])
	if err != nil {
		fmt.Println(err)
	}
	dataComments := make([]map[string]interface{}, 0)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=" + fmt.Sprintf("%v", data["id"]))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &dataComments)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range dataComments {
		go flux2ndGrade(db, value)
	}

}

func flux2ndGrade(db *sql.DB, data map[string]interface{}) {
	_, err := db.Exec("INSERT INTO comments VALUES(?, ?, ?, ?, ?)", data["postId"], data["id"], data["name"], data["email"], data["body"])
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	Parser()
}