package main

import (
	"log"
	"os"
)

func main1() {
	// file, _ := os.Create("file.log")
	// log.SetOutput(file)
	// log.Println("Hello World!...")
	// file.Close()
	// file, _ := os.Create("file.log")
	// log.SetOutput(file)

	flags := log.LstdFlags | log.Lshortfile
	//log.SetFlags(log.Ldate | log.Lshortfile)
	file, _ := os.OpenFile("file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	infoLogger := log.New(file, "INFO ", flags)
	warnLogger := log.New(file, "WARN ", flags)
	errorLogger := log.New(file, "ERROR ", flags)

	log.New(file, "", 0).Println()

	infoLogger.Println("this is info log...")
	warnLogger.Println("this is a warning...")
	errorLogger.Println("something went wrong...")

	//log.Error("some info")
	//log.SetOutput(file)

}
