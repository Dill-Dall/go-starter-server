package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/d-vignesh/go-starter-server/petapi"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type ServerInterface interface {
	petapi.ServerInterface
}

type PetServer struct {
	// any additional fields here
}

func (t PetServer) ListPets(w http.ResponseWriter, r *http.Request, params petapi.ListPetsParams) {
	pets := getPets()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pets)
}

func getPets() []petapi.Pet {
	hello := "jklds"
	pets := []petapi.Pet{{Id: 1, Name: "Garfield", Tag: &hello}, {Id: 2, Name: "Odie", Tag: &hello}}
	return pets
}

func (t PetServer) CreatePets(w http.ResponseWriter, r *http.Request) {
	// our logic to store the pet into a persistent layer
}

func (t PetServer) ShowPetById(w http.ResponseWriter, r *http.Request, petId string) {
	// our logic to get a pet by ID from the persistent layer
}

func (t PetServer) SwaggerDoc(w http.ResponseWriter, r *http.Request) {
	swagger, _ := petapi.GetSwagger()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(swagger)

}

const port = "3000"

func main() {
	envFilePath := flag.String("env", ".env", "Path to .env file")
	flag.Parse()

	err := godotenv.Load(*envFilePath)
	if err != nil {
		log.Fatal(err)
	}

	s := PetServer{}
	h := petapi.Handler(s)

	printRoutes(h.(*chi.Mux))

	confugureTlS(h)
}

func confugureTlS(h http.Handler) {
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	if certFile != "" && keyFile != "" {
		fmt.Printf("%s://%s:%s\n", "https", getHostname(), port)
		err := http.ListenAndServeTLS(getAddress(), certFile, keyFile, h)
		log.Fatal(err)
	} else {
		fmt.Printf("%s://%s:%s\n", "http", getHostname(), port)
		err := http.ListenAndServe(getAddress(), h)
		log.Fatal(err)
	}
}

func printRoutes(r chi.Router) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s: %s\n", method, route)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		// handle error
		log.Fatal(err)
	}
}

func getHostname() string {
	if os.Getenv("") != "local" {
		return "localhost"
	} else {
		return "0.0.0.0"

	}
}

func getAddress() string {
	return fmt.Sprintf("%s:%s", getHostname(), port)
}
