package main

import (
	"fmt"

	"strings"

	"github.com/alexflint/go-arg"
	"github.com/nlopes/slack"
)

func main() {
	var args struct {
		Token   string   `arg:"-t,required"`
		Channel string   `arg:"-c,required"`
		Message []string `arg:"positional"`
	}
	arg.MustParse(&args)

	api := slack.New(args.Token)
	// attachment := slack.Attachment{
	// 	// Pretext: "some pretext",
	// 	Text: strings.Join(args.message, " "),
	// 	// Uncomment the following part to send a field too
	// 	/*
	// 		Fields: []slack.AttachmentField{
	// 			slack.AttachmentField{
	// 				Title: "a",
	// 				Value: "no",
	// 			},
	// 		},
	// 	*/
	// }

	channelID, timestamp, err := api.PostMessage(args.Channel, slack.MsgOptionText(strings.Join(args.Message, " "), false) /*, slack.MsgOptionAttachments(attachment)*/)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}
