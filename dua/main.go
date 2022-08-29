package main

import (
	"errors"

	"github.com/labstack/gommon/log"
)

func main() {

	err := _r.repository.InsertArea(10, 10, persegi)
	if err != nil {
		log.Error().Msg(err.Error())
		err = errors.New(en.ERROR_DATABASE)
		return err
	}
}
