package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFibonacciHandler(t *testing.T) {
	tests := []struct {
		name           string
		queryString    string
		wantStatusCode int
		wantBody       []int
		wantErrBody    string
	}{
		{
			name:           "Default (n=10)",
			queryString:    "",
			wantStatusCode: http.StatusOK,
			wantBody:       []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
		{
			name:           "Specific n=5",
			queryString:    "?n=5",
			wantStatusCode: http.StatusOK,
			wantBody:       []int{0, 1, 1, 2, 3},
		},
		{
			name:           "Invalid n (string)",
			queryString:    "?n=abc",
			wantStatusCode: http.StatusBadRequest,
			wantErrBody:    "Invalid parameter 'n'\n",
		},
		{
			name:           "Invalid n (negative)",
			queryString:    "?n=-1",
			wantStatusCode: http.StatusBadRequest,
			wantErrBody:    "Invalid parameter 'n'\n",
		},
		{
			name:           "n too large",
			queryString:    "?n=91",
			wantStatusCode: http.StatusBadRequest,
			wantErrBody:    "Parameter 'n' too large (max 90)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/fibonacci"+tt.queryString, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := fibonacciHandler()

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatusCode)
			}

			if tt.wantStatusCode == http.StatusOK {
				var gotBody []int
				if err := json.Unmarshal(rr.Body.Bytes(), &gotBody); err != nil {
					t.Errorf("failed to unmarshal response body: %v", err)
				}
				if !reflect.DeepEqual(gotBody, tt.wantBody) {
					t.Errorf("handler returned unexpected body: got %v want %v",
						gotBody, tt.wantBody)
				}
			} else {
				if rr.Body.String() != tt.wantErrBody {
					t.Errorf("handler returned unexpected error body: got %q want %q",
						rr.Body.String(), tt.wantErrBody)
				}
			}
		})
	}
}
