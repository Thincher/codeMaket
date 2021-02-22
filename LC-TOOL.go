package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var flag int32

//"codeTop.txt","codeTop_done.txt"
// "必会题.txt","LC-Done.txt"
// "排序设计.txt","LC-Done.txt"
func main() {
	tool("codeTop.txt","codeTop_done.txt")
}

func tool(path,donePath string) {
	//os.Create("必会题.txt")
	//os.Create("LC-Done.txt")
	tobe, _ := os.OpenFile(path, os.O_RDONLY, 0600)
	done, _ := os.OpenFile(donePath, os.O_RDONLY, 0600)
	tobeContent, _ := ioutil.ReadAll(tobe)
	donecontent, _ := ioutil.ReadAll(done)

	tobeArr := strings.Split(string(tobeContent), ",")
	doneArr := strings.Split(string(donecontent), ",")
	doneMap := make(map[string]bool)
	tobeArr = RemoveRepeatedElement(tobeArr[:len(tobeArr)-1])
	doneArr = doneArr[:len(doneArr)-1]
	for _, val := range doneArr {
		doneMap[val] = true
	}
	start := time.Now()
	fmt.Println("还剩这么多了", len(tobeArr)-len(doneMap))
	if len(tobeArr)-len(doneArr) < 3 {
		for _, val := range tobeArr {
			if !doneMap[val] {
				fmt.Println(val)
				os.Exit(0)
			}

		}
		fmt.Println("做完了")
		os.Exit(0)
	}
	for {
		r := rand.New(rand.NewSource(time.Now().Unix()))

		len := len(tobeArr)
		cur := tobeArr[r.Intn(len)]
		if !doneMap[cur] {
			fmt.Print(cur)
			break
		}

		duration := time.Since(start)
		if duration.Seconds() >= 3 {
			fmt.Println("超时辣,请重试")
			break
		}

	}

	defer func() {
		done.Close()
		tobe.Close()
	}()
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
