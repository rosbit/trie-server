package rest

import (
	"github.com/rosbit/http-helper"
	"net/http"
	"trie-server/trie"
)

// PUT /clear
func Clear(c *helper.Context) {
	trie.Clear()

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg": "OK",
	})
}
