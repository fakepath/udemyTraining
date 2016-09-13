package main

//error checking enabled/disabled
import (
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
)

//production env
//func main() {
//	res, err := http.Get("http://sarabuna1.ro")
//	if err != nil {
//		log.Fatal(err)
//	}
//	page, err := ioutil.ReadAll(res.Body)
//	res.Body.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%s", page)
//}

func main() {
	res, _ := http.Get("http://sarabuna.ro")

	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", page)
}
