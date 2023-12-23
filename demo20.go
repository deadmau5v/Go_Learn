package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"_title_"`
	Year  string `json:"_year_"`
	Price string `json:"_price_"`
	By    string `json:"_by_"`
}

func main20() {
	var movie = Movie{"肖申克的救赎", "2001", "40", "US"}
	jsonData, err := json.Marshal(movie)
	if err != nil {
		println("Json Marshal Error!")
	}
	fmt.Printf("%s", jsonData)

	var movie2 Movie
	err2 := json.Unmarshal(jsonData, &movie2)
	if err2 != nil {
		println("Json Unmarshal Error!")
	}
	fmt.Println(movie2)
}
