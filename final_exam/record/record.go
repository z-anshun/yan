package record

import (
	"os"
)

func Record(filename string,str []byte)error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0600)
	if err!=nil{
		return nil
	}

	_, err = f.Write(str)
	if err!=nil{
		return err
	}

	return nil
}
