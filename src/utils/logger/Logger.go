package logger
import("fmt")

func Log(debugMode string, toLog ...interface{}) {
	fmt.Print(debugMode, " : ")
	fmt.Println(toLog)
}
