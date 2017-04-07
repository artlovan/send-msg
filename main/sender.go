package main

import (
	"os"
	"io/ioutil"
	"time"
	"net/http"
	"encoding/json"
	"bytes"
	"log"
	"flag"
	"runtime"
	"path"
)

var target = flag.String("watchedPath", currDir() + "/../source", "Path that will be watched for new files.")
var sendTo = flag.String("postUrl", "http://localhost:6061/api/msg", "Url to where the msg should be sent.")


func currDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}


func send(data []byte) (err error) {
	log.Println(">>> Sending msg")
	msg := map[string]string{"msg": string(data)}
	jsonValue, _ := json.Marshal(msg)
	_, err = http.Post(*sendTo, "application/json", bytes.NewBuffer(jsonValue))

	if err == nil {
		log.Println("<<< Msg sent succesfully")
	}
	return
}


func getFileData(filePath string) (data []byte) {
	f, _ := os.Open(filePath)
	data, _ = ioutil.ReadAll(f)
	defer f.Close()
	return
}


func main() {
	for {
		d, _ := os.Open(*target)
		files, _ := d.Readdir(-1)

		for _, fi := range files {
			filePath := *target + "/" + fi.Name()
			data := getFileData(filePath)
			send(data)
			os.Remove(filePath)
		}
		d.Close()
		time.Sleep(100 * time.Millisecond)
	}

}
