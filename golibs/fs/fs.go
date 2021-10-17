// The fs module provides easy and simple access to filesystem functions and other
// things, and acts an addition to the Lua standard library's I/O and fs functions.
package fs

import (
	"strconv"
	"os"
	"strings"

	"hilbish/util"
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
)

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)

	util.Document(L, mod, `The fs module provides easy and simple access to
filesystem functions and other things, and acts an
addition to the Lua standard library's I/O and fs functions.`)

	L.Push(mod)
	return 1
}

func luaErr(L *lua.LState, msg string) {
	L.Error(lua.LString(msg), 2)
}

var exports = map[string]lua.LGFunction{
	"cd": fcd,
	"mkdir": fmkdir,
	"stat": fstat,
	"readdir": freaddir,
}

// cd(dir)
// Changes directory to `dir`
func fcd(L *lua.LState) int {
	path := L.CheckString(1)

	err := os.Chdir(strings.TrimSpace(path))
	if err != nil {
		e := err.(*os.PathError).Err.Error()
		luaErr(L, e)
	}

	return 0
}

// mkdir(name, recursive)
// Makes a directory called `name`. If `recursive` is true, it will create its parent directories.
func fmkdir(L *lua.LState) int {
	dirname := L.CheckString(1)
	recursive := L.ToBool(2)
	path := strings.TrimSpace(dirname)
	var err error

	if recursive {
		err = os.MkdirAll(path, 0744)
	} else {
		err = os.Mkdir(path, 0744)
	}
	if err != nil {
		luaErr(L, err.Error())
	}

	return 0
}

// stat(path)
// Returns info about `path`
func fstat(L *lua.LState) int {
	path := L.CheckString(1)

	pathinfo, err := os.Stat(path)
	if err != nil {
		luaErr(L, err.Error())
		return 0
	}
	statTbl := L.NewTable()
	L.SetField(statTbl, "name", lua.LString(pathinfo.Name()))
	L.SetField(statTbl, "size", lua.LNumber(pathinfo.Size()))
	L.SetField(statTbl, "mode", lua.LString("0" + strconv.FormatInt(int64(pathinfo.Mode().Perm()), 8)))
	L.SetField(statTbl, "isDir", lua.LBool(pathinfo.IsDir()))
	L.Push(statTbl)

	return 1
}

// readdir(dir)
// Returns a table of files in `dir`
func freaddir(L *lua.LState) int {
	dir := L.CheckString(1)
	names := []string{}

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		luaErr(L, err.Error())
		return 0
	}
	for _, entry := range dirEntries {
		names = append(names, entry.Name())
	}

	L.Push(luar.New(L, names))

	return 1
}
