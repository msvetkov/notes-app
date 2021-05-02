package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/msvetkov/notes-app/pkg/domain"
	"net/http"
	"strconv"
	"time"
)

type getNotesResponse struct {
	Data []domain.Note `json:"data"`
}

// @Summary Get all notes
// @Security ApiKeyAuth
// @Tags notes
// @Description get all notes
// @ID get-all-notes
// @Accept  json
// @Produce  json
// @Success 200 {object} getNotesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/notes [get]
func (h *Handler) getNotes(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.services.Note.GetAll(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getNotesResponse{
		Data: list,
	})
}

// @Summary Get note by id
// @Security ApiKeyAuth
// @Tags notes
// @Description get note by id
// @ID get-note-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Note
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/notes [get]
func (h *Handler) getNote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	note, err := h.services.Note.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, note)
}

// @Summary Create note
// @Security ApiKeyAuth
// @Tags notes
// @Description create note
// @ID create-note
// @Accept  json
// @Produce  json
// @Param input body domain.Note true "note info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/notes [post]
func (h *Handler) createNote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input domain.Note
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Note.Create(domain.Note{
		Title:       input.Title,
		Body:        input.Body,
		DateCreated: time.Now(),
		UserId:      userId,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Update note by id
// @Security ApiKeyAuth
// @Tags notes
// @Description update note by id
// @ID get-all-notes
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/notes [put]
func (h *Handler) updateNote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	var input domain.UpdateNoteInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Note.Update(userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete note by id
// @Security ApiKeyAuth
// @Tags notes
// @Description delete note by id
// @ID delete-note-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/notes [delete]
func (h *Handler) deleteNote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	err = h.services.Note.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
