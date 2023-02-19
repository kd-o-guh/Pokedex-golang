package main

import (
	"fmt"
	"net/http"
	"pokedex-golang/manipulador"
)

func main() {

	fsc := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css", fsc))

	fsf := http.FileServer(http.Dir("favicons"))
	http.Handle("/favicons/", http.StripPrefix("/favicons", fsf))

	fsi := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images", fsi))

	http.HandleFunc("/", manipulador.Render)

	fmt.Println("Server is up and listening on port 8080.")
	http.ListenAndServe(":8080", nil)

}
