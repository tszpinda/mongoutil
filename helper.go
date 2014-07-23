package mgou

import (
	"labix.org/v2/mgo"
)

//TestCtx is a mongo test context
type TestCtx struct {
	Db *mgo.Database
	S  *mgo.Session
}

//WithTestCtx simple function wrapper providing mongodb ctx
func WithTestCtx(defaultDbName string, fn func(ctx *TestCtx)) {
	dbConn, session := db(defaultDbName)

	ctx := TestCtx{Db: dbConn, S: session}
	fn(&ctx)
	session.Close()
}

func db(defaultDbName string) (*mgo.Database, *mgo.Session) {
	url, dbName := GetMongoConfg(defaultDbName)
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	db := session.DB(dbName)
	db.DropDatabase()
	return db, session
}
