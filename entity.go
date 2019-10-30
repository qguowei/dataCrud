package dataCrud

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"text/template"
)

func CreateModel(tables ...string) error {
	tmpl, err := getTemplate(TplModel)
	if err != nil {
		return err
	}

	data, err := GetTableModel(tables...)
	if err != nil {
		return err
	}

	// 检查
	basPath, err := GenerateDir(Cfg.BasePath)
	if err != nil {
		return err
	}

	path, err := GenerateDir(basPath + Cfg.ModelPath)
	if err != nil {
		return err
	}

	for _, TableData := range data {
		content := bytes.NewBuffer([]byte{})
		tmpl.Execute(content, TableData)

		fileName := path + TableData.ModelName + ".go"
		if DirOrFileExist(fileName) {
			return errors.New("文件已存在:" + fileName)
		}

		_, err := WriteFile(fileName, content.String())
		if err != nil {
			return err
		}
		fmt.Printf("文件[%s]生成成功 \n", fileName)
	}
	return nil
}

func CreateSever(tables ...string) error {
	tmpl, err := getTemplate(TplSever)
	if err != nil {
		return err
	}

	data, err := GetTableModel(tables...)
	if err != nil {
		return err
	}

	// 检查
	basPath, err := GenerateDir(Cfg.BasePath)
	if err != nil {
		return err
	}

	path, err := GenerateDir(basPath + Cfg.SeverPath)
	if err != nil {
		return err
	}

	for _, TableData := range data {
		content := bytes.NewBuffer([]byte{})
		tmpl.Execute(content, TableData)

		fileName := path + TableData.ModelName + ".go"
		if DirOrFileExist(fileName) {
			return errors.New("文件已存在:" + fileName)
		}

		_, err := WriteFile(fileName, content.String())
		if err != nil {
			return err
		}
		fmt.Printf("文件[%s]生成成功 \n", fileName)
	}
	return nil
}

func CreateApi(tables ...string) error {
	tmpl, err := getTemplate(TplApi)
	if err != nil {
		return err
	}

	data, err := GetTableModel(tables...)
	if err != nil {
		return err
	}

	// 检查
	basPath, err := GenerateDir(Cfg.BasePath)
	if err != nil {
		return err
	}

	path, err := GenerateDir(basPath + Cfg.ApiPath)
	if err != nil {
		return err
	}

	for _, TableData := range data {
		content := bytes.NewBuffer([]byte{})
		tmpl.Execute(content, TableData)

		fileName := path + TableData.ModelName + ".go"
		if DirOrFileExist(fileName) {
			return errors.New("文件已存在:" + fileName)
		}

		_, err := WriteFile(fileName, content.String())
		if err != nil {
			return err
		}
		fmt.Printf("文件[%s]生成成功 \n", fileName)
	}
	return nil
}

func getTemplate(name string) (*template.Template, error) {
	fileName, ok := TplMap[name]
	if !ok {
		return nil, errors.New("模板不存在")
	}
	tplStr, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return template.New(name).Parse(string(tplStr))
}
