package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
	"notes/logging"
	"notes/model"
	"notes/storage"
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

func GetAll(ctx context.Context, _ *model.EmptyRequest) ([]model.NewNote, error) {

	res, err := storage.GetAllNotes(ctx)
	if err != nil {
		logging.Logger.Printf("[GetAll] get all notes failed | err:%v", err)
		return nil, err
	}

	return res, nil
}

func GetOne(ctx context.Context, note *model.Note) (model.NewNote, error) {
	return storage.GetNoteByID(ctx, note.ID)
}

func Add(ctx context.Context, newNote *model.Note) (model.EmptyResponse, error) {

	var err error

	content := newNote.Content

	if err = storage.AddNote(ctx, content); err != nil {
		logging.Logger.Printf("[Add] add note failed | err:%v | content:%v", err, content)
		return model.EmptyResponse{}, err
	}

	return model.EmptyResponse{}, err
}

func Delete(ctx context.Context, note *model.Note) (model.EmptyResponse, error) {

	var err error

	err = storage.RemoveNote(ctx, note.ID)
	return model.EmptyResponse{}, err
}

// 生成一个随机id
func genID(content string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10000)
	return fmt.Sprintf("%x", md5.Sum([]byte(content+string(i))))
}
