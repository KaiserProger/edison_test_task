package fs

import (
	"os"

	"github.com/gorilla/sessions"
)

func NewFSStore() sessions.Store {
	store := sessions.NewFilesystemStore("./fs", []byte(os.Getenv("SESSION_KEY")))
	store.MaxLength(0)
	return store
}
