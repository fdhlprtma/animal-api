package repository

import (
	"database/sql"
	"encoding/json"
	"animal-api/internal/model"
)

func GetAll(db *sql.DB) ([]model.Animal, error) {
	rows, err := db.Query("SELECT * FROM animals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []model.Animal
	for rows.Next() {
		var a model.Animal
		var examples, classification string
		err := rows.Scan(&a.ID, &a.Name, &a.ImageURL, &classification, &a.Characteristics, &examples, &a.Habitat, &a.EcologicalRole)
		if err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(classification), &a.Classification)
		json.Unmarshal([]byte(examples), &a.Examples)
		animals = append(animals, a)
	}
	return animals, nil
}

func GetByID(db *sql.DB, id string) (model.Animal, error) {
	var a model.Animal
	var examples, classification string
	err := db.QueryRow("SELECT * FROM animals WHERE id = ?", id).Scan(
		&a.ID, &a.Name, &a.ImageURL, &classification, &a.Characteristics, &examples, &a.Habitat, &a.EcologicalRole)
	if err != nil {
		return a, err
	}
	json.Unmarshal([]byte(classification), &a.Classification)
	json.Unmarshal([]byte(examples), &a.Examples)
	return a, nil
}

func Search(db *sql.DB, query string) ([]model.Animal, error) {
	rows, err := db.Query("SELECT * FROM animals WHERE name LIKE ?", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []model.Animal
	for rows.Next() {
		var a model.Animal
		var examples, classification string
		err := rows.Scan(&a.ID, &a.Name, &a.ImageURL, &classification, &a.Characteristics, &examples, &a.Habitat, &a.EcologicalRole)
		if err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(classification), &a.Classification)
		json.Unmarshal([]byte(examples), &a.Examples)
		animals = append(animals, a)
	}
	return animals, nil
}

func Create(db *sql.DB, a model.Animal) error {
	classification, _ := json.Marshal(a.Classification)
	examples, _ := json.Marshal(a.Examples)

	_, err := db.Exec(`INSERT INTO animals (name, image_url, classification, characteristics, examples, habitat, ecological_role)
		VALUES (?, ?, ?, ?, ?, ?, ?)`, a.Name, a.ImageURL, classification, a.Characteristics, examples, a.Habitat, a.EcologicalRole)
	return err
}
