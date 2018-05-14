package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

const (
	DISTANCE = "200km"
)

type Location struct {
	Lat float64 `json: "lat"`
	Lon float64 `json: "lon"`
}

type Post struct {
	User string `json: "user"`
	Message string `json: "message"`
	Location Location `json: "location"`
}

func main() {
	fmt.Println("Hello, world")
	fmt.Print("started server")
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("received one post request")
	decoder := json.NewDecoder(r.Body)
	var p Post
	// &：取地址符
	if err := decoder.Decode(&p); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "POST IS RECEIVED %s\n", p.Message)
}

func handleSearch(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("received on search request")
	// get lat and lon
	lt := r.URL.Query().Get("lat") // e.g.  "lt = "37.5""
	ln := r.URL.Query().Get("lon")

	// system.out.println()
	fmt.Println(lt + " and " + ln)

	// range is option
	ran := DISTANCE

	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}
	fmt.Println("range is " + ran)

	lat,_ := strconv.ParseFloat(lt, 64) // string to float
	lon,_ := strconv.ParseFloat(ln, 64)

	// Return a fake post
	// &：取地址符
	p := &Post {
		User:"1111",
		Message:"一生必去的100个地方",
		Location: Location{
			Lat:lat,
			Lon:lon,
		},
	}

	// TO JSON object
	js, err := json.Marshal(p)
	if err != nil {
		panic(err)
		return
	}

	w.Header().Set("Conten-Type", "application/json")
	w.Write(js)
	//fmt.Fprint(w, "Search received: ", lat, lon)
}
