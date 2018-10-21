package apis

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllProducts(t *testing.T) {

	tt := []struct {
		name string
		err  string
	}{
		{name: "Valid request"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8000/products", nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}
			rec := httptest.NewRecorder()

			GetAllProducts(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if tc.err != "" {
				//Check all error cases here
				//Example:
				// if res.StatusCode != http.StatusBadRequest {
				// 	t.Errorf("Expected status Bad Request, got %v", res.StatusCode)
				// }
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("Expected Status Ok, got %v", res.StatusCode)
			}
		})
	}
}
