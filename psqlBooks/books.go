package psqlbooks

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CreateABook(ctx echo.Context, db *sql.DB, isbn string, name string, author string, edition string, publication string) {
	query := fmt.Sprintf(`insert into "Books"("ISBN", "Name", "Author", "Edition", "Publication") values(%s, %s, %s, %s, %s)`, isbn, name, author, edition, publication)
	_, err := db.Exec(query)
	checkError(err, query)
}

func checkError(err error, query string) {
	if err != nil {
		log.Errorf("got error:%s, when trying to execute query:%s", err.Error(), query)
	}
}
