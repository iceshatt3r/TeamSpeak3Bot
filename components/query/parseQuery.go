package query

import (
	"fmt"
	"strings"
)

//Response , represents telnet response
type Response struct {
	Action string
	Params []map[string]string
}

//TSerror , prase string errot into Error()
type TSerror struct {
	id  string
	msg string
}

func (e TSerror) Error() string {
	return fmt.Sprintf("Error from telnet: %s %s", e.id, e.msg)
}

//FormatResponse , Formats output from telnet into Reponse struct
func FormatResponse(s string, action string) *Response {
	r := &Response{}
	var splitResponse []string

	if action == "cmd" {
		r.Action = "Cmd_Response"
		splitResponse = strings.Split(s, "|")
	} else {
		notifystr := strings.SplitN(s, " ", 2)
		r.Action = notifystr[0]
		splitResponse = strings.Split(notifystr[1], "|")
	}
	for i := range splitResponse {
		r.Params = append(r.Params, make(map[string]string))
		splitWhiteSpaces := strings.Split(splitResponse[i], " ")

		for j := range splitWhiteSpaces {
			splitParams := strings.SplitN(splitWhiteSpaces[j], "=", 2)
			if len(splitParams) > 1 {
				r.Params[i][splitParams[0]] = unescape(splitParams[1])
			} else {
				r.Params[i][splitParams[0]] = ""
			}
		}
	}
	return r
}

//FormatError , converts telnet error string to error struct
func FormatError(s string) error {
	e := &TSerror{}
	errorSplit := strings.Split(s, " ")
	for i := range errorSplit {

		eParams := strings.SplitN(errorSplit[i], "=", 2)
		if len(eParams) > 1 {
			if eParams[0] == "id" {
				e.id = eParams[1]
			} else if eParams[0] == "msg" {
				e.msg = unescape(eParams[1])
			}
		} else {
			continue
		}

	}
	if e.id != "0" && e.id != "" {
		return e
	}
	return nil
}
