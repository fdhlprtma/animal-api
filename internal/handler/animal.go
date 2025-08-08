package handler

import (
	"animal-api/internal/model"
	"animal-api/internal/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func GetAnimals(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		query := r.URL.Query().Get("q")
		var (
			data interface{}
			err  error
		)

		if query != "" {
			data, err = service.SearchAnimals(db, query)
		} else {
			data, err = service.GetAllAnimals(db)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		output, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(output)
	}
}

func GetAnimalByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		animal, err := service.GetAnimalByID(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		output, err := json.MarshalIndent(animal, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(output)
	}
}

func SearchAnimals(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		animals, err := service.SearchAnimals(db, q)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		output, err := json.MarshalIndent(animals, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(output)
	}
}

func CreateAnimal(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		// === Mode 1: multipart/form-data ===
		if strings.HasPrefix(contentType, "multipart/form-data") {
			err := r.ParseMultipartForm(10 << 20) // max 10 MB
			if err != nil {
				http.Error(w, "Gagal parsing form: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Ambil data form
			animal := model.Animal{
				Name:            r.FormValue("name"),
				Habitat:         r.FormValue("habitat"),
				Characteristics: r.FormValue("characteristics"),
				EcologicalRole:  r.FormValue("ecological_role"),
				Classification: model.Classification{
					Kingdom: r.FormValue("kingdom"),
					Phylum:  r.FormValue("phylum"),
					Class:   r.FormValue("class"),
					Order:   r.FormValue("order"),
					Family:  r.FormValue("family"),
				},
				Examples: r.MultipartForm.Value["examples"],
			}

			// Handle upload file
			file, handler, err := r.FormFile("image")
			if err != nil {
				http.Error(w, "Gagal ambil file gambar: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()

			// Simpan file ke folder
			uploadFolder := "uploads"
			os.MkdirAll(uploadFolder, os.ModePerm)
			filePath := fmt.Sprintf("%s/%s", uploadFolder, handler.Filename)

			dst, err := os.Create(filePath)
			if err != nil {
				http.Error(w, "Gagal simpan file: "+err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()
			io.Copy(dst, file)

			animal.ImageURL = fmt.Sprintf("https://yourdomain.com/animal-api/%s", filePath)

			// Simpan ke DB
			if err := service.CreateAnimal(db, animal); err != nil {
				http.Error(w, "Gagal simpan ke DB: "+err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message":   "Hewan berhasil dibuat",
				"image_url": animal.ImageURL,
			})
			return
		}

		// === Mode 2 & 3: application/json ===
		if strings.HasPrefix(contentType, "application/json") {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Gagal membaca body: "+err.Error(), http.StatusBadRequest)
				return
			}

			trimmed := strings.TrimSpace(string(body))

			if strings.HasPrefix(trimmed, "[") {
				// JSON array
				var animals []model.Animal
				if err := json.Unmarshal(body, &animals); err != nil {
					http.Error(w, "Gagal parsing JSON array: "+err.Error(), http.StatusBadRequest)
					return
				}

				for _, animal := range animals {
					// Default image URL kalau kosong
					if animal.ImageURL == "" {
						animal.ImageURL = "https://yourdomain.com/default.jpg"
					}

					if err := service.CreateAnimal(db, animal); err != nil {
						http.Error(w, "Gagal simpan salah satu hewan: "+err.Error(), http.StatusInternalServerError)
						return
					}
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Semua hewan berhasil disimpan",
				})
				return

			} else if strings.HasPrefix(trimmed, "{") {
				// JSON object
				var animal model.Animal
				if err := json.Unmarshal(body, &animal); err != nil {
					http.Error(w, "Gagal parsing JSON object: "+err.Error(), http.StatusBadRequest)
					return
				}

				if animal.ImageURL == "" {
					animal.ImageURL = "https://yourdomain.com/default.jpg"
				}

				if err := service.CreateAnimal(db, animal); err != nil {
					http.Error(w, "Gagal simpan hewan: "+err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"message":   "Hewan berhasil disimpan",
					"image_url": animal.ImageURL,
				})
				return
			}

			http.Error(w, "Format JSON tidak dikenali", http.StatusBadRequest)
			return
		}

		// === Tidak cocok semua ===
		http.Error(w, "Content-Type tidak didukung", http.StatusUnsupportedMediaType)
	}
}
