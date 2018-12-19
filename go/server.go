package main 

import (

	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	listener,err:=net.ListenUDP("udp",&net.UDPAddr{	IP:net.IPv4zero,Port:5555})
	if err!=nil{

	}
	log.Printf("local address %s",listener.LocalAddr().String())
	peers:=make([]net.UDPAddr,0,2)
	data:=make([]byte,1024)
	fmt.Println("开始进入udp服务器循环")
	for{
		n,remoteAddr,err:=listener.ReadFromUDP(data)
		if err!=nil{

		}
		fmt.Println(n)
		log.Printf("%s %s",remoteAddr.String(),data[:n])
		peers=append(peers,*remoteAddr)
		fmt.Println(peers)
		if len(peers)==2{
			log.Printf("进行udp打洞,建立 %s <------> %s",peers[0].String(),peers[1].String())
			listener.WriteToUDP([]byte(peers[1].String()),&peers[0])
			listener.WriteToUDP([]byte(peers[0].String()),&peers[1])
			time.Sleep(time.Duration(3)*time.Second)
			fmt.Println("中转服务器退出")
			return 

		}
	}

}