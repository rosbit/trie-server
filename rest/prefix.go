package rest

import (
	"github.com/rosbit/http-helper"
	"net/http"
	"strings"
	"trie-server/trie"
)

// POST /prefix
// { "text": "any text"}
func Prefix(c *helper.Context) {
	var params struct {
		Text string `json:"text"`
	}
	if code, err := c.ReadJSON(&params); err != nil {
		c.Error(code, err.Error())
		return
	}
	params.Text = strings.TrimSpace(params.Text)
	if len(params.Text) == 0 {
		c.Error(http.StatusBadRequest, "param text expected")
		return
	}
	keys := trie.PrefixSearch(params.Text)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg":  "OK",
		"text": params.Text,
		"prefixes": keys,
	})
}
