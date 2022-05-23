package leetcode

import (
	"fmt"
)

func simplifyPath(path string) string {
	var res string
	var splitPath []string
	startIndex := 0
	for i := 0; i < len(path); i++ {
		if path[i] != '/' {
			continue
		}
		addElement2SplitPath(&splitPath, path, startIndex, i)
		startIndex = i + 1
	}
	if startIndex < len(path) {
		addElement2SplitPath(&splitPath, path, startIndex, len(path))
	}

	for _, s := range splitPath {
		res += "/" + s
	}

	if res == "" {
		res = "/"
	}
	return res
}

func addElement2SplitPath(splitPath *[]string, path string, startIndex int, currentIndex int) {
	tempElement := path[startIndex:currentIndex]

	if len(tempElement) == 0 || tempElement == "." {
		return
	} else if tempElement == ".." {
		if len(*splitPath) == 0 {
			return
		}
		*splitPath = (*splitPath)[:len(*splitPath)-1]
	} else {
		*splitPath = append(*splitPath, tempElement)
	}
}

func TestsimplifyPath() {
	fmt.Println(simplifyPath("/home/"))                // /home
	fmt.Println(simplifyPath("/home/./"))              // /home
	fmt.Println(simplifyPath("/../"))                  // /
	fmt.Println(simplifyPath("/home/../"))             // /
	fmt.Println(simplifyPath("/home//foo/"))           // /home/foo
	fmt.Println(simplifyPath("/a/./b/../../c/"))       // /c
	fmt.Println(simplifyPath("/a//b////c/d//././/..")) // "/a/b/c"
}
