package logic

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"notes/model"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Hello(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := wr.Write([]byte("This is a awesome note app!"))
	if err != nil {
		return
	}
}

var notes = map[string]model.Note{}

func GetAll(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	data := map[string]interface{}{
		"msg":  "success",
		"data": notes,
	}

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	_, err = wr.Write(res)
	if err != nil {
		return
	}

}

func GetOne(wr http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if _, exist := notes[id]; !exist {
		return
	}

	data := map[string]interface{}{
		"msg":  "success",
		"data": notes[id],
	}

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	_, err = wr.Write(res)
	if err != nil {
		return
	}
}

func Add(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	content := r.URL.Query().Get("content")
	if content == "" {
		return
	}

	id := genID(content)
	note := model.Note{
		ID:         id,
		Content:    content,
		StartTime:  time.Now(),
		UpdateTime: time.Now(),
	}
	notes[id] = note

	data := map[string]interface{}{
		"msg":  "success",
		"data": "",
	}

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	_, err = wr.Write(res)
	if err != nil {
		return
	}
}

func Delete(wr http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := r.URL.Query().Get("id")
	if _, exist := notes[id]; exist {
		delete(notes, id)
	}

	data := map[string]interface{}{
		"msg":  "success",
		"data": "",
	}

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	_, err = wr.Write(res)
	if err != nil {
		return
	}

}

// 生成一个随机id
func genID(content string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10000)
	return fmt.Sprintf("%x", md5.Sum([]byte(content+string(i))))
}
