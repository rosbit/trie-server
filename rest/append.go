package rest

import (
	"github.com/rosbit/http-helper"
	"net/http"
	"strings"
	"trie-server/trie"
)

// POST /append
// {
//    "key": "adafs",
//    "meta": {JSON},
// }
func Append(c *helper.Context) {
	var params struct {
		Key string      `json:"key"`
		Meta interface{} `json:"meta"`
	}

	code, err := c.ReadJSON(&params)
	if err != nil {
		c.Error(code, err.Error())
		return
	}
	params.Key = strings.TrimSpace(params.Key)
	if len(params.Key) == 0 {
		c.Error(http.StatusBadRequest, "params key expected")
		return
	}

	trie.Append(params.Key, params.Meta)
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg": "OK",
	})
}
