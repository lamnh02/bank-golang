package api

type Server struct {
	store *db.Store
	router
}