package dispatcher

import "github.com/Overflow3D/ts3Bot_v2/components/query"

// Dispatch send event into correct dispatch function
func Dispatch(r *query.Response) {
	switch r.Action {
	case "notifytextmessage":
		dispatchMessage(r)
	}

}
