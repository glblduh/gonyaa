# Go-Nyaa
*A unofficial api for nyaa.si written in Go*

## Installation
```
go get -u github.com/glblduh/gonyaa
```

## Example
```
package main

import (
	"log"
	"github.com/glblduh/gonyaa"
)

func main() {
	cli, err := gonyaa.NewClient("https://nyaa.si")
	if err != nil {
		log.Fatalln(err)
	}
	res := cli.Search("ANIME_NAME", "&c=1_2&o=desc")
	log.Printf("Name: %s\n", res[0].Name)
	log.Printf("Category: %s\n", res[0].Category)
	log.Printf("Date: %s\n", res[0].Date)
	log.Printf("Seeders: %d\n", res[0].Seeders)
	log.Printf("Leechers: %d\n", res[0].Leechers)
	log.Printf("Downloads: %d\n", res[0].Downloads)
	log.Printf("Magnet: %s\n", res[0].Magnet)
}
```
