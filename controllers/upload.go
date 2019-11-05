package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/Luxurioust/excelize"
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) UpForm() {
	this.TplName = "upload/index.html"
}

func (this *UploadController) UpFile() {

	f, h, _ := this.GetFile("myfile") //获取上传的文件
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".xlsx": true,
		".xls":  true,
		".csv":  true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		this.Ctx.WriteString("后缀名不符合上传要求")
		return
	}
	//创建目录
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir, 777)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("%v", err))
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))

	fileName := fmt.Sprintf("%x", hashName) + ext
	//this.Ctx.WriteString(  fileName )

	fpath := uploadDir + fileName
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = this.SaveToFile("myfile", fpath)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("%v", err))
	}
	xlsx, err := excelize.OpenFile(fpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get value from cell by given sheet index and axis.
	//cell, _ := xlsx.GetCellValue("Sheet1", "B2")
	// Get sheet index.
	index := xlsx.GetSheetIndex("Sheet1")
	// Get all the rows in a sheet.
	rows, _ := xlsx.GetRows("Sheet" + strconv.Itoa(index))
	for k, row := range rows {
		if k != 0 {
			userName := ""
			a := 0.0
			b := 0.0
			c := 0.0
			d := 0.0
			for kCell, colCell := range row {
				switch kCell {
				case 0:
					userName = colCell
				case 1:
					a, _ = strconv.ParseFloat(colCell, 64)
				case 2:
					b, _ = strconv.ParseFloat(colCell, 64)
				case 3:
					c, _ = strconv.ParseFloat(colCell, 64)
				case 4:
					d, _ = strconv.ParseFloat(colCell, 64)
				}
			}
			err = scoreModel.AddRealScore(userName, a, b, c, d)
			if err != nil {
				this.Ctx.WriteString("Excel文件数据存在问题，请联系开发人员核对数据！")
			}
		}
	}
	this.Ctx.WriteString("上传成功！数据处理完成！")
}
