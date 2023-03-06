package service

import (
	"fmt"
	"log"

	"github.com/dill-dall/go-starter-server/petapi"
	"github.com/dill-dall/go-starter-server/repository"
)

func ListPets(limit uint8) ([]petapi.Pet, error) {
	entityPets, err := repository.NewPetRepository().GetPets(limit)
	if err != nil {
		log.Printf("Error in ListPets: %v", err)
		return nil, fmt.Errorf("error retrieving pets: %v", err)
	}

	entities := make([]petapi.Pet, len(entityPets))
	for i, pet := range entityPets {
		entities[i] = mapEntityToPet(pet)
	}

	return entities, nil
}

func ShowPetById(petId string) (petapi.Pet, error) {
	petEntityPtr, err := repository.NewPetRepository().GetPet(petId)

	if err != nil {
		log.Printf("Error in ShowPetById: %v", err)
		return petapi.Pet{}, fmt.Errorf("error retrieving pet: %v", err)
	}

	petEntity := *petEntityPtr
	apiPet := mapEntityToPet(petEntity)

	return apiPet, nil
}

func CreatePet(pet petapi.Pet) error {
	petEntity := mapPetToEntity(pet)
	err := repository.NewPetRepository().Add(petEntity)
	if err != nil {
		log.Printf("Error in CreatePet: %v", err)
		return fmt.Errorf("error creating pet: %v", err)
	}
	return nil
}
