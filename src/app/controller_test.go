package app

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"util"
)

func TestAddDataIntoCache(t *testing.T) {
	assert := assert.New(t)
	var jsonStr = []byte(`{"key":"name","value":"prabhu"}`)
	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(control.save)
	util.Size = 10
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
		t.Errorf("response body %v", rr.Body.String())

	}

	assert.Equal(http.StatusCreated, rr.Code)

	req, err = http.NewRequest("GET", "/fetchall", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(control.getAll)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"key":"name","value":"prabhu"}]`
	assert.Equal(http.StatusOK, rr.Code)
	assert.Equal(expected, rr.Body.String())
}

func TestUpdateDataInCache(t *testing.T) {
	assert := assert.New(t)
	var jsonStr = []byte(`{"key":"name","value":"namachivayam"}`)
	req, err := http.NewRequest("PUT", "/update", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(control.update)
	util.Size = 10
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Errorf("response body %v", rr.Body.String())

	}

	assert.Equal(http.StatusOK, rr.Code)

	req, err = http.NewRequest("GET", "/fetchall", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(control.getAll)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"key":"name","value":"namachivayam"}]`
	assert.Equal(http.StatusOK, rr.Code)
	assert.Equal(expected, rr.Body.String())

}

func TestDeleteDataFromCache(t *testing.T) {
	assert := assert.New(t)
	var jsonStr = []byte(`{"key":"name","value":"prabhu"}`)
	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(control.save)
	util.Size = 10
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	assert.Equal(http.StatusCreated, rr.Code)

	req, err = http.NewRequest("DELETE", "/remove/name", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"key": "name",
	}
	req = mux.SetURLVars(req, vars)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(control.remove)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
		t.Errorf("response body %v", rr.Body.String())

	}
	assert.Equal(http.StatusNoContent, rr.Code)

	req, err = http.NewRequest("GET", "/fetch/name", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, vars)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(control.get)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	expected := `{"error":"Key Not Found"}`
	assert.Equal(http.StatusNotFound, rr.Code)
	assert.Equal(expected, rr.Body.String())

}
