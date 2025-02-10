package presenters_test

import (
	"bytes"
	"encoding/json"
	"maps"
	"testing"

	"github.com/cl1ckname/cdf/internal/pkg/domain"
	"github.com/cl1ckname/cdf/internal/pkg/presenters"
)

func TestJSON(t *testing.T) {
	wr := bytes.NewBuffer(nil)
	js := presenters.NewJSON(wr, nil)
	marks := []domain.Mark{
		{Alias: "h", Path: "/home/username"},
		{Alias: "prj", Path: "/home/username/projects"},
	}
	err := js.Present(marks)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	bts := wr.Bytes()
	var res map[string]string
	if err := json.Unmarshal(bts, &res); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	exp := map[string]string{}
	for _, mark := range marks {
		exp[mark.Alias] = mark.Path
	}

	if !maps.Equal(exp, res) {
		t.Fatalf("wrong result, expected %v got %v", exp, res)
	}
}
