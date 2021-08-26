package main
//gopath设置为当前的工程路径，相对路径不能出现空格
import (
	lib1 "code/importPkg/lib1"
	lib2 "code/importPkg/lib2"
)

func main() {

	lib1.Lib1show()
	lib2.Lib2show()
}
