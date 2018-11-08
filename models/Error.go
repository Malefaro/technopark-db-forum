package models

import (
	"database/sql"
	"github.com/malefaro/technopark-db-forum/services"
	"log"
)

type Error struct {
	Message string `json:"message"`
}

func checkError(obj interface{}, err error) (interface{}, error) {
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return obj, nil
	default:
		funcname:=services.GetFunctionName()
		log.Printf("Function: %s, Error: %v",funcname , err)
		return nil, err
	}
}