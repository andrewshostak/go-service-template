package handler

import (
	"github.com/andrewshostak/awesome-service/errs"
	"github.com/andrewshostak/awesome-service/model"
	"github.com/andrewshostak/awesome-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QuestionHandler interface {
	Create(context *gin.Context)
	Delete(context *gin.Context)
	List(context *gin.Context)
	One(context *gin.Context)
	Update(context *gin.Context)
}

type questionHandler struct {
	qs service.QuestionService
}

func NewQuestionHandler(qs service.QuestionService) QuestionHandler {
	if qs == nil {
		panic("question service is nil")
	}
	return &questionHandler{qs: qs}
}

func (h *questionHandler) Create(context *gin.Context) {
	var question model.QuestionCreate
	if err := context.ShouldBindJSON(&question); err != nil {
		context.Error(errs.New(err, errs.UserError))
		return
	}

	created, err := h.qs.Create(context.Request.Context(), question)
	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"question": created,
	})
}

func (h *questionHandler) Delete(context *gin.Context) {
	var uriParams UriParams
	if err := context.ShouldBindUri(&uriParams); err != nil {
		context.Error(errs.New(err, errs.UserError))
		return
	}

	if err := h.qs.Delete(context.Request.Context(), uriParams.Id); err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func (h *questionHandler) List(context *gin.Context) {
	questions, err := h.qs.List(context.Request.Context())
	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"questions": questions,
	})
}

func (h *questionHandler) One(context *gin.Context) {
	var uriParams UriParams
	if err := context.ShouldBindUri(&uriParams); err != nil {
		context.Error(errs.New(err, errs.UserError))
		return
	}

	question, err := h.qs.One(context.Request.Context(), uriParams.Id)
	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"question": question,
	})
}

func (h *questionHandler) Update(context *gin.Context) {
	var uriParams UriParams
	if err := context.ShouldBindUri(&uriParams); err != nil {
		context.Error(errs.New(err, errs.UserError))
		return
	}

	var question model.QuestionUpdate
	if err := context.ShouldBindJSON(&question); err != nil {
		context.Error(errs.New(err, errs.UserError))
		return
	}

	updated, err := h.qs.Update(context.Request.Context(), uriParams.Id, question)
	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"question": updated,
	})
}

type UriParams struct {
	Id uint `uri:"id" binding:"required"`
}
