package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 2.
const (
	addUser    = `insert into tbUser (idUser, username, password, idKaryawan, status, createdBy, createdOn, updatedBy, updatedOn)values (?,?,?,?,?,?,?,?,?)`
	updateUser = `update tbUser set username=?, password=?, idKaryawan=?, status=?, createdBy=?, createdOn=?, updatedBy=?, updatedOn=? where idUser=?`
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddUser(u User) error {
	//fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		//return err
		fmt.Println(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(addUser, u.ID, u.Username, u.Pwd, u.IDkaryawan, OnAdd, u.CreatedBy, time.Now(), u.UpdatedBy, time.Now())

	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) UpdateUser(u User) error {
	fmt.Println("Update Anda berhasil gan! Silahkan lihat DB Anda.")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateUser, u.Username, u.Pwd, u.IDkaryawan, OnAdd, u.CreatedBy, time.Now(), u.UpdatedBy, time.Now(), u.ID)

	if err != nil {
		return err
	}

	return tx.Commit()
}
