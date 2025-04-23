package handler

import (
	"todo/pkg/service"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	service  *service.Service
	validate *validator.Validate
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.signUp).Methods("POST")
		auth.HandleFunc("/sign-in", h.signIn).Methods("POST")
	}

	api := router.PathPrefix("/api").Subrouter()
	{
		api.Use(h.userIdentity)
		lists := api.PathPrefix("/lists").Subrouter()
		{
			lists.HandleFunc("/", h.createList).Methods("POST")
			lists.HandleFunc("/", h.getAllLists).Methods("GET")
			lists.HandleFunc("/{id}", h.getListById).Methods("GET")
			lists.HandleFunc("/{id}", h.updateList).Methods("PUT")
			lists.HandleFunc("/{id}", h.deleteList).Methods("DELETE")

			items := lists.PathPrefix("/{id}/items").Subrouter()
			{
				items.HandleFunc("/", h.createItem).Methods("POST")
				items.HandleFunc("/", h.getAllItems).Methods("GET")
			}
		}

		items := api.PathPrefix("/items/{id}").Subrouter()
		{
			items.HandleFunc("/", h.getItemById).Methods("GET")
			items.HandleFunc("/", h.updateItem).Methods("PUT")
			items.HandleFunc("/", h.deleteItem).Methods("DELETE")
		}
	}
	return router
}
