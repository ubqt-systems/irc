package main

// Just a bunch of file write utilities to handle specific scenarios

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
	"text/template"
)

type Data struct {
	Name    string
	Message string
	file    string
	srv     string
}

func NewData(name, message, srv, fileprefix, filesuffix string) *Data {
	filePath := path.Join(fileprefix, filesuffix)
	return &Data{Name: name, Message: message, file: filePath, srv: srv}
}

// Make sure we have good paths for file writes
func init() {
	if _, err := os.Stat(*inPath); os.IsNotExist(err) {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		*inPath = path.Join(usr.HomeDir, *inPath)
	}
	if _, err := os.Stat(*config); os.IsNotExist(err) {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		*config = path.Join(usr.HomeDir, *config)
	}
}

func writeToFile(nick string, d *Data, format *template.Template) {
	filepath := path.Join(*inPath, d.srv, d.file)
	dirpath := path.Dir(filepath)
	// Make sure path to file exists
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		os.MkdirAll(dirpath, 0755)
	}
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		log.Println(err)
		return
	}
	err = format.Execute(f, d)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(f, "\n")
	switch { // This isn't the most pretty thing, find a better way in the future
	case strings.Contains(d.Message, nick):
		dirpath = path.Dir(d.file)
		filepath = path.Join(dirpath, "notification")
		d2 := &Data{srv: d.srv, file: filepath}
		writeToEvent(d2)
		WriteToNotification(format, d)
	case d.file == "notification":
		WriteToNotification(format, d)
	}
}

func WriteToNotification(format *template.Template, d *Data) {
	filepath := path.Join(*inPath, d.srv, d.file)
	dirpath := path.Dir(filepath)
	filepath = path.Join(dirpath, "notification")
	// Make sure path to file exists
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		os.MkdirAll(dirpath, 0755)
	}
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		log.Println(err)
		return
	}
	err = format.Execute(f, d)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(f, "\n")
}

func writeToEvent(d *Data) {
	paths := strings.Split(d.srv, string(os.PathSeparator))
	filepath := path.Join(*inPath, paths[0], "event")
	dirpath := path.Dir(filepath)
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		os.MkdirAll(dirpath, 0755)
	}
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(f, "%s\n", path.Join(d.srv, d.file))
}

func msgToFile(buff, msg string) {
	filepath := path.Join(*inPath, buff)
	dirpath := path.Dir(filepath)
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		os.MkdirAll(dirpath, 0755)
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(f, msg+"\n")
}

func msgToEvent(buff string) {
	paths := strings.Split(buff, string(os.PathSeparator))
	dirpath := path.Join(*inPath, paths[0])
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		os.MkdirAll(dirpath, 0755)
	}
	filepath := path.Join(dirpath, "event")
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(f, "%s\n", buff)
}

func setTopic(srv, buff, topic string) {
	filepath := path.Join(*inPath, srv, buff, "title")
	dirpath := path.Dir(filepath)
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		os.MkdirAll(dirpath, 0755)
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(f, topic+"\n")
	msgToEvent(path.Join(srv, buff, "title"))
}
