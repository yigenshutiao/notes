package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
	"notes/model"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Hello(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := wr.Write([]byte(`
  __  _  _ ____ ____  __  _  _ ____    __ _  __ ____ ____ ____     __  ____ ____ 
 / _\/ )( (  __) ___)/  \( \/ |  __)  (  ( \/  (_  _|  __) ___)   / _\(  _ (  _ \
/    \ /\ /) _)\___ (  O ) \/ \) _)   /    (  O ))(  ) _)\___ \  /    \) __/) __/
\_/\_(_/\_|____|____/\__/\_)(_(____)  \_)__)\__/(__)(____|____/  \_/\_(__) (__)

`))
	if err != nil {
		return
	}
}

var notes = map[string]model.Note{}

func GetAll(ctx context.Context, _ *model.EmptyRequest) (map[string]model.Note, error) {

	return notes, nil
}

func GetOne(ctx context.Context, note *model.Note) (model.Note, error) {

	var err error

	id := note.ID
	if _, exist := notes[id]; !exist {
		return model.Note{}, err
	}

	return notes[id], nil
}

func Add(ctx context.Context, newNote *model.Note) (model.EmptyResponse, error) {

	var err error

	content := newNote.Content

	id := genID(content)

	note := model.Note{
		ID:         id,
		Content:    content,
		StartTime:  time.Now(),
		UpdateTime: time.Now(),
	}
	notes[id] = note

	return model.EmptyResponse{}, err

}

func Delete(ctx context.Context, note *model.Note) (model.EmptyResponse, error) {

	var err error

	id := note.ID
	if _, exist := notes[id]; !exist {
		return model.EmptyResponse{}, err
	} else {
		delete(notes, id)
		return model.EmptyResponse{}, nil
	}
}

// 生成一个随机id
func genID(content string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10000)
	return fmt.Sprintf("%x", md5.Sum([]byte(content+string(i))))
}
