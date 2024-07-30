package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/yosssi/gohtml"
)

func execute() {
	dir := getDirectory()

	if len(os.Args) > 1 && (os.Args[1] == "watch" || os.Args[1] == "w") {
		watchAndCompile(dir)
	} else {
		compileOnce(dir)
	}
}

func compileOnce(dir string) {
	dictionary, err := BuildFileDictionary(dir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	indexContent, exists := dictionary["index"]
	if !exists {
		fmt.Println("Error: index.qtml not found")
		os.Exit(1)
	}

	compiledContent := compileHtml(indexContent, dictionary)
	writeToFile(dir, compiledContent)
	fmt.Println("Compilation successful.")
}

func getDirectory() string {
	var dir string
	if len(os.Args) < 2 || (len(os.Args) == 2 && os.Args[1] == "watch") {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}
	} else {
		dir = os.Args[1]
	}
	return dir
}

func writeToFile(dir, content string) {
	outputPath := filepath.Join(dir, "index.html")
	prettyHTML := gohtml.Format(content)
	err := os.WriteFile(outputPath, []byte(prettyHTML), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		os.Exit(1)
	}
}

func watchAndCompile(dir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					compileOnce(dir)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Watching directory: %s\n", dir)
	<-done
}
