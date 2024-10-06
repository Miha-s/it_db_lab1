package handlers

import (
	"net/http"
)

func (f *HandlersFactory) GetTable() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// db_name := chi.URLParam(r, "db_name")
		// table_name := chi.URLParam(r, "table_name")

	})
}

func (f *HandlersFactory) CreateTable() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// db_name := chi.URLParam(r, "db_name")
		// table_name := chi.URLParam(r, "table_name")

	})
}

func (f *HandlersFactory) DeleteTable() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// db_name := chi.URLParam(r, "db_name")
		// table_name := chi.URLParam(r, "table_name")

	})
}

func (f *HandlersFactory) UpdateTable() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// db_name := chi.URLParam(r, "db_name")
		// table_name := chi.URLParam(r, "table_name")

	})
}

func (f *HandlersFactory) RemoveDuplicates() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// db_name := chi.URLParam(r, "db_name")
		// table_name := chi.URLParam(r, "table_name")

	})
}
