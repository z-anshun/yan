package proto

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"seven_work/db"
	"seven_work/limite"
)

type UserService struct {
	UnimplementedRegisterServer
}

func (t *UserService) LoginName(coon Register_LoginNameServer) error {
	lim := limite.NewConnLimiter()
	for {
		user, err := coon.Recv()
		//先取，再往chan里面塞
		lim.GetConn()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		//开携程进行验证
		go func() {
			msg := db.Verify(user.Name, user.Passwd)
			s, err := json.Marshal(&msg)
			if err != nil {
				log.Panicln(err)
			}
			if err := coon.Send(&ResponseMsg{Status: string(s)}); err != nil {
				log.Println(err)
			}
			lim.ReleaseConn()
		}()

	}

}
func (t *UserService) Reg(coon Register_RegServer) error {
	lim := limite.NewConnLimiter()
	for {
		user, err := coon.Recv()
		//先取，再往chan里面塞
		lim.GetConn()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		//开携程进行验证
		go func() {
			msg := db.AddUser(user.Name, user.Passwd)
			s, err := json.Marshal(&msg)
			if err != nil {
				log.Panicln(err)
			}
			if err := coon.Send(&ResponseMsg{Status: string(s)}); err != nil {
				log.Println(err)
			}
			lim.ReleaseConn()
		}()

	}
}
func (t *UserService) Update(coon Register_UpdateServer) error {
	lim := limite.NewConnLimiter()
	for {
		in, err := coon.Recv()
		lim.GetConn()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		go func() {
			msg := db.Verify(in.V.Name, in.V.Passwd)
			fmt.Println(msg,in.V)
			if msg.Code != 204 {
				lim.ReleaseConn()
				coon.Send(&ResponseMsg{Status: "name or psssword error"})
				return
			}
			str := db.UpdateInfromation(in.Code, in.V.Name, in.Str).Message

			coon.Send(&ResponseMsg{Status: str})
			lim.ReleaseConn()
		}()
	}

}
