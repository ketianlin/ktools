package example

import (
	"encoding/json"
	"fmt"
	"github.com/ketianlin/ktools/excel"
	"github.com/xuri/excelize/v2"
	"testing"
)

type employee struct {
	Id       string `excel:"name:用户账号;"`
	Name     string `excel:"name:用户姓名;index:1;"`
	Email    string `excel:"name:用户邮箱;width:25;"`
	Com      string `excel:"name:所属公司;"`
	Dept     string `excel:"name:所在部门;"`
	RoleKey  string `excel:"name:角色代码;"`
	RoleName string `excel:"name:角色名称;replace:1_超级管理员,2_普通用户;"`
	Remark   string `excel:"name:备注;width:40;"`
}

// 读取数据源写入excel
func TestExport(t *testing.T) {
	// 下面这个 employeeList 可以通过mysql或者其他数据源获取,这里因为测试所以就写固定的
	var employeeList = []employee{
		{"fuhua", "符华", "fuhua@123.com", "太虚剑派", "开发部", "CJGLY", "1", "备注备注"},
		{"baiye", "白夜", "baiye@123.com", "天命科技有限公司", "执行部", "PTYG", "2", ""},
		{"chiling", "炽翎", "chiling@123.com", "太虚剑派", "行政部", "PTYG", "2", "备注备注备注备注"},
		{"yunmo", "云墨", "yunmo@123.com", "太虚剑派", "财务部", "CJGLY", "1", ""},
		{"yuelun", "月轮", "yuelun@123.com", "天命科技有限公司", "执行部", "CJGLY", "1", ""},
		{"diaomao", "吊毛",
			"diaomao@qq.com哈哈哈哈哈哈哈哈这里是最大行高测试哈哈哈哈哈哈哈哈这11111111111里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测试",
			"吊毛科技有限公司", "开发部", "PTYG", "2",
			"备注备注备注备注com哈哈哈哈哈哈哈哈这里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测里是最大行高测试哈哈哈哈哈哈哈哈这里是最大行高测试"},
	}
	// changeHead := map[string]string{}
	changeHead := map[string]string{"Id": "账号", "Name": "真实姓名"}
	//f, err := excel.NormalExport(employeeList, "Sheet1", "用户信息", "Id,Email,", true, true, changeHead)
	f, err := excel.NormalDynamicExport(employeeList, "Sheet1", "员工信息", "", true, false, changeHead)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Path = "./员工信息.xlsx"
	if err := f.Save(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(err)
}

// 测试读取excel
func TestReadExcel(t *testing.T) {
	f, err := excelize.OpenFile("./员工信息.xlsx")
	if err != nil {
		fmt.Println("文件打开失败")
	}
	importList := []employee{}
	err = excel.ImportExcel(f, &importList, 1, 2)
	if err != nil {
		fmt.Println(err)
	}
	marshal, _ := json.Marshal(importList)
	fmt.Println(string(marshal))
	fmt.Println("--------------------------------")
	for _, t := range importList {
		fmt.Printf("%#v\n", t)
	}
}
