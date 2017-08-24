package query

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
)

const defaultPort = "10011"

type TelNet struct {
	conn    net.Conn
	scanner *bufio.Scanner
	output  chan string
	err     chan string
	cmdRes  string
	Notify  chan string
}

//NewServerQuery , create a new server query connection
func NewServerQuery(addr string, isListener bool) (*TelNet, error) {
	telnet := &TelNet{}
	telnet.connect(addr)
	telnet.scanConnection()
	if !telnet.checkAndDiscard() {
		return nil, errors.New("Couldn't connect to teamspeak 3 server on adress: " + addr)
	}
	go telnet.checkResponseOutPut(isListener)

	return telnet, nil
}

func (t *TelNet) connect(adress string) error {
	if !strings.Contains(adress, ":") {
		adress = adress + ":" + defaultPort
	}
	con, err := net.Dial("tcp", adress)
	if err != nil {
		return nil
	}
	t.conn = con
	return nil
}

func (t *TelNet) scanConnection() {
	t.scanner = bufio.NewScanner(t.conn)
	t.scanner.Split(scanTelNet)
	t.output = make(chan string)
	go func() {
		for {
			t.scanner.Scan()
			t.output <- t.scanner.Text()
			e := t.scanner.Err()
			if e != nil {
				return
			}
		}
	}()
}

func scanTelNet(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\r")); i >= 0 {
		return i + 2, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func (t *TelNet) checkAndDiscard() bool {
	output := <-t.output
	if output != "TS3" {
		return false
	}
	<-t.output
	return true
}

func (t *TelNet) checkResponseOutPut(isListener bool) {
	t.err = make(chan string)
	t.Notify = make(chan string)
	for {
		output := <-t.output
		if strings.Index(output, "error") == 0 {
			t.err <- output
		} else if strings.Index(output, "notify") == 0 && isListener {
			t.Notify <- output
		} else {
			t.cmdRes = output
		}
	}

}

//Exec , executes command to serverQuery
func (t *TelNet) Exec(cmd *Command) (*Response, error) {
	fmt.Fprintf(t.conn, "%s\n\r", cmd)
	err := <-t.err
	return FormatResponse(t.cmdRes, ""), FormatError(err)
}

//ExecMultiple , executes multiple commands to serverQuery
//silence -> determinates if there will be console output or not
func (t *TelNet) ExecMultiple(cmd []*Command, silence bool) {
	for _, c := range cmd {
		fmt.Fprintf(t.conn, "%s\n\r", c)
		<-t.err
	}
}

//Close , closes the server query connection
func (t *TelNet) Close() {
	t.conn.Close()
}
