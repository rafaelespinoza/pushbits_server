package api

import (
	"log"
	"net/http"

	"github.com/eikendev/pushbits/authentication"
	"github.com/eikendev/pushbits/model"

	"github.com/gin-gonic/gin"
)

// The ApplicationDatabase interface for encapsulating database access.
type ApplicationDatabase interface {
	CreateApplication(application *model.Application) error
	DeleteApplication(application *model.Application) error
	GetApplicationByID(ID uint) (*model.Application, error)
	GetApplicationByToken(token string) (*model.Application, error)
}

// The ApplicationDispatcher interface for relaying notifications.
type ApplicationDispatcher interface {
	RegisterApplication(name, user string) (string, error)
	DeregisterApplication(matrixID string) error
}

// ApplicationHandler holds information for processing requests about applications.
type ApplicationHandler struct {
	DB         ApplicationDatabase
	Dispatcher ApplicationDispatcher
}

func (h *ApplicationHandler) applicationExists(token string) bool {
	application, _ := h.DB.GetApplicationByToken(token)
	return application != nil
}

// CreateApplication creates a user.
func (h *ApplicationHandler) CreateApplication(ctx *gin.Context) {
	var createApplication model.CreateApplication

	if success := successOrAbort(ctx, http.StatusBadRequest, ctx.Bind(&createApplication)); !success {
		return
	}

	user := authentication.GetUser(ctx)

	application := model.Application{}
	application.Token = authentication.GenerateNotExistingToken(authentication.GenerateApplicationToken, h.applicationExists)
	application.UserID = user.ID

	log.Printf("User %s will receive notifications for application %s.\n", user.Name, application.Name)

	matrixid, err := h.Dispatcher.RegisterApplication(application.Name, user.MatrixID)

	if success := successOrAbort(ctx, http.StatusInternalServerError, err); !success {
		return
	}

	application.MatrixID = matrixid

	if success := successOrAbort(ctx, http.StatusInternalServerError, h.DB.CreateApplication(&application)); !success {
		return
	}

	ctx.JSON(http.StatusOK, &application)
}

// DeleteApplication deletes a user with a certain ID.
func (h *ApplicationHandler) DeleteApplication(ctx *gin.Context) {
	var deleteApplication model.DeleteApplication

	if success := successOrAbort(ctx, http.StatusBadRequest, ctx.BindUri(&deleteApplication)); !success {
		return
	}

	application, err := h.DB.GetApplicationByID(deleteApplication.ID)

	log.Printf("Deleting application %s.\n", application.Name)

	if success := successOrAbort(ctx, http.StatusBadRequest, err); !success {
		return
	}

	if success := successOrAbort(ctx, http.StatusInternalServerError, h.Dispatcher.DeregisterApplication(application.MatrixID)); !success {
		return
	}

	if success := successOrAbort(ctx, http.StatusInternalServerError, h.DB.DeleteApplication(application)); !success {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
