package main

import (
	"fmt"
	"os"
	"sync"
)

// service struct
type service struct {
	id      string
	currDir string
	repo    *LsRepository
}

// insert or update
func (srv *service) Insert(params []string) {
	if len(params) != 4 {
		_, _ = fmt.Fprintf(os.Stderr, `invalid params
Usage: ls-pro -i direction description
`)
		return
	}

	err := srv.repo.Insert(Ls{
		ID:   srv.id,
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

// del
func (srv *service) Del(params []string) {
	if len(params) != 3 {
		_, _ = fmt.Fprintf(os.Stderr, `invalid params
Usage: ls-pro -d direction
`)
		return
	}

	err := srv.repo.Del(Ls{
		ID:   srv.id,
		Name: params[2],
	})

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `invalid params
Usage: ls-pro -d direction
`)
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, `delete success
`)
}

// get all info
func (srv *service) GetAll() {

	var wg sync.WaitGroup

	kvChan, dirChan := make(chan map[string]string), make(chan []string)

	wg.Add(2)
	// find list
	go func() {
		list := srv.kvList()
		wg.Done()
		kvChan <- list
	}()

	// find dir
	go func() {
		dirs, _ := ListDir(srv.currDir)
		wg.Done()
		dirChan <- dirs
	}()

	wg.Wait()

	kv, dirs := <-kvChan, <-dirChan

	if dirs == nil {
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

// get kv list
func (srv *service) kvList() map[string]string {
	lss, err := srv.repo.GetAll(Ls{ID: srv.id})
	if err != nil {
		return map[string]string{}
	}
	ret := make(map[string]string, len(lss))
	for _, v := range lss {
		ret[v.Name] = v.Desc
	}
	return ret
}
