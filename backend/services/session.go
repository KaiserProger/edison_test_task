package services

import (
	"github.com/gorilla/sessions"
)

func GetStoreOrCreate[T any](session *sessions.Session, key string) []*T {
	empty_array := make([]*T, 0)
	raw, found := session.Values[key]
	if !found {
		return empty_array
	}
	array, ok := raw.([]*T)
	if !ok {
		return empty_array
	}
	return array
}
