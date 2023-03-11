package sanple_test

import (
	"encoding/json"
	"gomulticache/e2e/helper"
	"net/http"
	"testing"
)

func TestE2ESamplePost(t *testing.T) {
	t.Parallel()

	t.Run("200", func(t *testing.T) {
		t.Parallel()

		client := helper.NewClient(t)

		type Req struct {
			Name string `json:"name"`
		}

		param := "post"

		req := &Req{
			Name: param,
		}

		code, res := client.Request(t, http.MethodPost, "api/v1/samples/", req)
		if code != http.StatusOK {
			t.Errorf("status code is %d", code)
		}

		type Res struct {
			Sample struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"sample"`
		}

		var body Res

		json.Unmarshal(res, &body)

		if body.Sample.ID == "" {
			t.Errorf("sample id is empty")
		}
		if body.Sample.Name != param {
			t.Errorf("sample name is %s", body.Sample.Name)
		}
	})
}
