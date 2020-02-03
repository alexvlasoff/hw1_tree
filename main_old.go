package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//	"strings"
//	//"path/filepath"
//	//"strings"
//)
//
//const END_SIGN = "└"
//
//func getFromDir(path string, printFiles bool) ([]os.FileInfo, error) {
//	file, err:= os.Open(path)
//	if err != nil {
//		return nil, err
//	}
//
//	sl1, err := file.Readdir(0)
//	if err != nil {
//		return nil, err
//	}
//
//	if !printFiles {
//		var sl []os.FileInfo
//		for _, fileOrDir := range sl1 {
//			if fileOrDir.IsDir(){
//				sl = append(sl, fileOrDir)
//			}
//		}
//		return sl, nil
//	}
//
//	return sl1, nil
//}
//
//
//func recursive(out io.Writer, path string, printFiles bool, level int, prevLevelSign string, prevRepeatPipes string) error  {
//	ent,  err := getFromDir(path, printFiles)
//	if err != nil {
//		return err
//	}
//
//	stopPipe := false
//
//	for i, fileOrDir := range ent {
//		if fileOrDir.IsDir() {
//			message := ""
//			prevSign := ""
//			sign := ""
//
//			if prevLevelSign == END_SIGN{
//				stopPipe = true
//			}
//
//			switch {
//			case i+1==len(ent) && stopPipe:
//				sign = END_SIGN
//				prevSign = " "
//			case i+1==len(ent) && !stopPipe:
//				sign = "\t"+END_SIGN
//				prevSign = "|"
//			default:
//				sign = "├"
//				prevSign = "|"
//			}
//
//			repeatPipes := strings.Repeat(prevSign+"\t", level)
//			if (prevRepeatPipes == "") {
//				prevRepeatPipes = repeatPipes
//			}
//
//
//			switch  {
//			case level==0:
//				message = sign + "───"
//			case level>0:
//				message = prevRepeatPipes+sign+"───"
//			}
//
//			fmt.Fprintln(out, message+fileOrDir.Name())
//			err = recursive(out, path+string(os.PathSeparator)+fileOrDir.Name(), printFiles, level+1, sign, repeatPipes)
//			if err != nil {
//				return err
//			}
//		}
//	}
//
//	return nil
//}
//
//func dirTree(out io.Writer, path string, printFiles bool) error  {
//	return recursive(out, path, printFiles, 0, "", "")
//}
//
//func main() {
//	out := os.Stdout
//	if !(len(os.Args) == 2 || len(os.Args) == 3) {
//		panic("usage go run main.go . [-f]")
//	}
//	path := os.Args[1]
//	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
//	err := dirTree(out, path, printFiles)
//	if err != nil {
//		panic(err.Error())
//	}
//}
