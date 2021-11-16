package models

import (
	"context"
	"encoding/json"
	"ghitub/julhan07/redis-go/entity"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func GetAllUser(ctx context.Context, rdb *redis.Client) ([]*entity.User, error) {
	var v []*entity.User

	for i := 0; i < 10; i++ {

		fullName := strconv.Itoa(i)
		address := strconv.Itoa(i)

		v = append(v, &entity.User{
			FullName: "Ulil" + fullName,
			Address:  "Jln Abc",
			Email:    "ulil" + address + "@email.com",
		})
	}

	ndata, _ := json.Marshal(v)

	err := rdb.HSet(ctx, "users", ndata).Err()

	if err != nil {
		return nil, err
	}

	return v, nil

}
