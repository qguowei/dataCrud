package {{.SevPackageName}}

import (
	"newjxc/{{.ModelsPackageName}}"
	"newjxc/utils"
	"strconv"
)

// 获取药品基本信息
type {{.ModelName}}InfoInput struct {
	ID    int32
	CusSN string
}

func (d *{{.ModelName}}InfoInput) Validate() error {
	if d.ID == 0 {
		return utils.NewCodesNotFont("id不能为空")
	}
	if d.CusSN == "" {
		return utils.NewCodesInternalServer("CusSN不能为空")
	}
	return nil
}

func {{.ModelName}}Info(in *{{.ModelName}}InfoInput) (info *{{.ModelsPackageName}}.{{.ModelName}}, err error) {
	if err := in.Validate(); err != nil {
		return info, err
	}
	info = {{.ModelsPackageName}}.Get{{.ModelName}}Info(in.ID)
	if info.ID == 0 {
		return info, utils.NewCodesNotFont("没有找到你要查询的信息")
	}
	if info.CusSN != in.CusSN {
		return &{{.ModelsPackageName}}.{{.ModelName}}{}, utils.NewCodesForbidden("你无权查看该数据")
	}
	return info, nil
}

// 获取药品列表信息
type {{.ModelName}}ListInput struct {
	CusSN string
	Page  int
	Limit int
	Where []utils.DbWhere
}

func (e *{{.ModelName}}ListInput) Validate() error {
	if e.CusSN == "" {
		return utils.NewCodesInternalServer("CusSN不能为空")
	}
	return nil
}

type {{.ModelName}}ListOut struct {
	Data  []{{.ModelsPackageName}}.{{.ModelName}} `json:"data"`
	Count int           `json:"count"`
}

func {{.ModelName}}List(in *{{.ModelName}}ListInput) (data *{{.ModelName}}ListOut, err error) {
	if err := in.Validate(); err != nil {
		return data, err
	}

	in.Where = append(in.Where, utils.DbWhere{
		FType:  "eq",
		FName:  "cus_sn",
		FValue: in.CusSN,
	})
	data = &{{.ModelName}}ListOut{}
	data.Count, data.Data, err = {{.ModelsPackageName}}.Get{{.ModelName}}List(in.Where, in.Limit, in.Page)
	return data, err
}

// 药品信息添加
type {{.ModelName}}CreateInput struct {
{{range .TableDesc}}
{{if eq .Edit "edit"}}{{.GolangName}}    {{.GolangType}}    `{{.Tag}}` //{{.ColumnComment}}{{end}}{{end}}
}

func (d *{{.ModelName}}CreateInput) Validate() error {
	if d.CusSN == "" {
		return utils.NewCodesInternalServer("CusSN不能为空")
	}
	{{range .TableDesc}}{{if eq .GolangType "string"}}
	if utf8.RuneCountInString(d.{{.GolangName}}) > {{.CharacterMaximumLength}} {
        return utils.NewCodesInvalidData("{{.ColumnComment}}最大长度为{{.CharacterMaximumLength}}")
    }{{end}}{{end}}
	return nil
}

func {{.ModelName}}Create(in *{{.ModelName}}CreateInput) (model *{{.ModelsPackageName}}.{{.ModelName}}, err error) {
	if err := in.Validate(); err != nil {
		return model, err
	}
	model = &{{.ModelsPackageName}}.{{.ModelName}}{}
    {{range .TableDesc}}
    {{if eq .Edit "edit"}}model.{{.GolangName}} = in.{{.GolangName}}{{end}}{{end}}
	err = {{.ModelsPackageName}}.Create{{.ModelName}}(model)
	return model, err
}

// 药品信息修改
type {{.ModelName}}UpdateInput struct {
{{range .TableDesc}}{{if eq .Edit "id"}}{{.GolangName}}    {{.GolangType}}    `{{.Tag}}` //{{.ColumnComment}} {{end}}
{{if eq .Edit "edit"}}{{.GolangName}}    {{.GolangType}}    `{{.Tag}}` //{{.ColumnComment}} {{end}}{{end}}
}

func (d *{{.ModelName}}UpdateInput) Validate() error {
	if d.ID == 0 {
		return utils.NewCodesNotFont("id不能为空")
	}
	if d.CusSN == "" {
		return utils.NewCodesInternalServer("CusSN不能为空")
	}
    {{range .TableDesc}}{{if eq .GolangType "string"}}
    if utf8.RuneCountInString(d.{{.GolangName}}) > {{.CharacterMaximumLength}} {
        return utils.NewCodesInvalidData("{{.ColumnComment}}最大长度为{{.CharacterMaximumLength}}")
    }{{end}}{{end}}
	return nil
}

func {{.ModelName}}Update(in *{{.ModelName}}UpdateInput) (model *{{.ModelsPackageName}}.{{.ModelName}}, err error) {
	if err := in.Validate(); err != nil {
		return model, err
	}
	model = {{.ModelsPackageName}}.Get{{.ModelName}}Info(in.ID)
	if model.ID == 0 {
		return model, utils.NewCodesNotFont("部门为空")
	}
	if model.CusSN != in.CusSN {
		return &{{.ModelsPackageName}}.{{.ModelName}}{}, utils.NewCodesForbidden("你无权修改该部门")
	}
	data := map[string]interface{}{
        {{range .TableDesc}}{{if eq .Edit "id"}}"{{.GolangName}}":in.{{.GolangName}},{{end}}
        {{if eq .Edit "edit"}}"{{.GolangName}}":in.{{.GolangName}},{{end}}{{end}}
	}
	err = model.Update(data)
	return model, err
}

// 药品信息删除
func {{.ModelName}}Delete(in *{{.ModelName}}InfoInput) (out *utils.DeleteOut, err error) {
	if err := in.Validate(); err != nil {
		return out, err
	}
	dep := {{.ModelsPackageName}}.Get{{.ModelName}}Info(in.ID)
	if dep.ID == 0 {
		return out, utils.NewCodesNotFont("部门为空")
	}
	if dep.CusSN != in.CusSN {
		return out, utils.NewCodesForbidden("你无权删除该部门")
	}
	err = dep.Delete()
	return out, err
}
