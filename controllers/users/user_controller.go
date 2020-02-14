package users

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/teukumulya-ichsan/couchDB-go/interfaces"
	"github.com/teukumulya-ichsan/couchDB-go/repositories"
	"net/http"
)

type UserController struct {
	repo interfaces.UsersRepos
}

func NewUserController() *UserController {
	repos := repositories.NewCouchRepo()
	return &UserController{
		repo: repos,
	}
}

func (c *UserController) GetUserById(res http.ResponseWriter, r *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	docId := chi.URLParam(r, "id")
	data , err := c.repo.FindById(docId)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(res).Encode(err.Error())
	}
	//fmt.Println(data)
	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(&data)
}