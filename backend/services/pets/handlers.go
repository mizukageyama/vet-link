package pets

import (
	"backend/framework"
	"backend/generated/models"
	. "backend/globals"
	"backend/services/users"
	"github.com/stephenafamo/bob"
	"time"
)

func (s *Service) handleList(c *framework.Context, data framework.ResultChan, err framework.ErrorChan) {
	params := new(listRequest)
	if e := c.Api.Bind(params); e != nil {
		err <- e
	} else if pets, e := listPets(
		c, s.Store.Db,
		models.SelectWhere.Pets.OwnerID.EQ(params.OwnerID),
		models.PreloadPetBreed(models.ThenLoadBreedAnimal()),
	); e != nil {
		err <- e
	} else {
		data <- pets
	}
}

func (s *Service) handleRead(c *framework.Context, data framework.ResultChan, err framework.ErrorChan) {
	params := new(readRequest)
	if e := c.Api.Bind(params); e != nil {
		err <- e
	} else if pet, e := models.Pets.Query(
		c.GetContext(), s.Store.Db,
		models.SelectWhere.Pets.ID.EQ(params.ID),
		models.PreloadPetBreed(models.ThenLoadBreedAnimal()),
	).One(); e != nil {
		err <- e
	} else if birthdate, e := time.Parse(time.DateOnly, pet.Birthdate); e != nil {
		err <- e
	} else {
		data <- PetDTO{
			ID:        pet.ID,
			Name:      pet.Name,
			Age:       CalculateAge(birthdate),
			Gender:    Elvis(pet.Gender, "Male", "Female"),
			Birthdate: pet.Birthdate,
			Breed: &BreedDTO{
				ID:      pet.R.Breed.ID,
				Breed:   pet.R.Breed.Description,
				Species: &pet.R.Breed.R.Animal.Description,
			},
		}
	}
}

func (s *Service) handleUpdate(c *framework.Context, data framework.ResultChan, err framework.ErrorChan) {
	params := new(updateRequest)
	if e := c.Api.Bind(params); e != nil {
		err <- e
	} else if e := s.Store.Transact(c.GetContext(), func(tx *bob.Tx) error {
		return updatePet(c, tx, params)
	}); e != nil {
		err <- e
	} else {
		updated, e := models.Pets.Query(
			c.GetContext(), s.Store.Db,
			models.SelectWhere.Pets.ID.EQ(params.ID),
			models.PreloadPetBreed(),
			models.PreloadPetOwnerUser(),
		).One()
		if e != nil {
			err <- e
			return
		}
		birthdate, _ := time.Parse(time.DateOnly, updated.Birthdate)
		data <- PetDTO{
			ID:        updated.ID,
			Name:      updated.Name,
			Age:       CalculateAge(birthdate),
			Gender:    Elvis(updated.Gender, "Male", "Female"),
			Birthdate: updated.Birthdate,
			Breed: &BreedDTO{
				ID:    updated.R.Breed.ID,
				Breed: updated.R.Breed.Description,
			},
			Owner: &users.UserResponse{
				ID:         updated.R.OwnerUser.ID,
				Email:      updated.R.OwnerUser.Email,
				FamilyName: updated.R.OwnerUser.FamilyName,
				GivenName:  updated.R.OwnerUser.GivenName,
			},
		}
	}
}

func (s *Service) handleDelete(c *framework.Context, data framework.ResultChan, err framework.ErrorChan) {
	params := new(deleteRequest)
	if e := c.Api.Bind(params); e != nil {
		err <- e
	} else if e := s.Store.Transact(c.GetContext(), func(tx *bob.Tx) error {
		if e = deletePet(c, tx, params.ID); e != nil {
			return e
		}
		data <- "Successfully deleted pet"
		return nil
	}); e != nil {
		err <- e
	}
}

func (s *Service) handleCreate(c *framework.Context, data framework.ResultChan, err framework.ErrorChan) {
	params := new(CreateRequest)
	if e := c.Api.Bind(params); e != nil {
		err <- e
	} else if e = s.Store.Transact(c.GetContext(), func(tx *bob.Tx) error {
		pet, e := insertPet(c, tx, params)
		if e != nil {
			return e
		}
		data <- pet
		return nil
	}); e != nil {
		err <- e
	}
}

func (s *Service) handleReadAllBreeds(c *framework.Context, data framework.ResultChan, err framework.ErrorChan) {
	params := new(listBreedsRequest)
	if e := c.Api.Bind(params); e != nil {
		err <- e
	} else if e = s.Store.Transact(c.GetContext(), func(tx *bob.Tx) error {
		breeds, e := listBreeds(c, tx,
			models.SelectWhere.Breeds.AnimalID.EQ(params.AnimalID),
			models.PreloadBreedAnimal(),
		)
		if e != nil {
			return e
		}
		data <- breeds
		return nil
	}); e != nil {
		err <- e
	}
}

type BreedDTO struct {
	ID      uint32  `json:"id"`
	Breed   string  `json:"breed"`
	Species *string `json:"species,omitempty"`
}
