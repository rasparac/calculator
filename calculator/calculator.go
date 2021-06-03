package calculator

import (
	"encoding/json"
	"net/http"
	"strings"
)

type (
	Calculator struct {
		m      *http.ServeMux
		cache  Cacher
		routes []route
	}

	route struct {
		method, operation string
	}

	Cacher interface {
		Get(string) (float64, bool)
		Set(string, float64)
	}

	jsonResult struct {
		Action string  `json:"action"`
		Answer float64 `json:"answer"`
		X      float64 `json:"x"`
		Y      float64 `json:"y"`
		Cached bool    `json:"cached"`
	}

	jsonErr struct {
		Msg     string `json:"message"`
		Details string `json:"details"`
		Code    int    `json:"code"`
	}

	Operation string
)

const (
	multiply Operation = "multiply"
	add      Operation = "add"
	divide   Operation = "divide"
	subtract Operation = "subtract"
)

func New() *Calculator {
	m := http.NewServeMux()
	r := &Calculator{
		m:      m,
		routes: []route{},
		cache:  NewCache(),
	}
	r.initOperations()
	return r
}

func (r *Calculator) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	setJSONHeader(w)

	var isOperationNotFound, isNotAllowed bool
	for _, route := range r.routes {
		if route.operation == req.URL.Path {
			if req.Method != route.method {
				isNotAllowed = true
				break
			}
			operation := Operation(strings.Trim(route.operation, "/"))
			r.calculate(w, req, operation)
			return
		}
	}

	if isNotAllowed {
		w.Header().Set("Access-Control-Allow-Methods", http.MethodGet)
		encodeError(w, jsonErr{
			Details: "only GET method allowed",
			Msg:     http.StatusText(http.StatusMethodNotAllowed),
			Code:    http.StatusMethodNotAllowed,
		})
		return
	}

	if !isOperationNotFound {
		encodeError(w, jsonErr{
			Details: "math operation not implemented",
			Msg:     http.StatusText(http.StatusNotFound),
			Code:    http.StatusNotFound,
		})
		return
	}
	r.m.ServeHTTP(w, req)
}

func (r *Calculator) initOperations() {
	r.addCalculatorOperation(http.MethodGet, string(add))
	r.addCalculatorOperation(http.MethodGet, string(subtract))
	r.addCalculatorOperation(http.MethodGet, string(multiply))
	r.addCalculatorOperation(http.MethodGet, string(divide))
}

func (r *Calculator) addCalculatorOperation(verb, operation string) {
	newRoute := route{
		method:    verb,
		operation: "/" + operation,
	}
	r.routes = append(r.routes, newRoute)
}

func setJSONHeader(w http.ResponseWriter) {
	w.Header().Add("Content-type", "application/json")
}

func encodeError(w http.ResponseWriter, err jsonErr) error {
	w.WriteHeader(err.Code)
	return json.NewEncoder(w).Encode(err)
}
