package rest

import (
	"github.com/rosbit/http-helper"
	"net/http"
	"fmt"
	"trie-server/trie"
)

// GET /get/:key
func Get(c *helper.Context) {
	key := c.Param("key")
	meta, ok := trie.Get(key)
	if !ok {
		c.Error(http.StatusNotFound, fmt.Sprintf("meta of key %s not found", key))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg": "OK",
		"key": key,
		"meta": meta,
	})
}
