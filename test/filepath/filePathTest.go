package main

import (
	//"path/filepath"
	//"strings"
	"strconv"
	"fmt"

	//"bytes"
	"os"
	"encoding/csv"
	//"path/filepath"
	//"strings"
	"path/filepath"
)
func main() {
	//files,_:=filepath.Glob("D:\\Repository\\src\\gitlab.33.cn\\lihailei\\chain33_tools\\gmsg-framework\\*.csv")
	//for _,file :=range files{
	//	fmt.Println(file)
     //   fmt.Println("ext:",filepath.Ext(file))
     //   fmt.Println("base:",filepath.Base(file))
     //   fmt.Println("dir:",filepath.Dir(file))
     //   fmt.Println("join:",filepath.Join(filepath.Dir(file),"csv"))
     //   fileName := filepath.Base(file)
     //   strings.Split(fileName,".csv")
     //   strs :=strings.Split(strings.Split(fileName,".csv")[0],"_")
     //   fmt.Println("strs:",strs)
     //   num,_:=strconv.ParseInt(strs[len(strs)-1],10,10)
     //   fmt.Println("num:",num)
	//
	//}
	taskDir :=filepath.Join(filepath.Dir("D:\\Repository\\src\\gitlab.33.cn\\lihailei\\chain33_tools\\gmsg-framework\\"),"datadir")
	err:=os.RemoveAll(taskDir)
	if err !=nil {
		fmt.Println(err.Error())
	}
	fmt.Println("taskDir:",taskDir)
	//fmt.Println("name:",cacelStoreFileName("xxxxxxxxx",1))
	index := &Index{
		10,
		10000,
		20000,
	}
	updateIndex("D:\\Repository\\src\\gitlab.33.cn\\lihailei\\chain33_tools\\gmsg-framework\\index",index)
}
func cacelStoreFileName(taskId string,index int)string{
	return fmt.Sprintf("%s_%08d.csv",taskId,index)
}
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func updateIndex(filePath string, index *Index) error {
	//if pathExists(filePath) {
	//	err := os.Remove(filePath)
	//	if err != nil {
	//		return err
	//	}
	//}
	//csvFile, err := os.Create(filePath)
	fmt.Println("filepath:",filePath)
	csvFile, err := os.OpenFile(filePath,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	line := []string{strconv.FormatInt(index.LastSliceNum, 10), strconv.FormatInt(index.CurrentCount, 10), strconv.FormatInt(index.PerSliceNum, 10)}
	fmt.Println(line)
	err = writer.Write(line)
	if err != nil {
		fmt.Println("err:",err.Error())
		return err
	}
	fmt.Println("11111111111")
	writer.Flush()
	return nil
}
type Index struct {
	LastSliceNum int64 `json:"last_slice_num"`   //最近第几个分片
	CurrentCount int64 `json:"current_count"`   //当前分片记录总数
	PerSliceNum  int64  `json:"per_slice_num"`  //每个分片，里面有多少条记录
}