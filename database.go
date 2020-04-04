package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "os"
)

type Ls struct {
    Id   string
    Name string
    Desc string
}

func CreateConnection(filepath string) (*sql.DB, error) {
    // check db
    if !Exist(filepath) {
        _, _ = fmt.Fprintf(os.Stderr, `database is not exist,creating the databse...
`)
        if err := createDb(filepath); err != nil {
            log.Fatal(err.Error())
        }
    }

    db, err := sql.Open("sqlite3", filepath)
    if err != nil {
        log.Fatal(err.Error())
    }

    if err = createTb(db); err != nil {
        return nil, err
    }

    return db, nil
}

func createDb(filepath string) error {
    file, err := os.Create(filepath) // Create SQLite file
    if err != nil {
        return err
    }
    defer file.Close()
    return nil
}

func createTb(db *sql.DB) error {
    stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS ls(ID CHAR(30),NAME TEXT,DESC TEXT);CREATE INDEX IF NOT EXISTS idx ON ls(ID,NAME);`)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()
    _, err = stmt.Exec()
    return err
}
