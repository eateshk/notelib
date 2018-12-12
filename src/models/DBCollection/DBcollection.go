package DBCollection
import (
	"gopkg.in/mgo.v2"
	"fmt"
	"configuration"
)



func GetSession() *mgo.Session{

	fmt.Println("configuration is -- ", configuration.GetConfiguration("local")["MongoPort"])
	session, err := mgo.Dial(configuration.GetConfiguration("local")["MongoUrl"])
	if err != nil {
		panic(err)
		return nil
	}

	session.SetMode(mgo.Monotonic, true)

	return session

}


func CloseSession(session *mgo.Session){

	session.Close()
}

type DBCollection struct {
	DBname string
	Collection string
}

func (dbc *DBCollection) GetDBSession(session *mgo.Session) *mgo.Collection{

	return session.DB(dbc.DBname).C(dbc.Collection)

}
