package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"seven_work/proto"
)

type User struct {
	name     string
	password string
}

func main() {
	coon, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer coon.Close()

	client := proto.NewRegisterClient(coon)
	do(client)

}
func do(client proto.RegisterClient) {
	var u User

	for {
		{
			fmt.Println("=============================")
			fmt.Println("1)注册  \t2)登录")
			fmt.Println("3)修改信息\t0)exit")
			fmt.Println("=============================")
		}
		var c int
		_, err := fmt.Scanln(&c)
		if err != nil {
			break
		}
		switch c {
		case 1:
			{
				fmt.Println("请输入你的姓名和密码")
				fmt.Scanln(&u.name, &u.password)
				registe(client, u.name, u.password)
			}
		case 2:
			{
				fmt.Println("请输入你的姓名和密码")
				fmt.Scanln(&u.name, &u.password)
				verifyUser(client, u.name, u.password)
			}
		case 3:
			{
				if len(u.name) == 0 || len(u.password) == 0 {
					fmt.Println("请先登录")
					continue
				}
				fmt.Println("请输入你想修改的姓名或密码（1为姓名，2为密码）")
				var code int32
				var str string
				fmt.Scanln(&code, &str)
				updateUser(client, u.name, u.password, code, str)
			}
		default:
			return

		}
	}
}

//注册
func registe(c proto.RegisterClient, n string, p string) {
	resp, err := c.Reg(context.Background())
	if err != nil {
		log.Panic(err)
	}
	err = resp.Send(&proto.Users{Name: n, Passwd: p})
	if err != nil {
		log.Println(err)
	}
	//读取
	go func() {
		recv, err := resp.Recv()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(recv.Status)
	}()

}

//验证
func verifyUser(c proto.RegisterClient, n string, p string) {
	resp, err := c.LoginName(context.Background())
	if err != nil {
		log.Panic(err)
	}
	err = resp.Send(&proto.Users{Name: n, Passwd: p})
	if err != nil {
		log.Println(err)
	}
	//读取
	go func() {
		recv, err := resp.Recv()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(recv.Status)
	}()
}

//修改
func updateUser(c proto.RegisterClient, n string, p string, code int32, newName string) {
	resp, err := c.Update(context.Background())
	if err != nil {
		log.Panic(err)
	}
	err = resp.Send(&proto.UpdateInfor{V: &proto.Users{Name: n, Passwd: p}, Code: code, Str: newName})
	if err != nil {
		log.Println(err)
	}

	go func() {
		recv, err := resp.Recv()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(recv.Status)
	}()

}
