package utils

import (
	"fmt"
	"mini-auction/models"
	"os"
)

var QueryHash string
var UpdateHash string

func LoadScripts() {

	queryByte, err := os.ReadFile("query.lua")
	if err != nil {
		fmt.Println("Error reading Lua script file:", err)
		return
	}

	updateByte, err := os.ReadFile("update.lua")
	if err != nil {
		fmt.Println("Error reading Lua script file:", err)
		return
	}

	queryString := string(queryByte)
	updateString := string(updateByte)

	queryHashCmd := models.Client.ScriptLoad(models.Ctx, queryString)
	QueryHash = queryHashCmd.Val()

	updateHashCmd := models.Client.ScriptLoad(models.Ctx, updateString)
	UpdateHash = updateHashCmd.Val()
}
