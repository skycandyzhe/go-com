package httpclient

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

var post_form = "http://0.0.0.0:9980/whales/v1/post/test"
var post_task_url = "http://0.0.0.0:9980/whales/v1/task/submit"

func TestPost(t *testing.T) {
	body, err := PostForm(
		post_form,
		url.Values{
			"id":   {"123"},
			"name": {"switch"},
		},
	)
	if err != nil {
		t.Fail()
		return
	}

	fmt.Println(string(body))
}

func TestPostFile(t *testing.T) {
	filepaths := "/Users/switch/Project/dbapp/whl-neptunenavy/whl-utils/tests/data/data.json"
	f, err := os.Open(filepaths)
	if err != nil {
		t.Fail()
		return
	}

	body, err := PostFormFile(
		post_task_url,
		url.Values{
			"timeout": {"15"},
			"dynamic": {"true"},
		},
		FileValues{
			"file": {
				Name:   filepath.Base(filepaths),
				Reader: f,
			},
		},
	)
	if err != nil {
		t.Fail()
		return
	}

	fmt.Println(string(body))
}
