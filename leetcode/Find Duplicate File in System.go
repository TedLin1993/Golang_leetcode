package leetcode

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	fileContentMap := map[string][]string{}
	for i := 0; i < len(paths); i++ {
		path := strings.Split(paths[i], " ")
		directoryPath := path[0]
		for j := 1; j < len(path); j++ {
			idx := strings.Index(path[j], "(")
			fileName := path[j][:idx]
			content := path[j][idx+1 : len(path[j])-1]
			filePath := fmt.Sprintf("%s/%s", directoryPath, fileName)
			fileContentMap[content] = append(fileContentMap[content], filePath)
		}
	}
	res := [][]string{}
	for _, paths := range fileContentMap {
		if len(paths) > 1 {
			res = append(res, paths)
		}
	}
	return res
}

func Test_findDuplicate() {
	fmt.Println(findDuplicate([]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"})) //[[root/a/1.txt root/c/3.txt] [root/a/2.txt root/c/d/4.txt root/4.txt]]
	fmt.Println(findDuplicate([]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"}))                     //[[root/a/1.txt root/c/3.txt] [root/a/2.txt root/c/d/4.txt]]
	fmt.Println(findDuplicate([]string{"root/a 1.txt(abcd) 2.txt(efsfgh)", "root/c 3.txt(abdfcd)", "root/c/d 4.txt(efggdfh)"}))              //[]
}
