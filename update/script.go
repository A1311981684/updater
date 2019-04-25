package update

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)
var PreExecuteScripts []string
var PostExecuteScripts []string

//There are pre-execute and post-execute scripts
//pre_1.sh, pre_2.sh, post_1.sh, post_2.sh
func loadScripts() error {
	//Checkout if scripts exist
	infos, err := ioutil.ReadDir(newFilePath)
	if err != nil {
		return err
	}
	if len(infos) == 0 {
		return nil
	}
	var scriptList []string
	for _, v := range infos {
		if !v.IsDir() {
			if strings.Contains(v.Name(), ".sh") {
				scriptList = append(scriptList, newFilePath + v.Name())
			}
		}
	}
	//Classify them into pre and post type
	var pre []string
	var post []string
	for _, v := range scriptList{
		if strings.Contains(filepath.Base(v), "pre") {
			pre = append(pre, v)
		}else if strings.Contains(filepath.Base(v), "post") {
			post = append(post, v)
		} else {
			log.Println("Unknown scripts:", filepath.Base(v))
		}
	}
	//Scripts want to be executed in a sequence, need to be sorted
	length := len(pre)
	var preSorted = make([]string, length)
	for _, v := range pre {
		scriptName := filepath.Base(v)
		split := strings.Split(strings.Replace(scriptName, ".sh", "", 1), "_")
		if len(split) != 2 {
			log.Println("Invalid script name:", scriptName, "haha",  split)
			length--
			preSorted = make([]string, length)
			continue
		}
		number, err := strconv.Atoi(split[1])
		if err != nil {
			log.Println("Invalid script name:", scriptName)
			length--
			preSorted = make([]string, length)
			continue
		}
		preSorted[number-1] = v
	}

	var postSorted = make([]string, len(post))
	length = len(post)
	for _, v := range post {
		scriptName := filepath.Base(v)
		split := strings.Split(strings.Replace(scriptName, ".sh", "", 1), "_")
		if len(split) != 2 {
			log.Println("Invalid script name:", scriptName)
			length--
			postSorted = make([]string, length)
			continue
		}
		number, err := strconv.Atoi(split[1])
		if err != nil {
			log.Println("Invalid script name:", scriptName)
			length--
			postSorted = make([]string, length)
			continue
		}
		postSorted[number-1] = v
	}

	PreExecuteScripts = preSorted
	PostExecuteScripts = postSorted
	return nil
}

func executeScript(script string)error{
	log.Println("Executing:", script)
	//cmd := exec.Command("/bin/bash", "-c", script )
	return nil
}

func executePre()error{
	if len(PreExecuteScripts) == 0 {
		return nil
	}
	for _, v := range PreExecuteScripts {
		err := executeScript(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func executePost()error {
	if len(PostExecuteScripts) == 0 {
		return nil
	}

	for _, v :=range PostExecuteScripts {
		err := executeScript(v)
		if err != nil {
			return nil
		}
	}

	return nil
}