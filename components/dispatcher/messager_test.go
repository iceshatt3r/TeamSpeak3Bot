package dispatcher

import (
	"fmt"
	"testing"
)

func TestCleanup(t *testing.T) {
	url := `https:\/\/www.youtube.com\/playlist?list=PLnMdq4ssEONRAB4tJI8BuM4OeGuF1XqJO`
	fmt.Println(cleanUpURL(url))
}
