/**
 * REST API router
 * Rosbit Xu
 */
package main

import (
	"github.com/rosbit/http-helper"
	"net/http"
	"fmt"
	"trie-server/rest"
)

func StartService() error {
	api := helper.NewHelper(helper.WithLogger("trie"))

	api.POST("/append",     rest.Append)
	api.GET("/get/:key",    rest.Get)
	api.DELETE("/del/:key", rest.Remove)
	api.PUT("/clear",    rest.Clear)
	api.POST("/search",  rest.Search)
	api.POST("/prefix",  rest.Prefix)

	listenParam := fmt.Sprintf(":%d", ListenPort)
	fmt.Printf("I am listening at %s...\n", listenParam)
	return http.ListenAndServe(listenParam, api)
}

