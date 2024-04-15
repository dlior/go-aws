package api

import (
	"fmt"
	"lambda/database"
	"lambda/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("username or password cannot be empty")
	}

	isUserExists, err := api.dbStore.IsUserExists(event.Username)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	if isUserExists {
		return fmt.Errorf("user with username %s already exists", event.Username)
	}

	if err := api.dbStore.InsertUser(event); err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
