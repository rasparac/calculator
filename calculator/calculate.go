package calculator

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Calculator) calculate(w http.ResponseWriter, req *http.Request, operation Operation) {
	log.Printf("msg:request received operation:%s status:success \n", operation)

	x, y, err := readQueryParams(req.URL.Query())
	if err != nil {
		log.Printf("msg:request finished operation: %s status:error error: %v \n", operation, err)
		encodeError(w, jsonErr{
			Msg:     http.StatusText(http.StatusUnprocessableEntity),
			Details: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		})
		return
	}

	key := generateCacheKey(operation, x, y)
	var result float64
	var found bool
	result, found = c.cache.Get(key)
	if !found {
		result = calculateResult(x, y, operation)
		c.cache.Set(key, result)
	}
	jsonResult := jsonResult{
		Action: string(operation),
		Answer: result,
		X:      x,
		Y:      y,
		Cached: found,
	}
	if err := encode(w, jsonResult); err != nil {
		log.Printf("msg:request finished operation:%s status:error error:%v \n", operation, err)
		encodeError(w, jsonErr{
			Msg:     http.StatusText(http.StatusInternalServerError),
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	log.Printf("msg:request received operation:%s status:success \n", operation)
}

func readQueryParams(v url.Values) (float64, float64, error) {
	x, err := strconv.ParseFloat(v.Get("x"), 64)
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.ParseFloat(v.Get("y"), 64)
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func calculateResult(x, y float64, action Operation) float64 {
	switch action {
	case add:
		return x + y
	case divide:
		return x / y
	case subtract:
		return x - y
	case multiply:
		return x * y
	default:
		return 0
	}
}

func encode(w http.ResponseWriter, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func generateCacheKey(operation Operation, x, y float64) string {
	if operation == add || operation == multiply {
		if y > x {
			return fmt.Sprintf("%s_%v_%v", operation, y, x)
		}
		return fmt.Sprintf("%s_%v_%v", operation, x, y)
	}
	return fmt.Sprintf("%s_%v_%v", operation, x, y)
}
