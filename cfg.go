package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/template"
	
	"github.com/vaughan0/go-ini"
)

// Hold our default configurations
type Cfg struct { 
	chanFmt *template.Template
	selfFmt *template.Template
	ntfyFmt *template.Template
	servFmt *template.Template
	highFmt *template.Template
	actiFmt *template.Template
	modeFmt *template.Template
}

func ParseFormat(conf ini.File) *Cfg {
    //Set some pretty printed defaults
    chanFmt := `[#5F87A7]({{.Name}}) {{.Data}}`
    selfFmt := `[#076678]({{.Name}}) {{.Data}}`
    highFmt := `[#9d0007]({{.Name}}) {{.Data}}`
    ntfyFmt := `[#5F87A7]({{.Name}}) {{.Data}}`
    servFmt := `--[#5F87A7]({{.Name}}) {{.Data}}--`
    actiFmt := `[#5F87A7( \* {{.Name}}) {{.Data}}`
    modeFmt := `--[#787878](Mode [{{.Data}}] by {{.Name}})`
    for key, value := range conf["options"] {
        switch key {
        case "channelfmt":
            chanFmt = value
        case "notificationfmt":
            ntfyFmt = value
        case "highfmt":
            highFmt = value
        case "selffmt":
            selfFmt = value
        case "actifmt":
            actiFmt = value
        case "modefmt":
            modeFmt = value
        }
    }
	return &Cfg{
    	chanFmt: template.Must(template.New("chan").Parse(chanFmt)),
    	ntfyFmt: template.Must(template.New("ntfy").Parse(ntfyFmt)),
    	servFmt: template.Must(template.New("serv").Parse(servFmt)),
    	selfFmt: template.Must(template.New("self").Parse(selfFmt)),
    	highFmt: template.Must(template.New("high").Parse(highFmt)),
    	actiFmt: template.Must(template.New("acti").Parse(actiFmt)),
    	modeFmt: template.Must(template.New("mode").Parse(modeFmt)),
	}
}

func ParseChannels(conf ini.File, section string) []string {
	channels, _ := conf.Get(section, "Channels")
	return strings.Split(channels, ",")
}

// TODO: Return useful config here
func ParseServer(conf ini.File, section string) () {
    server, ok := conf.Get(section, "Server")
    if ! ok {
        log.Println("Server entry not found!")
    }
    p, ok := conf.Get(section, "Port")
    port, _ := strconv.Atoi(p)
    if !ok {
        fmt.Println("No port set, using default")
        port = 6667
    }
    nick, ok := conf.Get(section, "Nick")
    if !ok {
        fmt.Println("nick entry not found")
    }
    user, ok := conf.Get(section, "User")
    if !ok {
        fmt.Println("user entry not found")
    }
    name, ok := conf.Get(section, "Name")
    if !ok {
        fmt.Println("name entry not found")
    }
    pw, ok := conf.Get(section, "Password")
    if !ok {
        fmt.Println("password entry not found")
    }
    // return &girc.Config{Server: server, Port: port, Nick: nick, User: user, Name: name, ServerPass: pw}
}
