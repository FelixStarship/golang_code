package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

//写文件
func (m *Monster) Store() error {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal err=", err)
		return err
	}
	filePath := "d:/monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err=", err)
		return err
	}
	return nil
}

//读文件
func (m *Monster) ReStore() error {
	filePath := "d:/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err=", err)
		return err
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("Unmarshal err=", err)
		return err
	}
	return nil
}
