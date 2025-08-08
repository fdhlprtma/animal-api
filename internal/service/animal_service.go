package service

import (
	"database/sql"
	"animal-api/internal/model"
	"animal-api/internal/repository"
)

func GetAllAnimals(db *sql.DB) ([]model.Animal, error) {
	return repository.GetAll(db)
}

func GetAnimalByID(db *sql.DB, id string) (model.Animal, error) {
	return repository.GetByID(db, id)
}

func SearchAnimals(db *sql.DB, q string) ([]model.Animal, error) {
	return repository.Search(db, q)
}

func CreateAnimal(db *sql.DB, a model.Animal) error {
	return repository.Create(db, a)
}