package main

import (
    "log"
    "testing"
)

var (
    ls = Ls{"664b5f3c4b1cb0a62e1528dcf6c88edb", "test", "This is test"}
)

func getConn() IRepository {
    db, err := CreateConnection("ls_pro.db")
    if err != nil {
        log.Fatal(err)
        return nil
    }
    return &LsRepository{db: db}
}

func TestInsert(t *testing.T) {
    conn := getConn()
    if conn == nil {
        t.Errorf("Connect the Db faield")
        return
    }
    err := conn.Insert(ls)
    if err != nil {
        t.Errorf("Insert failed")
        return
    }
}

func TestQuery(t *testing.T) {
    conn := getConn()
    if conn == nil {
        t.Errorf("Connect the Db faield")
        return
    }
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
    conn := getConn()
    if conn == nil {
        t.Errorf("Connect the Db faield")
        return
    }
    ls.Desc = "This is new value"
    err := conn.Update(ls)
    if err != nil {
        t.Errorf("update failed")
        return
    }
}

func TestDel(t *testing.T) {
    conn := getConn()
    if conn == nil {
        t.Errorf("Connect the Db faield")
        return
    }
    err := conn.Del(ls)
    if err != nil {
        t.Errorf("delete failed")
        return
    }
}



