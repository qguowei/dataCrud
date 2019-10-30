package dataCrud

type Config struct {
	Host     string //地址
	Port     int    //端口
	Name     string //名称
	Pass     string //密码
	DBName   string //库名
	Charset  string //编码
	Timezone string //时区
	Prefix   string //前缀

	BasePath  string
	ModelPath string
	SeverPath string
	ApiPath   string
}

var Cfg = Config{
	Host:     "localhost",
	Port:     3306,
	Name:     "root",
	Pass:     "hcs@2017",
	DBName:   "jxc",
	Charset:  "utf8",
	Timezone: "Local",
	Prefix:   "jxc_",

	BasePath:  "../test",
	ModelPath: "models",
	SeverPath: "proSev",
	ApiPath:   "api",
}

const (
	TplModel = "model"
	TplSever = "sever"
	TplApi   = "api"
)

var TplMap = map[string]string{
	"model": "tpl/model.tpl",
	"sever": "tpl/sever.tpl",
	"api":   "tpl/api.tpl",
}
