package service

import (
	"badminton-go/db/dao"
	"encoding/json"
	"fmt"
	"net/http"
)

const DefPageSize int = 10

// CounterHandler 计数器接口
func MatchListHandler(w http.ResponseWriter, r *http.Request) {
	page, size := getListParam(r)
	res := &JsonResult{}
	list, err := dao.MatchDaoIns.List(page, size)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
	} else {
		res.Data = list
	}
	response(w, res)
}

func response(w http.ResponseWriter, res *JsonResult) {
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func getListParam(r *http.Request) (int, int) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return 1, DefPageSize
	}
	defer r.Body.Close()
	page := 1
	size := DefPageSize
	if v, ok := body["page"]; ok {
		page = v.(int)
	}
	if v, ok := body["size"]; ok {
		size = v.(int)
	}
	return page, size
}
