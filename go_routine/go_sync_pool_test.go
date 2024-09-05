package go_routine

import (
	"fmt"
	"sync"
	"testing"
)

type DBConnection struct {
	creds string
}

var credsId string
var connPool *sync.Pool

func init() {
	connPool = &sync.Pool{
		New: func() interface{} {
			credsId += "aku"
			return &DBConnection{creds: credsId}
		},
	}
}

func TestGetDBCreds(t *testing.T) {
	conn := connPool.Get()
	fmt.Println(conn)
}
