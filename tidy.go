package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main(){
	str := []string{
	"0厨余垃圾_饼干",
	"1厨余垃圾_蛋壳",
	"2厨余垃圾_香蕉皮",
	"3厨余垃圾_落叶",
	"4厨余垃圾_unknown",
	"5可回收物_易拉罐",
	"6可回收物_钥匙",
	"7可回收物_钉子",
	"8可回收物_纸牌",
    "9可回收物_unknown",
    "10其他垃圾_烟",
    "11其他垃圾_牙线",
    "12其他垃圾_透明胶",
    "13其他垃圾_橡皮擦",
    "14其他垃圾_unknown",
    "15有害垃圾_电池",
    "16有害垃圾_维生素C含片",
    "17有害垃圾_水笔芯",
    "18有害垃圾_指甲油",
    "19有害垃圾_unknown",
    }
	lens := len(str)
	imgpt := 0
	var newpath, newpathtxt string
	for i := 0; i < lens; i++{
		err := filepath.Walk("images/" + str[i], func(path string, info os.FileInfo, err error) error{
			if err != nil {
				return err
			}
			if info.IsDir(){
				return nil
			}
			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			newpath = "imgsc/" + "imgsc_" + strconv.Itoa(imgpt) + filepath.Ext(path)
			newpathtxt = "imgsc/" + "imgsc_" + strconv.Itoa(imgpt) + ".txt"
			ioutil.WriteFile(newpath, buf, 0755)
			if err != nil {
				return err
			}
			txtfile, err := os.Create(newpathtxt)
			if err != nil {
				return err
			}
			txtfile.Write([]byte(filepath.Base(newpath) + ", " + strconv.Itoa(i)))
			imgpt++
			return nil
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
