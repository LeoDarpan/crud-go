package sessions

import (
	"github.com/gorilla/sessions"
)

//Declaring a sessions object
var Store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))
