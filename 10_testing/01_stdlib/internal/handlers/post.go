package handlers

import (
	"encoding/json"
	"errors"
	"example/internal/models"
	"example/internal/rest"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetPosts(db *gorm.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var posts []models.Post
		err := db.Order("id desc").Find(&posts).Error
		if err != nil {
			rest.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		rest.WriteJSON(w, http.StatusCreated, rest.Response{
			Ok:     true,
			Result: posts,
		})
	})
}

func GetPost(db *gorm.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			rest.WriteError(w, http.StatusBadRequest, err)
			return
		}

		var post models.Post
		err = db.Take(&post, id).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				rest.WriteError(w, http.StatusNotFound, errors.New("not found"))
				return
			}
			rest.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		rest.WriteJSON(w, http.StatusCreated, rest.Response{
			Ok:     true,
			Result: post,
		})
	})
}

func CreatePost(db *gorm.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var post models.Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			rest.WriteError(w, http.StatusBadRequest, err)
			return
		}

		// TODO: validation

		err = db.Create(&post).Error
		if err != nil {
			rest.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		rest.WriteJSON(w, http.StatusCreated, rest.Response{
			Ok:     true,
			Result: post.Id,
		})
	})
}

func UpdatePost(db *gorm.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			rest.WriteError(w, http.StatusBadRequest, err)
			return
		}

		var post models.Post
		err = json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			rest.WriteError(w, http.StatusBadRequest, err)
			return
		}
		post.Id = id

		// TODO: validation

		err = db.Select("*").Updates(post).Error
		if err != nil {
			rest.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		rest.WriteJSON(w, http.StatusCreated, rest.Response{
			Ok:     true,
			Result: post.Id,
		})
	})
}

func DeletePost(db *gorm.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			rest.WriteError(w, http.StatusBadRequest, err)
			return
		}

		err = db.Delete(&models.Post{}, id).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				rest.WriteError(w, http.StatusNotFound, errors.New("not found"))
				return
			}
			rest.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		rest.WriteJSON(w, http.StatusCreated, rest.Response{
			Ok: true,
		})
	})
}
