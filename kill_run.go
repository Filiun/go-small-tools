
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {


	//const jsonStream = `
    //  { "Name" : "Ed" , "Text" : "Knock knock." }
    //  { "Name" : "Sam" , "Text" : "Who's there?" }
    //  { "Name" : "Ed" , "Text" : "Go fmt." }
    //  { "Name" : "Sam" , "Text" : "Go fmt who?" }
    //  { "Name" : "Ed" , "Text" : "Go fmt yourself!" }
    //`
	//
	//reader:=strings.NewReader(jsonStream)
	//dec:=json.NewDecoder(reader)
	//type Message struct {
	//	Name , Text string
	//}
	//var m  Message
	//dec.Decode(&m)
	//fmt.Println(m.Name,m.Text)
	//fmt.Println(time.Now().Unix())


	var   name,program_name,run_name  string
	flag.StringVar(&name,"k"," ","the name you will kill process  no usage")   //要关闭的进程 关键字
	flag.StringVar(&run_name,"r"," ","the name you will  run process  no usage") //要运行的脚本
	flag.Parse()

	program_name=strings.Replace(os.Args[0],"./","",-1)
	fmt.Println(program_name)

	cmd,_ := exec.Command("bash", "-c", " ps  aux |  grep  "+name+"  | grep  -v  grep | grep  -v  "+program_name+"  | awk '{print  $2}'   ").Output()
	cmd1:=exec.Command("bash","-c"," kill "+ string(cmd)+" ")
	cmd1.Start()
	res:=exist(run_name)

	if res {
		if run_name != name  {
			err:=os.Rename(run_name,name)
			if err !=nil {
				fmt.Println("文件重命名失败："+err.Error())
			}
		}

		fileinfo,_:=os.Stat(name)
		fileMode:=fileinfo.Mode()
		fmt.Println(fileMode.String())
		perm:=fileMode.Perm()
		fmt.Println(perm)
		flag1:=perm & os.FileMode(73)

		if uint32(flag1) == uint32(73) {

			cmd:=exec.Command("bash","-c"," nohup  ./"+name+"   > nohup.out  & ")
			err1:=cmd.Start()
			fmt.Println(err1)
			fmt.Println("程序启动完成！")
		}else{

			file,_:=os.Open(name)
			err:=file.Chmod(0777)
			if err !=nil  {
				fmt.Println("文件权限修改失败"+err.Error())
				os.Exit(1)
			}
			cmd:=exec.Command("bash","-c"," nohup  ./"+name+"  > nohup.out   & ")
			err1:=cmd.Start()
			fmt.Println(err1)
			fmt.Println("程序启动完成！")

		}
	}



// aes 加密解密

	//key:=[]byte("ABCDEFGHIJKLMNOP")
	//block,_:=aes.NewCipher(key)
	//size:=block.BlockSize()
	//data:=[]byte("tianxiahuizainali")
	//encry:=make([]byte,len(data))
	//block.Encrypt(encry,data)
	//fmt.Println(encry)
	//fmt.Println(size)
	//fmt.Println(string(encry))
	//dest:=make([]byte,len(encry))
	//block.Decrypt(dest,encry)
	//fmt.Println(string(dest))


}

//遍历目录
func scanDirFile(dirname string){
	
	fileinfo,err:=ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
	}
	for _,v:= range fileinfo{
		if v.IsDir() {
			dir:=dirname+"/"+v.Name()
			scanDirFile(dir)
		}else{
			fmt.Println(v.Name())
		}

	}
}

func tracefile(str_content string)  {
	fd,_:=os.OpenFile("./log.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_time:=time.Now().Format("2006-01-02 15:04:05")
	fd_content:=strings.Join([]string{"======",fd_time,"=====",str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}


func  getMonthDay(num int)string{

	loc, _ := time.LoadLocation("Asia/Shanghai")
	time1 := time.Unix(time.Now().Unix(), 0)
	timeFormat := "2006-01"
	str := time1.In(loc).Format(timeFormat)
	firstDay:=str+"-01"
	tt,_:=time.ParseInLocation("2006-01-02",firstDay,loc)
	numDay:=tt.AddDate(0,0,num)
	ss:=numDay.Format("2006-01-02")
	return ss
}

func getMonthDays() int {

	loc, _ := time.LoadLocation("Asia/Shanghai")
	time1 := time.Unix(time.Now().Unix(), 0)
	timeFormat := "2006-01"
	str := time1.In(loc).Format(timeFormat)
	firstDay:=str+"-01"
	tt,_:=time.ParseInLocation("2006-01-02",firstDay,loc)
	numDay:=tt.AddDate(0,1,-1)
	return numDay.Day()
}

func parseXml()  {
	type Address struct{
		City string
		Area string
	}

	type Email struct{
		Where string `xml:"where,attr"`
		Addr string
	}

	type Student struct{
		Id int `xml:"id,attr"`
		Address
		Email []Email
		FirstName string `xml:"name>first"`
		LastName string `xml:"name>last"`
	}

	//实例化对象
	stu := Student{23, Address{"shanghai","pudong"},[]Email{Email{"home","home@qq.com"}, Email{"work","work@qq.com"}},"chain","zhang"}
	fmt.Println("stu:", stu)
	//序列化
	buf,err := xml.Marshal(stu)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("xml: ", string(buf))
	var newStu Student
	//反序列化
	err1 := xml.Unmarshal(buf, &newStu)
	if err1 != nil{
		fmt.Println(err1.Error())
		return
	}
	fmt.Println("newStu: ", newStu)
}


func  exist (file string)bool{
	_,err:=os.Stat(file)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
