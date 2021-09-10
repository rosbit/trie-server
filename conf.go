/**
 * global conf
 * ENV:
 *   LISTEN_PORT    --- 端口号
 *   TZ             --- 时区名称"Asia/Shanghai"
 *
 * Rosbit Xu
 */
package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
)

var (
	ListenPort = 7080
	Loc = time.FixedZone("UTC+8", 8*60*60)
)

func getEnv(name string, result *string, must bool) error {
	s := os.Getenv(name)
	if s == "" {
		if must {
			return fmt.Errorf("env \"%s\" not set", name)
		}
	}
	*result = s
	return nil
}

func CheckGlobalConf() error {
	var p string
	getEnv("TZ", &p, false)
	if p != "" {
		if loc, err := time.LoadLocation(p); err == nil {
			Loc = loc
		}
	}

	getEnv("LISTEN_PORT", &p, false)
	if p != "" {
		port, err := strconv.Atoi(p)
		if err != nil {
			return err
		}
		if port <= 0 {
			return fmt.Errorf("listening port must be greater than 0")
		}
		ListenPort = port
	}

	return nil
}

func DumpConf() {
	fmt.Printf("ListenPort: %v\n", ListenPort)
	fmt.Printf("TZ time location: %v\n", Loc)
}
