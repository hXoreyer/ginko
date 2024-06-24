package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"
)

//go:embed templates/*
var content embed.FS

type Name struct {
	Name    string
	Version string
}

func createFile(outPath, templateName string, data interface{}) error {
	// 读取模板文件内容
	fileContent, err := content.ReadFile("templates/" + templateName + ".tpl")
	if err != nil {
		return err
	}

	// 解析模板文件内容
	tmpl, err := template.New(templateName).Parse(string(fileContent))
	if err != nil {
		return err
	}

	// 创建输出文件
	outputFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 执行模板并将生成的内容写入输出文件
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		return err
	}
	return nil
}

func Create(baseDir string) error {
	files := []struct {
		Name string
		Path string
	}{
		{Name: "config", Path: "config.yaml"},
		{Name: "mod", Path: "go.mod"},
		{Name: "main", Path: "main.go"},
		{Name: "utils/directory", Path: "utils/directory.go"},
		{Name: "routes/api", Path: "routes/api.go"},
		{Name: "routes/group", Path: "routes/group.go"},
		{Name: "models/user", Path: "models/user.go"},
		{Name: "models/common", Path: "models/common.go"},
		{Name: "middlewares/cors", Path: "middlewares/cors.go"},
		{Name: "internal/user/user", Path: "internal/user/user.go"},
		{Name: "global/app", Path: "global/app.go"},
		{Name: "config/app", Path: "config/app.go"},
		{Name: "config/config", Path: "config/config.go"},
		{Name: "config/database", Path: "config/database.go"},
		{Name: "config/log", Path: "config/log.go"},
		{Name: "common/code/code", Path: "common/code/code.go"},
		{Name: "common/code/en-us", Path: "common/code/en-us.go"},
		{Name: "common/code/zh-cn", Path: "common/code/zh-cn.go"},
		{Name: "common/request/validator", Path: "common/request/validator.go"},
		{Name: "bootstrap/config", Path: "bootstrap/config.go"},
		{Name: "bootstrap/db", Path: "bootstrap/db.go"},
		{Name: "bootstrap/log", Path: "bootstrap/log.go"},
		{Name: "bootstrap/router", Path: "bootstrap/router.go"},
		{Name: "api/user/user", Path: "api/user/user.go"},
	}

	version := runtime.Version()
	version = strings.TrimPrefix(version, "go")
	for _, file := range files {
		templateName := file.Name
		outputPath := filepath.Join(baseDir, file.Path)

		data := Name{
			Name:    baseDir, // 将模板名称作为示例数据传递给模板
			Version: version,
		}

		err := createFile(outputPath, templateName, data)
		if err != nil {
			fmt.Println(err)
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
