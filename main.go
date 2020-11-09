package main

import (
	"fmt"
	redis "gopkg.in/redis.v4"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

// import "fmt"
// func main() {
//     fmt.Println("Hello World")
// }

// import (
// 	"fmt"
// 	"html"
// 	"log"
// 	"net/http"
// )

// func main() {

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	})

// 	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hi")
// 	})

// 	log.Fatal(http.ListenAndServe(":8081", nil))

// }
