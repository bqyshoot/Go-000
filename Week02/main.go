package main

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

type sth_from_db struct {
}

func dao() (*sth_from_db, error) {
	return nil, sql.ErrNoRows
}

func biz() (*sth_from_db, error) {
	result, err := dao()
	if err != nil {
		return nil, errors.Wrap(err, "find nothing from db")
	}
	return result, nil
}

func service() {
	_, err := biz()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func main() {
	service()
}
