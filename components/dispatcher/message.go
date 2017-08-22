package dispatcher

import (
	"fmt"
	"strings"

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
	executeCommand(command)
}

func executeCommand(command string) {
	switch command {
	case "!addsong":
		fmt.Println(command)
	default:
		fmt.Println("Unknow action")
	}
}
