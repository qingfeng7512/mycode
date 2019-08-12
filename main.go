package test

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Projects struct {
	name,
	fileList,
	urlList []string
}

//项目父级目录
var project string

var proj Projects

var f = excelize.NewFile()

const splitBackslash = "\\"

func main() {
	println("正在启动....")
	println("****************************************操作说明********************************")
	println("*  上git-->show log-->Copy clipboard-->full paths  粘贴到 text.txt文件中即可   *")
	println("********************************************************************************")

	println("请输入项目父级目录 例: workspace/winsea-saas-common-base  输入 workspace")
	fmt.Scanln(&project)

	var funcationName string
	println("请输入此次功能名称......")
	fmt.Scan(&funcationName)

	var gitName string
	println("请输入git账户名称...")
	fmt.Scan(&gitName)

	var name string
	println("请输入你的名字...")
	fmt.Scan(&name)

	//读取文件
	data := readFile()
	proj = rinseData(project+"\\", data)

	ExcelInit()
	f.SetCellValue("Sheet1", "D1", proj.name[0])
	f.SetCellValue("Sheet1", "D2", funcationName)

	nowDate := time.Now()
	x := time.Date(nowDate.Year(), nowDate.Month(), nowDate.Day(), nowDate.Hour(), nowDate.Minute(), nowDate.Second(), 20, time.Local)

	for index := range proj.urlList {
		f.MergeCell("Sheet1", "A"+strconv.Itoa(index+4), "E"+strconv.Itoa(index+4))
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(index+4), proj.urlList[index])

		f.MergeCell("Sheet1", "F"+strconv.Itoa(index+4), "J"+strconv.Itoa(index+4))
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(index+4), proj.fileList[index])

		f.SetCellValue("Sheet1", "K"+strconv.Itoa(index+4), gitName)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(index+4), x.Format("2006-01-02 15:04:05"))
	}

	var porjectN = "代码提交清单" + "-" + name + "-" + strconv.Itoa(nowDate.Year()) + strconv.Itoa(int(nowDate.Month())) + strconv.Itoa(nowDate.Day())

	var err = f.SaveAs("./" + porjectN + "_1" + ".xlsx")

	if err != nil {
		fmt.Println(err)
	}

}

//读取文件
func readFile() string {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(data)
}

//清洗数据
func rinseData(projectUrl string, data string) Projects {

	var projectList []string
	var fileList []string
	var urlList []string

	oldData := strings.Split(data, projectUrl)
	deleteString := oldData[0]

	for index := range oldData {
		oldData[index] = strings.Trim(oldData[index], deleteString)
	}

	//删选过后的集合
	var newData = noNil(oldData)

	for index := range newData {
		var debris = strings.Split(newData[index], splitBackslash)
		projectList = append(projectList, debris[0])
		//已截取 项目名的路径
		z := debris[1 : len(debris)-1]
		urlList = append(urlList, strings.Join(z, splitBackslash))
		fileList = append(fileList, debris[len(debris)-1])
	}

	projectList = checkRepetition(projectList)

	projects := Projects{name: projectList, urlList: urlList, fileList: fileList}

	fmt.Println("项目名称:", projectList)
	fmt.Println("文件路径:", urlList)
	fmt.Println("文件名称:", fileList)
	return projects
}

func ExcelInit() {
	f.MergeCell("Sheet1", "A1", "C1")
	f.MergeCell("Sheet1", "D1", "L1")

	f.MergeCell("Sheet1", "A2", "C2")
	f.MergeCell("Sheet1", "D2", "L2")

	f.MergeCell("Sheet1", "A3", "E3")
	f.MergeCell("Sheet1", "F3", "J3")

	f.SetCellValue("Sheet1", "A1", "功能名")
	f.SetCellValue("Sheet1", "B2", "工程名")

	f.SetCellValue("Sheet1", "E3", "报名/目录名")
	f.SetCellValue("Sheet1", "F3", "文件名")
	f.SetCellValue("Sheet1", "K3", "GIT账号")
	f.SetCellValue("Sheet1", "L3", "提交时间")
	style, error := f.NewStyle(`{
		"fill":{
			"type":"pattern",
			"color":["#445E6A"],
			"pattern":1
		}
	}`)
	borderStyle, borderError := f.NewStyle(`{
	"border":[
		{"type":"top","color":"#000000","style":1},
		{"type":"left","color":"#000000","style":1},
		{"type":"right","color":"#000000","style":1},
		{"type":"bottom","color":"#000000","style":1}
	]}`)

	f.SetCellStyle("Sheet1", "A1", "A4", style)
	f.SetCellStyle("Sheet1", "A1", "L"+strconv.Itoa(len(proj.fileList)), borderStyle)

	if error != nil || borderError != nil {
		fmt.Println("borderError====================>", borderError)
		fmt.Println("error=========================>", error)
	}
}

//集合 非空判断
func noNil(data []string) []string {
	var newData []string
	for index := range data {
		if data[index] != "" {
			newData = append(newData, data[index])
		}
	}
	return newData
}


