package http

import (
	"encoding/json"
	"github.com/hb0730/go-check-certs/certs"
	"io/ioutil"
	"net/http"
)

func Request(addr string) error {
	handler()
	return http.ListenAndServe(addr, nil)
}
func handler() {
	http.Handle("/get", getHandler())
	http.Handle("/list", listHandler())
}

func getHandler() http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			write(response, result(false, nil, "invalid http method"))
			return
		}
		addr := request.URL.Query().Get("addr")
		if addr == "" {
			write(response, result(false, nil, "request parameter error"))
			return
		}
		checkCerts(response, addr)
	})
}
func listHandler() http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			write(response, result(false, nil, "invalid http method"))
			return
		}
		defer request.Body.Close()
		addr, err := ioutil.ReadAll(request.Body)
		if err != nil {
			write(response, result(false, nil, "request parameter error"))
			return
		}

		var addrs []string
		err = json.Unmarshal(addr, &addrs)
		if err != nil {
			write(response, result(false, nil, "request parameter error"))
			return
		}
		checkCerts(response, addrs...)
	})
}

func checkCerts(w http.ResponseWriter, addrs ...string) {
	var result = make([]interface{}, 0, len(addrs))
	for _, addr := range addrs {
		cs, err := certs.Check(addr)
		s := struct {
			Host  string       `json:"host"`
			Certs []certs.Cert `json:"certs"`
			Error string       `json:"error"`
		}{}
		s.Host = addr
		if err != nil {
			s.Error = err.Error()
		} else {
			s.Certs = cs
		}

		result = append(result, s)
	}
	write(w, resultSuccess(result))
}

func write(w http.ResponseWriter, rt []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rt)
}
func resultSuccess(data interface{}) (bt []byte) {
	return result(true, data, "")
}
func result(success bool, data interface{}, err string) (rt []byte) {
	r := Result{
		Success: success,
		Data:    data,
		Error:   err,
	}
	rt, _ = json.Marshal(&r)
	return
}

type Result struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}
