package controllers

import (
	"fmt"
	"mini-auction/middlewares"
	"mini-auction/models"
	"mini-auction/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auction(c *gin.Context) {
	c.File("./index.html")
}

func GetAuction() string {
	result, err := models.Client.EvalSha(models.Ctx, utils.QueryHash, []string{"product:1"}).Result()
	if err != nil {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	data, ok := result.([]interface{})
	if !ok {
		fmt.Println("Type assertion failed!")
		return ""
	}

	return data[7].(string)
}

func PostAuction(bid int, bidder string) (bool, int64) {

	bidStr := strconv.Itoa(bid)
	result, err := models.Client.EvalSha(models.Ctx, utils.UpdateHash, nil, bidStr, bidder).Result()
	if err != nil {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	if int64(bid) < result.(int64) {
		return false, int64(bid)
	}
	return true, result.(int64)
}
