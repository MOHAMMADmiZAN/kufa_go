package KufaSessions

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// random String //
//const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//
//func RandStringBytes(n int) string {
//	b := make([]byte, n)
//	for i := range b {
//		b[i] = letterBytes[rand.Intn(len(letterBytes))]
//	}
//	return string(b)
//}

var (
	Key   = []byte(securecookie.GenerateRandomKey(32))
	Store = sessions.NewCookieStore(Key)
)
