package service

import (
	"badminton-go/db/dao"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func MatchCreateHandler(w http.ResponseWriter, r *http.Request) {
	name, startTime, err := getCreateParam(r)
	if err != nil {
		responseErr(w, err)
		return
	}
	matchId, err := dao.MatchDaoIns.Create(0, name, startTime)
	if err != nil {
		responseErr(w, err)
	} else {
		res := &JsonResult{}
		res.Data = matchId
		response(w, res)
	}
}

func responseErr(w http.ResponseWriter, err error) {
	res := &JsonResult{}
	res.Code = -1
	res.ErrorMsg = err.Error()
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

func getCreateParam(r *http.Request) (string, time.Time, error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return "", time.Now(), err
	}
	defer r.Body.Close()
	name := ""
	startTime := time.Now()
	if v, ok := body["name"]; !ok || v == "" {
		return "", time.Now(), fmt.Errorf("name err")
	} else {
		name = v.(string)
	}
	if v, ok := body["startTime"]; !ok {
		return "", time.Now(), fmt.Errorf("startTime err")
	} else {
		layout := "2006-01-02 15:04:05"
		startTime, _ = time.Parse(layout, v.(string))
	}
	return name, startTime, nil
}
