package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Copyright 2013 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author：polaris	studygolang@gmail.com

func GetProjectRoot() string {
	currentDir, _ := os.Getwd()
	for {
		goModFilePath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModFilePath); os.IsNotExist(err) {
			newPath := filepath.Dir(goModFilePath)
			currentDir = newPath
		} else {
			return currentDir
		}
		if goModFilePath == "go.mod" {
			return currentDir
		}
	}
}

func Welcome() {
	fmt.Println("***********************************")
	fmt.Println("*******欢迎来到Go语言中文网*******")
	fmt.Println("***********************************")
}

// strings.Index的UTF-8版本
// 即 Utf8Index("Go语言中文网", "学习") 返回 4，而不是strings.Index的 8
func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}
