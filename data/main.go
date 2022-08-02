package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Penampung struct {
	Name string
	Data []*Data
}

type Data struct {
	Title       string
	Link        string
	Description string
	Content     string
}

func main() {
	container := getData()

	//ngambil konten berdasarkan link
	for _,j := range container.Data{
		j.Content = getContent(j.Link)
		fmt.Println(*j)
	}

	hasil, err := json.Marshal(container)
	errorHandler(err)

	writeData(hasil)


}
func errorHandler(e error)  {
	if e!=nil {
		log.Fatal(e)
	}
}


func writeData(param []byte)  {
	err := ioutil.WriteFile("exported.txt", param, 0775)
	errorHandler(err)
	
}


func getContent(path string) string  {
	hasil, err := ioutil.ReadFile(path)
	errorHandler(err)
	return string(hasil)
}

func getData() Penampung {
	hasil, err := ioutil.ReadFile("./data/data.json")
	errorHandler(err)
	var p Penampung
	jErr:=json.Unmarshal(hasil, &p) // unmarshal : byte => golang
	errorHandler(jErr)
	return p
}