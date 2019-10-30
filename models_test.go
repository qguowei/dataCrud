package dataCrud

import (
	"fmt"
	"testing"
)

func TestGetTableNames(t *testing.T) {
	data, err := GetTableNames()
	if err != nil {
		t.Fatalf("TestGetTableNames err:%v", err)
	}
	fmt.Printf("data:%+v \n err:%v \n", data, err)
}

func TestGetTableDesc(t *testing.T) {
	data, err := GetTableDesc("jxc_customer")
	if err != nil {
		t.Fatalf("TestGetTableDesc err:%v", err)
	}
	fmt.Printf("data:%+v \n err:%v \n", data, err)
}

func TestGetTableModel(t *testing.T) {
	data, err := GetTableModel()
	if err != nil {
		t.Fatalf("TestGetTableDesc err:%v", err)
	}
	for _, d := range data {
		fmt.Printf("model:%s 表名：%s  表注释:%s \n", d.ModelName, d.TableName, d.TableComment)
		for _, desc := range d.TableDesc {
			fmt.Printf("%v \n", desc)
		}
		fmt.Println()
	}
}
