package main

import (
	"fmt"
	"net/http"
	"os"
)

func World(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello,world!")
}

// func F1() {
// 	fs := http.FileServer(http.Dir("static/page1.html"))
// 	http.Handle("/static/page1.html", fs)
// }

func Future(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello,future!")
}

// func F2() {
// 	fs := http.FileServer(http.Dir("static/page2.html"))
// 	http.Handle("/static/page2.html", fs)
// }

func DynamicRequest1(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("static/page1.html")
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s", content)
}

func DynamicRequest2(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("static/page2.html")
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s", content)
}

func main() {

	http.HandleFunc("/index/p1", World)
	http.HandleFunc("/index/p2", Future)
	// F1()
	// F2()
	http.HandleFunc("/index/p5", DynamicRequest1)
	http.HandleFunc("/index/p6", DynamicRequest2)
	fmt.Println("服务器开始监听...")
	http.ListenAndServe(":8080", nil)

}
