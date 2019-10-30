package dataCrud

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strings"
)

var g *gorm.DB
var tableNames map[string]string

func init() {
	cdn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		Cfg.Name, Cfg.Pass, Cfg.Host, Cfg.Port, Cfg.DBName, Cfg.Charset)
	var err error
	g, err = gorm.Open("mysql", cdn)
	if err != nil {
		log.Fatal("数据库链接失败", err)
	}
	log.Println("数据库初始化")
}

type TablesName struct {
	TableName    string
	TableComment string
}

// 需要 表 Model 和表
type TableModel struct {
	ModelName         string
	TableName         string
	TableComment      string
	SevPackageName    string
	ModelsPackageName string
	TableDesc         []TableDesc
}

type TableDesc struct {
	ColumnName             string
	GolangName             string
	DataType               string
	GolangType             string
	ColumnKey              string
	CharacterMaximumLength int64
	ColumnComment          string
	Tag                    string
	Edit                   string
}

func GetTableNames() (map[string]string, error) {
	var data []TablesName
	fmt.Printf("Cfg.DBName:%s \n", Cfg.DBName)
	err := g.Raw("SELECT `table_name`, table_comment FROM information_schema.tables WHERE table_schema = ?", Cfg.DBName).Find(&data).Error
	tableComment := make(map[string]string)
	for _, tb := range data {
		tableComment[tb.TableName] = tb.TableComment
	}
	return tableComment, err
}

// 获取 表的model 可以获取多个  如果参数为零 表示获取全部的
func GetTableModel(names ...string) ([]TableModel, error) {
	var data []TableModel
	var tables map[string]string
	if len(tableNames) == 0 {
		var err error
		tableNames, err = GetTableNames()
		if err != nil {
			return data, err
		}
	}

	if len(names) > 0 {
		tables = make(map[string]string)
		for _, name := range names {
			comment, ok := tableNames[name]
			if !ok {
				return data, errors.New("错误的数据表名:" + name)
			} else {
				tables[name] = comment
			}
		}
	} else {
		tables = tableNames
	}

	for name, comment := range tables {
		var tableModel TableModel
		modelName, err := GetTableModelName(name)
		if err != nil {
			return data, err
		}
		tableModel.ModelName = modelName
		tableModel.TableName = name
		tableModel.TableComment = comment
		tableModel.SevPackageName = Cfg.SeverPath
		tableModel.ModelsPackageName = Cfg.ModelPath
		desc, err := GetTableDesc(name)
		if err != nil {
			return data, err
		}
		tableModel.TableDesc = desc
		data = append(data, tableModel)
	}
	return data, nil
}

func GetTableModelName(name string) (string, error) {
	if strings.HasPrefix(name, Cfg.Prefix) == false {
		return "", errors.New("数据表前缀和设置的前缀不符")
	}
	name = strings.TrimPrefix(name, Cfg.Prefix)
	return UnderlineToHump(name), nil
}

func (t *TableDesc) GetEdit() {
	if t.ColumnName == "id" {
		t.Edit = "id"
	} else if t.ColumnName == "created_at" || t.ColumnName == "updated_at" || t.ColumnName == "deleted_at" {
		t.Edit = ""
	} else {
		t.Edit = "edit"
	}
}

func (t *TableDesc) GetTag() {
	if t.ColumnName == "deleted_at" {
		t.Tag = `json:"-"`
	} else {
		t.Tag = `json:"` + t.ColumnName + `"`
	}
	if t.ColumnKey == "PRI" {
		t.Tag += ` gorm:"primary_key"`
	}
}

func GetTableDesc(TBName string) ([]TableDesc, error) {
	var data []TableDesc
	err := g.Raw("select column_name,data_type,column_key,character_maximum_length,column_comment from information_schema.columns where table_name = ? and table_schema = ?", TBName, Cfg.DBName).Find(&data).Error
	if err != nil {
		return data, err
	}
	for i, table := range data {
		table.GolangName = getGolangColumnName(table.ColumnName)
		table.GolangType, err = getGolangColumnType(table.ColumnName, table.DataType)
		if err != nil {
			return data, err
		}
		table.GetTag()
		table.GetEdit()
		data[i] = table
	}
	return data, err
}

func getGolangColumnName(name string) string {
	if field, exit := GolangColumnName[name]; exit {
		return field
	}
	return UnderlineToHump(name)
}

func getGolangColumnType(name, TName string) (string, error) {
	if name == "deleted_at" {
		return "*utils.JSONTime", nil
	}
	if field, exit := GolangColumnType[TName]; exit {
		return field, nil
	}
	return "", errors.New("不能识别的数据库字段类型:" + TName)
}

var GolangColumnName = map[string]string{
	"id":         "ID",
	"cus_sn":     "CusSN",
	"user_sn":    "UserSN",
	"created_at": "CreatedAt",
	"updated_at": "UpdatedAt",
	"deleted_at": "DeletedAt",
}

var GolangColumnType = map[string]string{
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"int":        "int32",
	"integer":    "int32",
	"bigint":     "int64",
	"float":      "float32",
	"double":     "float64",
	"decimal":    "float64",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "utils.JSONTime",
	"timestamp":  "utils.JSONTime",
	"char":       "string",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}
