package main

import (
    "fmt"
    "os"
)

type IService interface {
    GetAll()
    Insert(params []string)
    Del(params []string)
}

type service struct {
    id      string
    currDir string
    repo    *LsRepository
}

func (srv *service) Insert(params []string) {
    if len(params) != 4 {
        _, _ = fmt.Fprintf(os.Stderr, `invalid params
Usage: ls-pro -i direction description
`)
        return
    }

    err := srv.repo.Insert(Ls{
        Id:   srv.id,
        Name: params[2],
        Desc: params[3],
    })

    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, `insert failed
Usage: ls-pro -i direction description
`)
        return
    }

    _, _ = fmt.Fprintf(os.Stderr, `add success
`)
}

func (srv *service) Del(params []string) {
    if len(params) != 3 {
        _, _ = fmt.Fprintf(os.Stderr, `invalid params
Usage: ls-pro -d direction
`)
        return
    }

    err := srv.repo.Del(Ls{
        Id:   srv.id,
        Name: params[2],
    })

    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, `invalid params
Usage: ls-pro -i direction description
`)
        return
    }

    _, _ = fmt.Fprintf(os.Stderr, `delete success
`)
}

func (srv *service) GetAll() {

    // find list
    kv := srv.kvList()
    // find dir
    dirs, err := ListDir(srv.currDir)

    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, `There are no directory
`)
        return
    }
    for _, dir := range dirs {
        if a, ok := kv[dir]; ok {
            _, _ = fmt.Fprintf(os.Stderr, "%-35s%s%s%s\n", dir, string(colorGreen), a, string(colorReset))
        } else {
            _, _ = fmt.Fprintf(os.Stderr, "%s\n", dir)

        }
    }

}

func (srv *service) kvList() map[string]string {
    lss, err := srv.repo.GetAll(Ls{Id: srv.id})
    if err != nil {
        return map[string]string{}
    }
    ret := make(map[string]string, len(lss))
    for _, v := range lss {
        ret[v.Name] = v.Desc
    }
    return ret
}
