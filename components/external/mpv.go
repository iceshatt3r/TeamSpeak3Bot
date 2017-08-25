package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const externalurl = "url"

type Action struct {
	SocketName string
	ActionName string   `json:"action"`
	Arguments  []string `json:"arguments"`
}

func (a *Action) Send() error {
	url := a.endPoint()
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return makeRequest(url, data)
}

func (a *Action) endPoint() string {
	return fmt.Sprintf(`%s/player/%s/song/add`, externalurl, a.SocketName)
}

func makeRequest(url string, data []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Token", "<secrekt>")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code: %d", resp.StatusCode)
	}
	return nil
}
