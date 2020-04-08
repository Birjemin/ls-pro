package main

import (
    "github.com/DATA-DOG/go-sqlmock"
    "regexp"
    "testing"
)

var (
    id      = "664b5f3c4b1cb0a62e1528dcf6c88edb"
    currDir = "/Users/birjemin/Developer/Go/src/ls-pro"
)

func TestInsertSrv(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectPrepare(regexp.QuoteMeta("INSERT OR REPLACE INTO LS(ID,NAME,DESC)")).
        ExpectExec().
        WithArgs(ls.Id, ls.Name, ls.Desc).
        WillReturnResult(sqlmock.NewResult(1, 1))

    srv := &service{
        repo:    &LsRepository{db: db},
        id:      id,
        currDir: currDir,
    }

    srv.Insert([]string{
        "./ls-pro", "-a", "srv", "啊哈哈",
    })
}

func TestGetAllSrv(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()
    columns := []string{"ID", "NAME", "DESC"}
    mock.ExpectQuery("SELECT (.+) FROM LS").
        WithArgs(ls.Id).
        WillReturnRows(sqlmock.NewRows(columns).AddRow("664b5f3c4b1cb0a62e1528dcf6c88edb", "test", "This is test"))

    srv := &service{
        repo:    &LsRepository{db: db},
        id:      id,
        currDir: currDir,
    }

    srv.GetAll()
}

func TestDelSrv(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM LS WHERE")).
        ExpectExec().
        WithArgs(ls.Id, ls.Name).
        WillReturnResult(sqlmock.NewResult(1, 1))

    srv := &service{
        repo:    &LsRepository{db: db},
        id:      id,
        currDir: currDir,
    }

    srv.Del([]string{
        "./ls-pro", "-d", "srv",
    })
}
