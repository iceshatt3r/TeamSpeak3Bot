package dispatcher

import (
	"fmt"
	"strings"

	"github.com/Overflow3D/ts3Bot_v2/components/external"
	"github.com/Overflow3D/ts3Bot_v2/components/query"
)

func dispatchMessage(r *query.Response) {
	indexOfWhiteSpace := strings.Index(r.Params[0]["msg"], " ")
	if indexOfWhiteSpace == -1 {
		indexOfWhiteSpace = len(r.Params[0]["msg"])
	}
	command := strings.ToLower(r.Params[0]["msg"][:indexOfWhiteSpace])
	if strings.Index(command, "!") != 0 {
		return
	}
	executeCommand(command, r)
}

func executeCommand(command string, r *query.Response) {
	slicedMsg := strings.Split(r.Params[0]["msg"], " ")
	switch command {
	case "!addsong":
		fmt.Println(slicedMsg)
		addSong(slicedMsg)
	default:
		fmt.Println("Unknow action")
	}
}

func addSong(slicedMsg []string) {
	if len(slicedMsg) < 3 {
		return
	}
	action := &external.Action{SocketName: slicedMsg[1], ActionName: "loadfile", Arguments: []string{removeBBCodeTags(slicedMsg[2]), "append"}}
	if err := action.Send(); err != nil {
		fmt.Println(err)
	}
}

func removeBBCodeTags(url string) string {
	if strings.Contains(url, "[") {
		return url[5 : len(url)-7]
	}
	return url
}
