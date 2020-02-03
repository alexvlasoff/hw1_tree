package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

const EndSign = "└───"
const DirSign = "├───"
const PipeSign = "│\t"

type ByName []os.FileInfo

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].Name() < n[j].Name() }

func getFromDir(path string, printFiles bool) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if !printFiles {
		var dirs []os.FileInfo
		for _, fileOrDir := range files {
			if fileOrDir.IsDir(){
				dirs = append(dirs, fileOrDir)
			}
		}
		sort.Sort(ByName(dirs))
		return dirs, nil
	}

	sort.Sort(ByName(files))
	return files, nil
}


func recursive(out io.Writer, path string, printFiles bool, level int, ends map[int]int, prevLevelSign string) error  {
	entities,  err := getFromDir(path, printFiles)
	if err != nil {
		return err
	}

	for i, fileOrDir := range entities {

		message := ""
		size := ""
		sign := ""

		switch {
		case i+1==len(entities):
			sign = EndSign
		default:
			sign = DirSign
		}

		if prevLevelSign == EndSign {
			ends[level] = 1
		} else {
			ends[level] = 0
		}

		if len(ends)>level+1 {
			for i:=len(ends); i>level; i-- {
				delete(ends, i)
			}
		}

		for idx:=0; idx<len(ends); idx++  {
			if idx == 0 {
				continue
			}

			if ends[idx] == 0 {
				message += PipeSign
			} else {
				message += "\t"
			}
		}

		message += sign

		if !fileOrDir.IsDir() {
			size = " (empty)"
			if fileOrDir.Size() > 0 {
				size = " ("+ strconv.Itoa(int(fileOrDir.Size())) + "b)"
 			}
		}

		_, err = fmt.Fprintln(out, message+fileOrDir.Name()+size)
		if err != nil {
			return err
		}

		if fileOrDir.IsDir() {
			err = recursive(out, path+string(os.PathSeparator)+fileOrDir.Name(), printFiles, level+1, ends, sign)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error  {
	return recursive(out, path, printFiles, 0, make(map[int]int), "")
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
