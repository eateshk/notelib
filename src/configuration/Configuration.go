package configuration

import (
	"fmt"
)



func GetConfiguration(environment string) (map[string]string){
	fmt.Println("someone asked for configuration")	
	var Local map[string]string
	Local = make(map[string]string)

	Local["Id"] = "-1"
	Local["MongoPort"] = "27017"
	Local["MongoNode"] = "127.0.0.1"
	Local["MongoUrl"] = "127.0.0.1:27017"
	switch environment{
		case "local":
			return Local
		default :
			return Local
	}

}

