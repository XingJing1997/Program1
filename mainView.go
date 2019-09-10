package main

import (
	"190831bookManage2.0/service"
	"fmt"
	"os"
)
var(
	id,title,author string
	price float32
	publish bool
)
func AddInfo(){
	fmt.Println("please input the infos of Book that needs to add:")
	//输入要添加的书籍信息
	fmt.Print("id:")
	fmt.Scanln(&id)
	fmt.Print("title:")
	fmt.Scanln(&title)
	fmt.Print("author:")
	fmt.Scanln(&author)
	fmt.Print("price:")
	fmt.Scanln(&price)
	fmt.Print("publish:")
	_,_=fmt.Scanln(&publish)
	service.Add(id,title,author,price,publish)
}
func ChangeInfo(){
	//var Id string
	service.Show()
	fmt.Print("please input the book id that needs to change:")
	fmt.Scanln(&id)
	fmt.Println("please input changed infos:")
	flag:=service.Change(id)
	if(flag==0){
		return
	}
}
func ShowInfo(){
	service.Show()
}
func DeleteInfo(){
	var Id string
	service.Show()
	fmt.Println("请输入你要删除的书籍编号ID：")
	fmt.Scanln(&Id)
	flag:=service.Delete(Id)
	if(flag==0){
		return
	}
}


//用户主视图
func mainView(){
	for {
		var choice int
		fmt.Println("---------------------------------")
		fmt.Println("|           1.添加书籍           |")
		fmt.Println("| 			2.修改书籍           |")
		fmt.Println("| 			3.展示所有书籍        |")
		fmt.Println("| 			4.删除所有书籍        |")
		fmt.Println("| 			5.退出程序		   	|")
		fmt.Println("---------------------------------")
		fmt.Print("please make a choice:")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			AddInfo()
		case 2:
			ChangeInfo()
		case 3:
			ShowInfo()
		case 4:
			DeleteInfo()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("invalidate operate,please make a new choice:")
		}
	}
}
func main(){
	//调用用户主视图
	mainView()
}
