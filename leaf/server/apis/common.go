package apis

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetInt64Param(req *http.Request, key string) (int64, error) {
	value := req.FormValue(key)
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("query param %s[%s] not is be int err:%w", key, value, err)
	}
	return int64(number), nil
}

func Pong(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	return 200, "pong", nil
}

func Test(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	return 200, "test", nil
}
