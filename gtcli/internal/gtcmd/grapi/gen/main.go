package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a path as an argument")
	}
	grapiDir := filepath.FromSlash(os.Args[1])
	var grapiData []map[string]any
	_ = filepath.Walk(grapiDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if ok := skipPath(path); !ok {
			relativePath := strings.ReplaceAll(path, grapiDir, "")
			if info.IsDir() {
				grapiData = append(grapiData, map[string]any{
					"path":    relativePath,
					"isDir":   true,
					"content": []byte(""),
				})

			} else {
				content, _ := os.ReadFile(path)
				grapiData = append(grapiData, map[string]any{
					"path":    relativePath,
					"isDir":   false,
					"content": content,
				})
			}
		}
		return nil
	})
	workDir, _ := os.Getwd()
	grapiJsonPath := filepath.Join(filepath.Dir(workDir), "newtpl.go")
	grapiJson, _ := json.Marshal(grapiData)
	err := os.WriteFile(
		grapiJsonPath,
		[]byte(fmt.Sprintf("package grapi\n\nconst grapiJson = `%s`\n", grapiJson)),
		fs.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("grapiJson written to: ", grapiJsonPath)
}

func skipPath(path string) bool {
	re := regexp.MustCompile(`.*\.git($|[/\\])|.*\.idea($|[/\\])|.*\.vscode($|[/\\])|.*\.exe$|.*\.log$|.*\.db$`)
	if re.MatchString(path) {
		return true
	}
	return false
}
