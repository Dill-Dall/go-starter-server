package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dill-dall/go-starter-server/petapi"
	"github.com/dill-dall/go-starter-server/service"
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
	pets, err := service.ListPets(uint8(*params.Limit))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error retrieving pets: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pets)
}

func (t PetServer) ShowPetById(w http.ResponseWriter, r *http.Request, petId string) {
	pet, err := service.ShowPetById(petId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error retrieving pet: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pet)
}

func (t PetServer) SwaggerDoc(w http.ResponseWriter, r *http.Request) {
	swagger, err := petapi.GetSwagger()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error retrieving Swagger doc: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(swagger)
}

func (t PetServer) CreatePets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)

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

	configureTLS(h)
}

func configureTLS(h http.Handler) {
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
