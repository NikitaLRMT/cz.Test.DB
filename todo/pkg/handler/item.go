package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo"

	"github.com/gorilla/mux"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid list id param")
		return
	}

	defer r.Body.Close()
	var input todo.TodoItem
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.service.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, items)
}

func (h *Handler) getItemById(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.service.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, item)
}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid id param")
		return
	}

	defer r.Body.Close()
	var input todo.UpdateItemInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid item id param")
		return
	}

	err = h.service.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, statusResponse{"ok"})
}
