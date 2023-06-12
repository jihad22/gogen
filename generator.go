package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var repository_tmpl = `
package repository

import (	
	"gorm.io/gorm"
)

type {{.FileNameToCapitalize}}Repository interface {
	Create({{.TableName}} {{.TablePackageName}}.{{.TableNameCapitalize}}) ({{.TablePackageName}}.{{.TableNameCapitalize}}, error)
	FindByID(id string) ({{.TablePackageName}}.{{.TableNameCapitalize}}, error)
	ReadAll() ([]{{.TablePackageName}}.{{.TableNameCapitalize}}, error)
	Update({{.TableNameToLower}} {{.TablePackageName}}.{{.TableNameCapitalize}}) ({{.TablePackageName}}.{{.TableNameCapitalize}}, error)
	Delete(id string) error
}

type {{.FileNameToCapitalize}}RepositoryImpl struct {
	Db *gorm.DB
}

func New{{.FileNameToCapitalize}}RepositoryImpl(Db *gorm.DB) {{.FileNameToCapitalize}}Repository {
	return &{{.FileNameToCapitalize}}RepositoryImpl{Db: Db}
}

func ({{.InitialCharFileName}} {{.FileNameToCapitalize}}RepositoryImpl) Create({{.TableName}} {{.TablePackageName}}.{{.TableNameCapitalize}}) ({{.TablePackageName}}.{{.TableNameCapitalize}}, error) {
	err := {{.InitialCharFileName}}.Db.Model(&{{.TablePackageName}}.{{.TableNameCapitalize}}{}).Create(&{{.TableName}}).Error
	if err != nil {
		return {{.TableName}}, err
	}
	return {{.TableName}}, nil
}

func ({{.InitialCharFileName}} {{.FileNameToCapitalize}}RepositoryImpl) FindByID(id string) ({{.TablePackageName}}.{{.TableNameCapitalize}}, error) {
	//uncomment if you need
	/*id, err := strconv.Atoi(paramId) //convert param to Integer
    if err != nil {
        return {{.TableName}}, err
    }*/

	//change param id in where clause
	var {{.TableName}} {{.TablePackageName}}.{{.TableNameCapitalize}}
	err := {{.InitialCharFileName}}.Db.Model(&{{.TablePackageName}}.{{.TableNameCapitalize}}{}).Where("id = ?", id).First(&{{.TableName}}).Error
	if err != nil {
		return {{.TableName}}, err
	}
	return {{.TableName}}, nil
}

func ({{.InitialCharFileName}} {{.FileNameToCapitalize}}RepositoryImpl) ReadAll() ([]{{.TablePackageName}}.{{.TableNameCapitalize}}, error) {
	var {{.TableName}} []{{.TablePackageName}}.{{.TableNameCapitalize}}
	err := {{.InitialCharFileName}}.Db.Model(&{{.TablePackageName}}.{{.TableNameCapitalize}}{}).Find(&{{.TableName}}).Error
	if err != nil {
		return {{.TableName}}, err
	}
	return {{.TableName}}, nil
}

func ({{.InitialCharFileName}} {{.FileNameToCapitalize}}RepositoryImpl) Update({{.TableName}} {{.TablePackageName}}.{{.TableNameCapitalize}}) ({{.TablePackageName}}.{{.TableNameCapitalize}}, error) {
	//uncomment if you need
	/*id, err := strconv.Atoi(paramId) //convert param to Integer
    if err != nil {
        return {{.TableName}}, err
    }
	//for spesific update
	{{.TableName}} := {{.TablePackageName}}.{{.TableName}}{
		Contoh : {{.TableName}}.Field,
	}
	*/
	//Where("id = ?", {{.TableName}}.Id) change with param you want 
	err := {{.InitialCharFileName}}.Db.Model(&{{.TablePackageName}}.{{.TableNameCapitalize}}{}).Where("id = ?", {{.TableName}}.Id).Updates({{.TableName}}).Error
	if err != nil {
		return {{.TableName}}, err
	}
	return {{.TableName}}, nil
}

func ({{.InitialCharFileName}} {{.FileNameToCapitalize}}RepositoryImpl) Delete(id string) error {
	err := {{.InitialCharFileName}}.Db.Model(&{{.TablePackageName}}.{{.TableNameCapitalize}}{}).Where("id = ?", id).Delete(&{{.TablePackageName}}.{{.TableNameCapitalize}}{}).Error
	if err != nil {
		return err
	}
	return nil
}
`

