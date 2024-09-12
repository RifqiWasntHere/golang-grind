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

// //////////////Sync.Map////////////// //

// var wg = sync.WaitGroup

func AddDataToMap(data *sync.Map, payload int, wg *sync.WaitGroup) {
	defer wg.Done()

	data.Store(payload, payload)
}

func TestAddDataToMap(t *testing.T) {
	data := &sync.Map{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(payload int) {

			AddDataToMap(data, payload, wg)
		}(i)
	}

	wg.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
