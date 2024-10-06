package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (f *HandlersFactory) GetAllDb() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func (f *HandlersFactory) CreateDb() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db_name := chi.URLParam(r, "db_name")

	})
}

func (f *HandlersFactory) DeleteDb() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db_name := chi.URLParam(r, "db_name")

	})
}

func (f *HandlersFactory) GetDb() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db_name := chi.URLParam(r, "db_name")

	})
}
