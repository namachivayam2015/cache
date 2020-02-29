package cache

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"test.mydomain.com/cache/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Controller struct {

}

func (c *Controller) save(w http.ResponseWriter, r *http.Request) {
	var data CacheData
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		util.BuildErrorResponse(w,  http.StatusInternalServerError,"Key cannot be empty!")
	}

	if err := json.Unmarshal(body, &data); err != nil {
		util.BuildErrorResponse(w,  http.StatusBadRequest,"Input Data Parse Error")
		return
	}

	if data.Key == "" {
		util.BuildErrorResponse(w,  http.StatusBadRequest,"Key cannot be empty!")
		return
	}

	error := util.AddElement(data.Key, data.Value)
	if error != nil {
		util.BuildErrorResponse(w,  http.StatusBadRequest, error.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func (c *Controller) update(w http.ResponseWriter, r *http.Request) {
	var data CacheData
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		util.BuildErrorResponse(w,  http.StatusInternalServerError,"Key cannot be empty!")
	}

	if err := json.Unmarshal(body, &data); err != nil {
		util.BuildErrorResponse(w,  http.StatusBadRequest,"Input Data Parse Error")
		return
	}

	if data.Key == "" {
		util.BuildErrorResponse(w,  http.StatusBadRequest,"Key cannot be empty!")
		return
	}

	var prevVal = util.GetElement(data.Key)

	if prevVal == "" {
		util.BuildErrorResponse(w,  http.StatusBadRequest,"Invalid Key Supplied!!!")
		return
	}

	util.UpdateElement(data.Key, data.Value)

	w.WriteHeader(http.StatusOK)
	return
}

func (c *Controller) remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var val = util.GetElement(vars["key"])
	if val == "" {
		util.BuildErrorResponse(w,  http.StatusBadRequest,"Invalid Key Supplied!!!")
		return
	}

	util.DeleteElement(vars["key"])

	w.WriteHeader(http.StatusNoContent)
}


func (c *Controller) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var data CacheData
	var val = util.GetElement(vars["key"])
	if val == "" {
		util.BuildErrorResponse(w,  http.StatusNotFound,"Key Not Found")
		return
	}
	data.Key = vars["key"]
	data.Value = val

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	output, _ := json.Marshal(data)
	w.Write(output)
}

func (c *Controller) getAll(w http.ResponseWriter, r *http.Request) {
	var m = util.GetAllElement()
	var size = len(m)
	var counter = 0
	var dataList = make(CacheDataList, size)
	for key, val := range m {
		var data CacheData
		data.Value = val
		data.Key = key
		dataList[counter] = data
		counter++
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	output, _ := json.Marshal(dataList)
	w.Write(output)
}

func (c *Controller) basicAuth(handler http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var user, pass, ok= r.BasicAuth();
		if !ok {
			w.WriteHeader(http.StatusForbidden)
		} else {
			if !(user == os.Getenv("BASIC_AUTH_USERNAME") && pass == os.Getenv("BASIC_AUTH_PASSWORD")) {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				if r.Header.Get("Content-Type") == "application/json" {
					handler(w, r)
				}else{
					w.WriteHeader(http.StatusUnsupportedMediaType)
				}
			}
		}
	}
}
