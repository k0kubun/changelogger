package main

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var contentByPath map[string]string

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	contentByPath = map[string]string{}
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return nil
		}

		if info.IsDir() {
			startLogging(path)
		} else {
			buffer, _ := ioutil.ReadFile(path)
			contentByPath[path] = string(buffer)
		}
		return nil
	})

	for {
		runtime.Gosched()
	}
}

func startLogging(path string) {
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		err = watcher.Watch(path)
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		for {
			select {
			case event := <-watcher.Event:
				if event.IsModify() {
					modifiedPath := event.Name
					checkContentChanged(modifiedPath)
				}
			case err := <-watcher.Error:
				log.Println("Error:", err)
			}
		}
	}()
}

func checkContentChanged(path string) bool {
	buffer, _ := ioutil.ReadFile(path)
	newContent := string(buffer)
	oldContent := contentByPath[path]

	if oldContent != newContent {
		log.Println("Changed:", path)
		showDiff(oldContent, newContent)
		fmt.Println("")

		contentByPath[path] = newContent
		return true
	}
	return false
}
