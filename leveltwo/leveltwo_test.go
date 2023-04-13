package leveltwo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetData(t *testing.T) {

	// Prepare test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Required data")
	}))
	defer ts.Close()

	err := GetData(ts.URL)
	if err != nil {
		t.Errorf("error GetData(), expected no error, got %v", err)
	}

}
