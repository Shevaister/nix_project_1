package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type comments struct {
	gorm.Model
	ID     uint
	PostID uint
	Name   string
	Email  string
	Body   string
}

type posts struct {
	gorm.Model
	ID     uint
	UserID uint
	Title  string
	Body   string
}

func parser() {
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

	dsn := "mysql:mysql@tcp(127.0.0.1:3306)/parse?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&posts{})
	db.AutoMigrate(&comments{})

	for _, value := range data {
		go flux1stGrade(db, value)
	}

	var input string
	fmt.Scanln(&input)
}

func flux1stGrade(db *gorm.DB, data map[string]interface{}) {
	post := posts{ID: uint(data["id"].(float64)), UserID: uint(data["userId"].(float64)), Title: data["title"].(string), Body: data["body"].(string)}
	db.Create(&post)
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

func flux2ndGrade(db *gorm.DB, data map[string]interface{}) {
	comment := comments{ID: uint(data["id"].(float64)), PostID: uint(data["postId"].(float64)), Name: data["name"].(string), Email: data["email"].(string), Body: data["body"].(string)}
	db.Create(&comment)
}

func main() {
	parser()
}
