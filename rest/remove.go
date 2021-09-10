package rest

import (
	"github.com/rosbit/http-helper"
	"net/http"
	"strings"
	"fmt"
	"trie-server/trie"
)

// DELETE /del/:key
func Remove(c *helper.Context) {
	key := strings.TrimSpace(c.Param("key"))
	fmt.Printf("key: %s\n", key)
	if len(key) > 0 {
		trie.Remove(key)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg": "OK",
	})
}
