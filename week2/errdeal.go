package main

import (
	"database/sql"
	"fmt"
	"log"
	xerrors "github.com/pkg/errors"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func ExecuteSql() error {
	return xerrors.Wrap(sql.ErrNoRows,"具体sql")
}

func main() {
	err := ExecuteSql()
	if xerrors.Cause(err) == sql.ErrNoRows {
		fmt.Printf(err.Error())
		log.Printf(err.Error())
	}
}