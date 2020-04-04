package main

import (
    "log"
    "testing"
)

var (
    id      = "664b5f3c4b1cb0a62e1528dcf6c88edb"
    currDir = "/Users/birjemin/Developer/Go/src/ls-pro"
)

func getSrv() IService {
    db, err := CreateConnection("ls_pro.db")
    if err != nil {
        log.Fatal(err)
    }

    return &service{
        repo:    &LsRepository{db: db},
        id:      id,
        currDir: currDir,
    }
}

func TestInsertSrv(t *testing.T) {
    srv := getSrv()
    if srv == nil {
        t.Errorf("Connect the Db faield")
        return
    }
    srv.Insert([]string{
        "./ls-pro", "-a", "srv", "啊哈哈",
    })
}

func TestGetAllSrv(t *testing.T) {
    srv := getSrv()
    if srv == nil {
        t.Errorf("Connect the Db faield")
        return
    }
    srv.GetAll()
}

func TestDelSrv(t *testing.T) {
    srv := getSrv()
    if srv == nil {
        t.Errorf("Connect the Db faield")
        return
    }
    srv.Del([]string{
        "./ls-pro", "-d", "srv",
    })
}
