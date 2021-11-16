package controllers

import (
	"ghitub/julhan07/redis-go/models"
	"ghitub/julhan07/redis-go/utils"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type userController struct {
	rdb *redis.Client
}

type UserControllerMapping interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

func NewUserController(rdb *redis.Client) UserControllerMapping {
	return &userController{
		rdb: rdb,
	}
}

func (c *userController) GetAll(w http.ResponseWriter, r *http.Request) {

	v, err := models.GetAllUser(r.Context(), c.rdb)

	if err != nil {
		utils.CreateResponseError(w, 400, err)
		return
	}

	utils.CreateResponseSuccess(w, 200, v)
	return
}
