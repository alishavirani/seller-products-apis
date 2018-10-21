package apis

import (
	"net/http"

	"github.com/globalsign/mgo"
)

// var sess *mgo.Session

func InitSession(session *mgo.Session, r *http.Request) {
	// fmt.Println("session in init session", session)

	// context.Set(r, "session", session)

	// utils.SetMyKey(r, session)
}