var controller_tmpl = `
package controller

import (
    "github.com/gin-gonic/gin"
	"net/http"
    "repository"
)

type {{.FileNameToCapitalize}}Controller struct {
    {{.FileName}} repository.{{.FileNameToCapitalize}}Repository
}

func New{{.FileNameToCapitalize}}Controller({{.FileName}} repository.{{.FileNameToCapitalize}}Repository) *{{.FileNameToCapitalize}}Controller {
	return &{{.FileNameToCapitalize}}Controller{ {{.FileName}}: {{.FileName}} }
}

func (controller *{{.FileNameToCapitalize}}Controller) Create(ctx *gin.Context) {
    var {{.TableName}} {{.TablePackageName}}.{{.TableNameCapitalize}}
    if err := ctx.ShouldBindJSON(&{{.TableName}}); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
    }

    {{.TableName}}, err := controller.{{.FileName}}.Create({{.TableName}})
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
    }
    ctx.JSON(http.StatusOK, &{{.TableName}})
}


func (controller *{{.FileNameToCapitalize}}Controller) FindByID(ctx *gin.Context) {
    var {{.TableName}} {{.TablePackageName}}.{{.TableNameCapitalize}}
    id := ctx.Param("id") //change param with you want
    {{.TableName}}, err := controller.{{.FileName}}.FindByID(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
    }
    ctx.JSON(http.StatusOK, &{{.TableName}})
}

func (controller *{{.FileNameToCapitalize}}Controller) ReadAll(ctx *gin.Context) {
    var {{.TableName}} []{{.TablePackageName}}.{{.TableNameCapitalize}}
    {{.TableName}}, err := controller.{{.FileName}}.ReadAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
    }
    ctx.JSON(http.StatusOK, &{{.TableName}})
}

func (controller *{{.FileNameToCapitalize}}Controller) Update(ctx *gin.Context) {
    var {{.TableName}} {{.TablePackageName}}.{{.TableNameCapitalize}}
    if err := ctx.ShouldBindJSON(&{{.TableName}}); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
    }

	//checking data, change param id
	if _, err := controller.{{.FileName}}.FindByID({{.TableName}}.id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

    {{.TableName}}, err := controller.{{.FileName}}.Update({{.TableName}})
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
    }
    ctx.JSON(http.StatusOK, &{{.TableName}})
}

func (controller *{{.FileNameToCapitalize}}Controller) Delete(ctx *gin.Context) {
    id := ctx.Param("id") //change param with you want
    err := controller.{{.FileName}}.Delete(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
}
`

type ComponenName struct {
	FileName             string
	FileNameToCapitalize string
	InitialCharFileName  string
	TablePackageName     string
	TableName            string
	TableNameCapitalize  string
	TableNameToLower     string
	InitialCharTableName string
}

func Generator() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: gogen [file name] [table package] [table name]")
		fmt.Println("Remember : Remember To Move your generated file to absolute directory")
		os.Exit(1)
	}

	fileName := strings.TrimSpace(os.Args[1])
	tablePackage := strings.TrimSpace(os.Args[2])
	tableName := strings.TrimSpace(os.Args[3])
	//ambil direktori saat ini
	outDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data := &ComponenName{
		FileName:             fileName,
		FileNameToCapitalize: strings.Title(fileName),
		InitialCharFileName:  fileName[0:1],
		TablePackageName:     tablePackage,
		TableName:            tableName,
		TableNameCapitalize:  strings.Title(tableName),
		TableNameToLower:     strings.ToLower(tableName),
		InitialCharTableName: tableName[0:1],
	}

	//Compile Template
	tmpl, err := template.New("repository").Parse(repository_tmpl)
	// tmpl, err := template.ParseFiles("repository_tmpl.tpl")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//buat file baru untuk hasil generate
	fName := strings.ToLower(fileName) + ".repository.go"
	outputFilePath := filepath.Join(outDir, fName)

	file, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//render template ke file baru
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s Repository generated to %s.\n", fName, outputFilePath)

	//CONTROLLER GENERATOR

	tmpl, err = template.New("controller").Parse(controller_tmpl)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//buat file baru untuk hasil generate
	fName = strings.ToLower(fileName) + ".controller.go"
	outputFilePath = filepath.Join(outDir, fName)

	file, err = os.Create(outputFilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	//render template ke file baru
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s Controller generated to %s.\n", fName, outputFilePath)
	fmt.Println(">> ===========<<>>========== <<")
	fmt.Println("Remember : Remember To Move your generated file to absolute directory")
}
