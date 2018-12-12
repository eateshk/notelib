package tagmap

import (
//"container/list"
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
)

type TagMap struct {
	Tags       []string          `json:"tags" binding:"required" bson:_id,omitempty"`
	EmailId    string            `json:"emailId" binding:"required"`
	TagData    string            `json:"tagData" binding:"required"`
	UpdateTime uint              `json:"dob" bson:"dob,omitempty"`       // this is not mandatory, we'll fil it from our end
	Others     map[string]string `json:"others" bson:"others,omitempty"` // this is for more stuff if we want to put later

}

func mapper(tm TagMap) map[string]interface{} {

	tagMap := make(map[string]interface{})

	tagMap["Tags"] = tm.Tags
	tagMap["EmailId"] = tm.EmailId
	tagMap["TagData"] = tm.TagData
	tagMap["UpdateTime"] = tm.UpdateTime
	tagMap["Others"] = tm.Others

	return tagMap
}
