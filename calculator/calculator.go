package calculator

import (
	"encoding/json"
	"net/http"
	"strings"
)

type (
	// Calculator holds main application logic to run math operations
	Calculator struct {
		m      *http.ServeMux
		cache  Cacher
		routes []route
	}

	// Cacher interface is use to cache math opertaion problem resutls
	Cacher interface {
		Get(string) (float64, bool)
		Set(string, float64)
	}

	// Operation holds math operation that can be use by calculator
	Operation string

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

	// route holds info about http method and math operation
	route struct {
		method, operation string
	}
)

const (
	multiply Operation = "multiply"
	add      Operation = "add"
	divide   Operation = "divide"
	subtract Operation = "subtract"
)

//New returns new calculator instance
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

// ServeHTTP implements Handler interface so we can use Calculator as HTTP server
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
	// cover other cases with default mux
	r.m.ServeHTTP(w, req)
}

// initOperations initalize calculator operations
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
