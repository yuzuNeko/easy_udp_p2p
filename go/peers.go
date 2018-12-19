package main 

import (

	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)
var tag string
const Hand_shake_msg="打洞消息"
func main(){
	//命令行给进程传递一个名字
	tag=os.Args[1]
	port,_:=strconv.Atoi( os.Args[2])
	srcAddr:=&net.UDPAddr{IP:net.IPv4zero,Port:port}
	dstAddr:=&net.UDPAddr{IP:net.ParseIP("your server"),Port:9100}
	conn,err:=net.DialUDP("udp",srcAddr,dstAddr)
	if err!=nil{

	}
	if _,err:=conn.Write([]byte("hello,peer"+tag));err!=nil{
		log.Panic(err)
	}
	data:=make([]byte,1024)
	fmt.Println("开始卡住")
	n,remoteAddr,err:=conn.ReadFromUDP(data)
	if err!=nil{

	}
	fmt.Println("结束卡住")
	conn.Close()
	anotherPeer:=parseAddr(string(data[:n]))
	fmt.Println("开始打洞")
	fmt.Printf("local:%s server:%s another:%s\n", srcAddr, remoteAddr, anotherPeer.String())
	bidirectionHole(srcAddr,&anotherPeer)

}

func parseAddr(addr string)net.UDPAddr{
	t:=strings.Split(addr,":")
	port,_:=strconv.Atoi(t[1])
	return net.UDPAddr{
		IP:net.ParseIP(t[0]),
		Port:port,
	}
}

func bidirectionHole(srcAddr *net.UDPAddr,anotherPeer *net.UDPAddr){
	conn,err:=net.DialUDP("udp",srcAddr,anotherPeer)
	if err!=nil{

	}
	defer conn.Close()
	if _,err=conn.Write([]byte(Hand_shake_msg));err!=nil{

	}
	go func(){
		for{
			time.Sleep(10*time.Second)
			if _,err=conn.Write([]byte("from ["+tag+"]"));err!=nil{

			}
		}
	}()
	for {
		data:=make([]byte,1024)
		n,_,err:=conn.ReadFromUDP(data)
		if err!=nil{

			}else{
				fmt.Println("收到数据",string(data[:n]))
			}
	}

}

