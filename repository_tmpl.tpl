package {{.PackageName}}

import (	
	"gorm.io/gorm"
)

type {{.FileNameToCapitalize}}Repository interface {
	Create({{.TableName}} {{.TablePackageName}}.{{.TableName}}) ({{.TablePackageName}}.{{.TableName}}, error)
	FindByID(id string) ({{.TablePackageName}}.{{.TableName}}, error)
	ReadAll() ([]{{.TablePackageName}}.{{.TableName}}, error)
	Update({{.TableNameToLower}} {{.TablePackageName}}.{{.TableName}}) ({{.TablePackageName}}.{{.TableName}}, error)
	Delete(id string) error
}

type {{.FileNameToCapitalize}}RepositoryImpl struct {
	Db *gorm.DB
}

func New{{.FileNameToCapitalize}}RepositoryImpl(Db *gorm.DB) {{.FileName}} {
	return &{{.FileNameToCapitalize}}RepositoryImpl{Db: Db}
}

func ({{.InitialCharFileName}} {{.FileNameToCapitalize}}RepositoryImpl) Create({{.TableName}} {{.TablePackageName}}.{{.TableName}}) ({{.TablePackageName}}.{{.TableName}}, error) {
	err := {{.InitialCharFileName}}.Db.Create(&{{.TableName}}).Error
	if err != nil {
		return {{.TableName}}, err
	}
	return {{.TableName}}, nil
}

func ({{.InitialCharFileName}} {{.FileNameToCapitalize}}RepositoryImpl) FindByID(id string) ({{.TablePackageName}}.{{.TableName}}, error) {
	//uncomment if you need
	/*id, err := strconv.Atoi(paramId) //convert param to Integer
    if err != nil {
        return {{.TableNameToLower}}, err
    }*/

	var {{.TableNameToLower}} []{{.TablePackageName}}.{{.FileName}}
	err := {{.InitialCharFileName}}.Db.Where("id = ?", id).Find(&TableNameToLower).Error
	if err != nil {
		return {{.TableNameToLower}}, err
	}
	return {{.TableNameToLower}}, nil
}


func ({{.InitialCharFileName}} {{.FileNameToCapitalize}}RepositoryImpl) ReadAll() ([]{{.TablePackageName}}.{{.TableName}}, error) {
	var {{.TableNameToLower}} []{{.TablePackageName}}.{{.FileName}}
	err := {{.InitialCharFileName}}.Db.Find(&{{.TableNameToLower}}).Error
	if err != nil {
		return {{.TableNameToLower}}, err
	}
	return {{.TableNameToLower}}, nil
}

func ({{.InitialCharFileName}} {{.FileNameToCapitalize}}RepositoryImpl) Update({{.TableNameToLower}} {{.TablePackageName}}.{{.TableName}}) ({{.TablePackageName}}.{{.TableName}}, error) {
	//uncomment if you need
	/*id, err := strconv.Atoi(paramId) //convert param to Integer
    if err != nil {
        return {{.TableNameToLower}}, err
    }
	//for spesific update
	{{.TableNameToLower}} := {{.TablePackageName}}.{{.TableName}}{
		Contoh : {{.TableNameToLower}}.Field,
	}
	*/
	//Where("id = ?", {{.TableNameToLower}}.Id) change with param you want 
	err := {{.InitialCharFileName}}.Db.Model(&{{.TableName}}).Where("id = ?", {{.TableNameToLower}}.Id).Updates({{.TableNameToLower}}).Error
	if err != nil {
		return {{.TableName}}, err
	}
	return {{.TableName}}, nil
}

func ({{.InitialCharFileName}} {{.FileNameToCapitalize}}RepositoryImpl) Delete(id string) error {
	var {{.TableNameToLower}} {{.TablePackageName}}.{{.FileName}}
	err := {{.InitialCharFileName}}.Db.Where("id = ?", id).Delete(&{{.TableNameToLower}}).Error
	if err != nil {
		return err
	}
	return nil
}