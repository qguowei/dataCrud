package {{.ModelsPackageName}}

import (
	"github.com/jinzhu/gorm"
	"log"
	"newjxc/utils"
)


// {{.TableComment}}
type {{.ModelName}} struct {
    {{range $j,$item := .TableDesc}}
    {{$item.GolangName}}    {{$item.GolangType}}    `{{$item.Tag}}` //{{$item.ColumnComment}} {{end}}
}


func (model *{{.ModelName}}) Update(data map[string]interface{}) error {
	return g.Model(model).Updates(data).Error
}

func (model *{{.ModelName}}) Delete() error {
	return g.Delete(model).Error
}

func Create{{.ModelName}}(model *{{.ModelName}}) error {
	return g.Create(model).Error
}

func Get{{.ModelName}}Info(id int32) *{{.ModelName}} {
	var info {{.ModelName}}
	err := g.Where("id=?", id).First(&info).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Get{{.ModelName}}Info is err:", err)
	}
	return &info
}

func Get{{.ModelName}}ListByCusSN(cus string) (data []{{.ModelName}}) {
	err := g.Where("cus_sn=?", cus).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Get{{.ModelName}}ListByCusSN is err:", err)
	}
	return data
}

func Get{{.ModelName}}List(where []utils.DbWhere, limit, page int) (count int, data []{{.ModelName}}, err error) {
	db := newDbByDbWhere(where)
	err = db.Offset((page - 1) * limit).Limit(limit).Find(&data).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		log.Println("Get{{.ModelName}}List err:", err)
	}
	return count, data, err
}
