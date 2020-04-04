package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

type IRepository interface {
    GetAll(ls Ls) ([]Ls, error)
    Insert(ls Ls) error
    Update(ls Ls) error
    Del(ls Ls) error
}

type LsRepository struct {
    db *sql.DB
}

func (l *LsRepository) GetAll(ls Ls) ([]Ls, error) {
    lsSql := fmt.Sprintf("select * from ls where id ='%s'", ls.Id)
    rows, err := l.db.Query(lsSql)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var lss []Ls
    for rows.Next() {
        ls := Ls{}
        _ = rows.Scan(&ls.Id, &ls.Name, &ls.Desc)
        lss = append(lss, ls)
    }
    return lss, nil
}

func (l *LsRepository) Insert(ls Ls) error {
    return l.call(`INSERT OR REPLACE INTO LS(ID,NAME, DESC) VALUES(?, ?,?)`, ls.Id, ls.Name, ls.Desc)
}

func (l *LsRepository) Update(ls Ls) error {
    return l.call(`UPDATE LS SET DESC=? WHERE ID=? AND NAME=?`, ls.Desc, ls.Id, ls.Name)
}

func (l *LsRepository) Del(ls Ls) error {
    return l.call(`DELETE FROM LS WHERE ID=? AND NAME=?`, ls.Id, ls.Name)
}

func (l *LsRepository) call(sql string, str ...interface{}) error {
    stmt, err := l.db.Prepare(sql)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(str...)
    return err
}
