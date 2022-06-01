package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-01 16:33
 **/
/*
func main(){
	start:=time.Now()
	var wg sync.WaitGroup
	for i:=1;i<65535;i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			address:= fmt.Sprintf("127.0.0.1:%d", i)
			conn,err:=net.Dial("tcp",address)
			if err!=nil{
				fmt.Printf("%s 关闭了\n",address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了\n",address)
		}(i)
	}
	wg.Wait()
	elapsed:=time.Since(start)/1e9
	fmt.Printf("%d seconds",elapsed)
}

func main(){
	for i:=21;i<120;i++{
		address:= fmt.Sprintf("20.194.168.28:%d", i)
		conn,err:=net.Dial("tcp",address)
		if err!=nil{
			fmt.Printf("%s 关闭了\n",address)
			continue
		}
		conn.Close()
		fmt.Printf("%s 打开了\n",address)
	}
}
*/

func worker(ports chan int,res chan int){
	for p:=range ports{
		address:=fmt.Sprintf("127.0.0.1:%d",p)
		conn,err:=net.Dial("tcp",address)
		if err!=nil{
			fmt.Printf("%d errors!!!\n",p)
			res <- 0
			continue
		}
		conn.Close()
		fmt.Printf("%d 打开了\n",p)
		res<-p

	}
}

func main(){
	start := time.Now()
	ports:=make(chan int,100)
	res:=make(chan int)
	var openPorts []int
	var closedPorts []int


	for i:=0;i<cap(ports);i++{
		go worker(ports,res)
	}

	go func(){
		for i:=1;i<1024;i++{
			ports<-i
		}
	}()


	for i:=1;i<1024;i++{
		port:=<-res
		if port!=0{
			openPorts=append(openPorts,port)
		}else{
			closedPorts=append(closedPorts,port)
		}
	}

	close(ports)
	close(res)


	sort.Ints(openPorts)
	sort.Ints(closedPorts)
	for _,port:=range openPorts{
		fmt.Printf("%d 端口打开了\n",port)
	}

	elapsed := time.Since(start)
	fmt.Println("运行时间：", elapsed)
}