package api

import (
	"errors"
	"net/http"

	"github.com/Mohamadreza-shad/notepad/repository"
	"github.com/gin-gonic/gin"
)

type NotepadHandler struct {
	NotepadRepository repository.Repository
}

func (h *NotepadHandler) CreateNotepad(c *gin.Context) {
	var params repository.CreateNotepadParams
	err := c.BindJSON(&params)
	if err != nil {
		MakeErrorResponseWithCode(c.Writer,http.StatusBadRequest,"invalid request")
		return
	}
	err = h.NotepadRepository.CreateNotepad(c.Request.Context(), params)
	if errors.Is(err, repository.ErrNotepadAlreadyExists) {
		MakeErrorResponseWithCode(c.Writer,http.StatusConflict,repository.ErrNotepadAlreadyExists.Error())
		return
	}
	if err != nil {
		MakeErrorResponseWithCode(c.Writer,http.StatusInternalServerError,repository.ErrSomethingWentWrong.Error())
		return
	}
	MakeSuccessResponse(c.Writer, nil, "notepad created successfully")
}

func NewNotepadHandler(r repository.Repository) *NotepadHandler {
	return &NotepadHandler{NotepadRepository: r}
}
