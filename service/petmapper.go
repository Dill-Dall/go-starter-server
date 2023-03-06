package service

import (
	"github.com/dill-dall/go-starter-server/petapi"
	"github.com/dill-dall/go-starter-server/repository"
)

func mapPetToEntity(pet petapi.Pet) repository.PetEntity {
	return repository.PetEntity{
		Name: pet.Name,
		Tag:  pet.Tag,
	}
}

func mapEntityToPet(entity repository.PetEntity) petapi.Pet {
	return petapi.Pet{
		Id:   entity.ID,
		Name: entity.Name,
		Tag:  entity.Tag,
	}
}