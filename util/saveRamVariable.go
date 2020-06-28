package util

import (
	"encoding/json"
	"expvar"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var ramVarMap sync.Map

const ramVarPath = "./asset/cache/var/"

func SaveRamVar(name string, v expvar.Var) {
	expvar.Publish(name, v)
	ramVarMap.Store(name, v)
}

func RestoreRamVar() {
	fmt.Println("RestoreRamVar")

	s, err := os.Stat(ramVarPath)
	if err != nil {
		fmt.Println("could not find userinfo", err.Error())
		os.Mkdir(ramVarPath, os.ModePerm)
		return
	}

	if !s.IsDir() {
		fmt.Println("userinfo is not a dir")
	}

	filepath.Walk(ramVarPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println("RestoreRamVar.path:" + path)
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("could not open file", err)
		}
		val, ok := ramVarMap.Load(info.Name())
		if !ok {
			fmt.Println("not find", info.Name())
			return nil
		}
		err = json.Unmarshal(b, &val)
		if err != nil {
			fmt.Println("unmarshal faild", err)
			return nil
		}
		return nil
	})
}
