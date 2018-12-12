package tagmap

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
	"configuration"
	//"reflect" // for reflect.TypeOf(<myobject>)
	//"models/modelconstants"
)

var (
	dbname     = "platform"
	collection = "tagbase"
)

/**
All dao functions return status(true/false), error, data

**/

/*
func DeletePerson(person Person) bool {

	fmt.Println("Going to delete a person type records")

	session, err := mgo.Dial(configuration.GetConfiguration("local")["MongoUrl"])
	if err != nil {
		panic(err)
		return false
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbname).C(collection)

	err = c.Remove(person)

	if err != nil {
		panic(err)
		return false
	}
	return true
}


func UpdatePerson(person Person) bool {

	fmt.Println("Going to update a person record")

	session, err := mgo.Dial(configuration.GetConfiguration("local")["MongoUrl"])
	if err != nil {
		panic(err)
		return false
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbname).C(collection)

	err = c.Update(bson.M{"emailid": person.EmailId}, person)

	if err != nil {
		panic(err)
		return false
	}
	return true

}

*/

func SaveTagMap(tagMap TagMap) (bool, string, interface{}) {
	fmt.Println("configuration is -- ", configuration.GetConfiguration("local")["MongoPort"])
	session, err := mgo.Dial(configuration.GetConfiguration("local")["MongoUrl"])
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbname).C(collection)

	err = c.Insert(tagMap)
	if err != nil {
		if mgo.IsDup(err) {
			return false, "This can't happen, because no key is there for tags, this section should never execute :)", nil
			// Is a duplicate key, but we don't know which one
		}
		panic(err)
	}
	return true, "", nil
}

func GetTagData(emailid string, tagList []string) (bool, string, []map[string]interface{}) {

	fmt.Println(configuration.GetConfiguration("local")["MongoPort"])
	session, err := mgo.Dial(configuration.GetConfiguration("local")["MongoUrl"])
	if err != nil {
		panic(err)
		return false, "can't connect to db", nil
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbname).C(collection)

	data := make([]map[string]interface{}, 0, 1000)
	fmt.Println("going to search tags for emailid/tags - ", emailid, tagList)

	q := new([]interface{})

	for _, element := range tagList {
		*q = append(*q, bson.M{"tags": element})
	}

	fmt.Println("your query is : ", q)
	err = c.Find(bson.M{"emailid": emailid, "$or": *q}).All(&data)

	if err != nil {
		fmt.Println("some error occurred while fetching tags for emailid/tags, ", emailid, tagList, err)
		return false, "error occurred, check logs", data
	}
	return true, "success", data
}

/*
func GetPersons(person Person) []map[string]interface{} {
	fmt.Println(configuration.GetConfiguration("local")["MongoPort"])
	session, err := mgo.Dial(configuration.GetConfiguration("local")["MongoUrl"])
	if err != nil {
		panic(err)
		return nil
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbname).C(collection)

	data := []map[string]interface{}{}
	fmt.Println("going to search db for emailid - ", person.EmailId)
	err = c.Find(person).All(&data)

	if err != nil {
		fmt.Println("some error occurred while fetching person, ", err)
		return data
	}
	return data
}
*/
