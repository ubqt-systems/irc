package main

import(
	"fmt"
	"github.com/go-irc/irc"
)

func InitHandler(channels string) (irc.Handler) {
	return irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
		switch m.Command {
		// This is sent on server connection, join channels here
		case "001":
			c.Writef("JOIN %s\n", channels) 
		//case "INVITE":	
		//case "NOTICE":
		//case "PRIVMSG":
		// - :ACTION
		// - :TOPIC
		// - :FINGER
		// - etcetera
		//case "JOIN":
		// JOIN for our user implies we're joining a channel. We need to clear out sidebar so we can harvest the name list without a FSM
		//case "PART":
		//case "KICK"
		//case "MODE"
		//case "TIME"
		//case "TOPIC"
		//case "VERSION"
		//case "FINGER"
		//case "USERINFO"
		//case "CLIENTINFO"
		//case "SOURCE"
		//case "301" // <client> <nick> :<message> //away message reply
		//case "305" // no longer away
		//case "306" // now away
		//case "332" // topic - log to channel, as well as set up title
		//case "333" // who set the topic, when - log to channel
		//case "375" // Start of message of the day
		//case "372" // MOTD
		//case "376" // End of MOTD, MODE
		//case "353" // List of names - set up sidebar
		//case "366" // End of name list
		case "PING":
			c.Writef("PONG %s", m.Params[0])
		//case "QUIT"
		default: // Log to server for all other messages so far
			fmt.Printf("%s\n", m.String())
		}
	})
}