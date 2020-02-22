package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/teukumulya-ichsan/couchDB-go/interfaces"
	"github.com/teukumulya-ichsan/couchDB-go/models"
	"github.com/teukumulya-ichsan/couchDB-go/repositories"
)

// UserController struct ...
type UserController struct {
	repo interfaces.UsersRepos
}

// NewUserController Constructor ...
func NewUserController() *UserController {
	repos := repositories.NewCouchRepo()
	return &UserController{
		repo: repos,
	}
}

// GetUserByID Method (mutation) ...
func (c *UserController) GetUserByID(res http.ResponseWriter, r *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	docID := chi.URLParam(r, "id")

	data, err := c.repo.FindByID(docID)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(res).Encode(err.Error())
	}

	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(&data)
}

func (c *UserController) CreateUser(res http.ResponseWriter, r *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDocoder.
	if err!= nil{
	}
	c.repo.SaveDocs()
}
