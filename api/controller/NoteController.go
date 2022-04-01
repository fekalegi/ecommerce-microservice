package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ecommerce-microservice/api/middlewares"
	"ecommerce-microservice/common/dto"
	"ecommerce-microservice/common/models"
	"ecommerce-microservice/usecase"
	"strconv"
)

// NoteController : Controller for Note Api
type noteController struct {
	noteService usecase.NoteService
}

// NewNoteController : Instance for register  NoteService
func NewNoteController(noteService usecase.NoteService) *noteController {
	return &noteController{noteService: noteService}
}

// Route : Setting Note Route
func (n *noteController) Route(group *echo.Group) {
	group.POST("/notes/", n.CreateNote, middleware.KeyAuth(middlewares.AuthCheck))
	group.GET("/notes", n.FetchAllNote)
	group.GET("/notes/:id", n.FetchNote)
	group.PUT("/notes/:id", n.UpdateNote, middleware.KeyAuth(middlewares.AuthCheck))
	group.DELETE("/notes/:id", n.DeleteNote, middleware.KeyAuth(middlewares.AuthCheck))
}

// CreateNote godoc
// @Summary Create New Note
// @Description This endpoint for create new Note
// @Tags Note
// @Accept  json
// @Produce  json
// @Param services body dto.RequestNote true "Create new note"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /notes/ [post]
func (n *noteController) CreateNote(ctx echo.Context) error {
	userID := ctx.Get("user_id")
	conv, err := strconv.Atoi(fmt.Sprint(userID))
	if err != nil {
		return err
	}

	request := dto.RequestNote{}

	err = models.ValidateRequest(ctx, &request)

	if err != nil {
		return err
	}

	responses, err := n.noteService.CreateNote(request, conv)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchAllNote : Fetch All Note
// @Summary Fetch All Note
// @Description This endpoint for fetch all product category
// @Tags Note
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=[]domain.Note} "desc"
// @Router /notes [get]
func (n *noteController) FetchAllNote(ctx echo.Context) error {

	responses, err := n.noteService.FetchAllNote()
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchNote : Fetch Note by ID
// @Summary Fetch Note by ID
// @Description This endpoint for fetch product category by ID
// @Tags Note
// @Accept  json
// @Produce  json
// @Param id path string true "Note ID"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Note} "desc"
// @Router /notes/{id} [get]
func (n *noteController) FetchNote(ctx echo.Context) error {
	parameter := ctx.Param("id")
	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := n.noteService.FetchNote(converted)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// UpdateNote : Update Note by ID
// @Summary Update Note by ID
// @Description This endpoint for Update product category by ID
// @Tags Note
// @Accept  json
// @Produce  json
// @Param id path string true "Note ID"
// @Param services body dto.RequestNote true "Create new note"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Note} "desc"
// @Router /notes/{id} [put]
func (n *noteController) UpdateNote(ctx echo.Context) error {
	parameter := ctx.Param("id")
	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	userID := ctx.Get("user_id")
	conv, err := strconv.Atoi(fmt.Sprint(userID))
	if err != nil {
		return err
	}

	request := dto.RequestNote{}

	err = models.ValidateRequest(ctx, &request)

	if err != nil {
		return err
	}

	responses, err := n.noteService.UpdateNote(converted, conv, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// DeleteNote : Delete Note by ID
// @Summary Delete Note by ID
// @Description This endpoint for Delete product category by ID
// @Tags Note
// @Accept  json
// @Produce  json
// @Param id path string true "Note ID"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Note} "desc"
// @Router /notes/{id} [delete]
func (n *noteController) DeleteNote(ctx echo.Context) error {
	parameter := ctx.Param("id")
	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	userID := ctx.Get("user_id")
	conv, err := strconv.Atoi(fmt.Sprint(userID))
	if err != nil {
		return err
	}

	responses, err := n.noteService.DeleteNote(converted, conv)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
