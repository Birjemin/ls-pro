package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// LsRepository ls repository
type LsRepository struct {
	db *sql.DB
}

// GetAll get all
func (l *LsRepository) GetAll(ls Ls) ([]Ls, error) {
	lsSQL := "SELECT * FROM LS WHERE ID =?"
	rows, err := l.db.Query(lsSQL, ls.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var lss []Ls
	for rows.Next() {
		ls := Ls{}
		_ = rows.Scan(&ls.ID, &ls.Name, &ls.Desc)
		lss = append(lss, ls)
	}
	return lss, nil
}

// Insert insert or update
func (l *LsRepository) Insert(ls Ls) error {
	return l.call(`INSERT OR REPLACE INTO LS(ID,NAME,DESC) VALUES(?,?,?)`, ls.ID, ls.Name, ls.Desc)
}

// Update update
func (l *LsRepository) Update(ls Ls) error {
	return l.call(`UPDATE LS SET DESC=? WHERE ID=? AND NAME=?`, ls.Desc, ls.ID, ls.Name)
}

// Del del
func (l *LsRepository) Del(ls Ls) error {
	return l.call(`DELETE FROM LS WHERE ID=? AND NAME=?`, ls.ID, ls.Name)
}

// call function
func (l *LsRepository) call(sql string, str ...interface{}) error {
	stmt, err := l.db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(str...)
	return err
}
