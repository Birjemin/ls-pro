package main

import (
    "github.com/DATA-DOG/go-sqlmock"
    "regexp"
    "testing"
)

var (
    ls = Ls{"664b5f3c4b1cb0a62e1528dcf6c88edb", "test", "This is test"}
)

func TestInsert(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectPrepare(regexp.QuoteMeta("INSERT OR REPLACE INTO LS(ID,NAME,DESC)")).
        ExpectExec().
        WithArgs(ls.ID, ls.Name, ls.Desc).
        WillReturnResult(sqlmock.NewResult(1, 1))

    conn := &LsRepository{db: db}
    err = conn.Insert(ls)
    if err != nil {
        t.Errorf("Insert failed")
    }

}

func TestQuery(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()
    columns := []string{"ID", "NAME", "DESC"}
    mock.ExpectQuery("SELECT (.+) FROM LS").
        WithArgs(ls.ID).
        WillReturnRows(sqlmock.NewRows(columns).AddRow("664b5f3c4b1cb0a62e1528dcf6c88edb", "test", "This is test"))

    conn := &LsRepository{db: db}
    ret, err := conn.GetAll(ls)
    if err != nil {
        t.Errorf("GetAll empty")
        return
    }
    if len(ret) != 1 {
        t.Errorf("GetAll failed")
    }
}

func TestUpdate(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    ls.Desc = "This is new value"

    mock.ExpectPrepare(regexp.QuoteMeta("UPDATE LS SET DESC")).
        ExpectExec().
        WithArgs(ls.Desc, ls.ID, ls.Name).
        WillReturnResult(sqlmock.NewResult(1, 1))

    conn := &LsRepository{db: db}
    err = conn.Update(ls)
    if err != nil {
        t.Errorf("update failed")
        return
    }
}

func TestDel(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM LS WHERE")).
        ExpectExec().
        WithArgs(ls.ID, ls.Name).
        WillReturnResult(sqlmock.NewResult(1, 1))

    conn := &LsRepository{db: db}
    err = conn.Del(ls)
    if err != nil {
        t.Errorf("delete failed")
        return
    }
}
