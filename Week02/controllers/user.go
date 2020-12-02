package controllers

import (
	"fmt"
	"github.com/unknwon/com"
	"homeworkweek2/dao"
	"net/http"
)

func GetUser(r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("解析表单数据失败!")
	}
	id := com.StrTo(com.ToStr(r.Form["uid"])).MustInt()
	user, err := dao.GetUserById(id)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
	if user.Phone == "" {
		//接口返回没有查找到数据就好
	}
	//其余的流程
}
