package helper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

type Client struct {
	client *http.Client
	root   string
}

func NewClient(
	t *testing.T,
) Client {
	root := os.Getenv("E2E_TEST_ENDPOINT")

	return Client{
		client: http.DefaultClient,
		root:   root,
	}
}

func (c Client) Request(
	t *testing.T,
	method string,
	path string,
	req any,
) (int, []byte) {
	t.Helper()

	var body io.Reader

	if req != nil {
		b, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}
		body = bytes.NewBuffer(b)
	}

	url := strings.Join([]string{c.root, path}, "/")

	r, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}

	r.Header.Set("Content-Type", "application/json")

	return c.do(t, r)
}

func (c Client) do(
	t *testing.T,
	r *http.Request,
) (int, []byte) {
	t.Helper()

	res, err := c.client.Do(r)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	return res.StatusCode, b
}
