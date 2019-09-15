package work

import (
	"fmt"
	"net"
)

func logError(key string, err error) {
	if _, ok := err.(*net.OpError); !ok {
		fmt.Printf("ERROR: %s - %s\n", key, err.Error())
	}
}
