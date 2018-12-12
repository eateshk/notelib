package person

import (
//"fmt"
//"net/url"
//"net/http"
//"encoding/json"
//"gopkg.in/mgo.v2"
//"gopkg.in/mgo.v2/bson"
//"github.com/dougblack/sleepy"
//"time"
//"math/big"
//"configuration"
//"dao"
	"crypto/md5"
	"encoding/hex"
)

type Person struct {
	EmailId  string            `json:"emailid" binding:"required" bson:"_id,omitempty"`
	Password string            `json:"password" bson:"password,omitempty"`
	Name     string            `json:"name" bson:"name,omitempty"`
	Sex      string            `json:"sex" bson:"sex,omitempty"`
	NickName string            `json:"nickname" bson:"nickname,omitempty"`
	DOB      uint              `json:"dob" bson:"dob,omitempty"`
	Others   map[string]string `json:"others" bson:"others,omitempty"`
}

func mapper(p Person) map[string]interface{} {

	personMap := make(map[string]interface{})

	personMap["Name"] = p.Name
	personMap["Sex"] = p.Sex
	personMap["NickName"] = p.NickName
	personMap["DOB"] = p.DOB
	personMap["Others"] = p.Others

	return personMap
}


func (person *Person) Init(){

	person.Password = getMD5Hash(person.Password)

}


func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}