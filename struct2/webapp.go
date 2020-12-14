package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1> Hey there</h1> We are Here in WebAPP
	<form>
	<br>
	Enter the Username: 
	<input type="text>
	<br>
	Enter the Password:
	<input type="password">
	<input type="submit">
	</form>
	
	`)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
