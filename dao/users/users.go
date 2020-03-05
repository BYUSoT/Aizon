package users

import (
	"pml/dao"
	"database/sql"
)

var (
	create, read, delete *sql.Stmt
)

func init() {
	create, err := dao.DB.Prepare("insert into users(passhash, superuser, username, fname, lname, email, staff) values(?, ?, ?, ?, ?, ?, ?);"
	if err != nil {
		log.Fatal(err)
	}
	read, err = dao.DB.Prepare("select * from users where username = ?;")
	if err != nil {
		log.Fatal(err)
	}
	delete, err = dao.DB.Prepare("delete from users where username = ?;")
	if err != nil {
		log.Fatal(err)
	}
}

func Create(User u) {

}