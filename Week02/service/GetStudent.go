package service

import (
	"demo/dao"
	"fmt"
)

func GetStudent() {
	defer dao.Db.Close()
	student, err := dao.GetBasicStudentByNo(123456)
	if err != nil {
		fmt.Printf("query student err : %+v", err)
		return
	}
	fmt.Println("query student : ", student)
}
