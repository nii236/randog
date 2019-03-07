package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
		if resp.StatusCode > 299 {
			log.Println("non 200 response:", resp.StatusCode)
		}
		if err != nil {
			log.Println("error fetching doggo:", err)
		}
		dogResp := &DogResp{}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("error reading doggo response:", err)
		}
		defer resp.Body.Close()
		err = json.Unmarshal(b, dogResp)
		if err != nil {
			log.Println("error unmarshalling doggo response:", err)
		}
		http.Redirect(w, r, dogResp.Message, http.StatusTemporaryRedirect)
	})
	log.Println("Hosting on :3333")
	log.Fatalln(http.ListenAndServe(":3333", r))
}

type DogResp struct {
	Status  string
	Message string
}
