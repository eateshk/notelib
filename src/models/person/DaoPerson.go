package person

import (
	"fmt"
	//"net/url"
	//"net/http"
	//"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/dougblack/sleepy"
	//"time"
	//"math/big"
	//"reflect" // for reflect.TypeOf(<myobject>)
	//"models/modelconstants"
	"models/DBCollection"
	"utils/logger"
)

var (
	dbcollection = &DBCollection.DBCollection{ "platform","person"}
)



/**
All dao functions return status(true/false), error, data

**/



func DeletePerson(person Person) bool {

	session := DBCollection.GetSession()
	defer DBCollection.CloseSession(session)
	dbSession := dbcollection.GetDBSession(session)


	err := dbSession.Remove(person)
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func UpdatePerson(person Person) bool {

	session := DBCollection.GetSession()
	defer DBCollection.CloseSession(session)
	dbSession := dbcollection.GetDBSession(session)

	err := dbSession.Update(bson.M{"emailid": person.EmailId}, person)

	if err != nil {
		panic(err)
		return false
	}
	return true

}

func SavePerson(person Person) (bool, string, interface{}) {

	session := DBCollection.GetSession()
	defer DBCollection.CloseSession(session)
	dbSession := dbcollection.GetDBSession(session)

	person.Init()
	err := dbSession.Insert(person)
	if err != nil {
		if mgo.IsDup(err) {
			return false, "This Person Already Exists", nil
			// Is a duplicate key, but we don't know which one
		}
		panic(err)
	}
	return true, "", nil
}

func GetPerson(emailid string) map[string]interface{} {

	session := DBCollection.GetSession()
	defer DBCollection.CloseSession(session)
	dbSession := dbcollection.GetDBSession(session)

	data := make(map[string]interface{})
	fmt.Println("going to search db for emailid - ", emailid)
	err := dbSession.Find(bson.M{"emailid": emailid}).One(&data)

	if err != nil {
		fmt.Println("some error occurred while fetching person, ", err)
		return data
	}
	return data
}

func GetPersons(person Person) []map[string]interface{} {
	session := DBCollection.GetSession()
	defer DBCollection.CloseSession(session)
	dbSession := dbcollection.GetDBSession(session)

	data := []map[string]interface{}{}
	fmt.Println("going to search db for emailid - ", person.EmailId)
	err := dbSession.Find(person).All(&data)

	if err != nil {
		fmt.Println("some error occurred while fetching person, ", err)
		return data
	}


	return data
}


func (person *Person) Authenticate() (bool,string){

	person.Init()
	session := DBCollection.GetSession()
	defer DBCollection.CloseSession(session)
	dbSession := dbcollection.GetDBSession(session)

	result := Person{}
	err := dbSession.Find(bson.M{"_id": person.EmailId}).One(&result)

	if err != nil{
		logger.Log("DaoPerson","Authentication failed for person due to some error " + person.EmailId + " " +err.Error())
		return false, "Authentication failure, User does not exist"


	}

	return authenticationHelper(person,&result)


}

func authenticationHelper(source *Person, target *Person)(bool, string){

	userExist ,err := false , "Authentication failure, User does not exist"
	if target != nil {
		if target.Password == source.Password {
			userExist , err = true, "Authentication successfull"
		} else
		{
			userExist ,err = false , "Authentication failure, Incorrect Password"
		}

	}

	return userExist , err

}