package simple

import "fmt"

type Connection struct {
	File *File
}

func NewConnection(file *File) (*Connection, func()) { // The key to initialize a cleanup function is to RETURNS A FUNCTION
	connection := &Connection{File: file}

	return connection, func() {
		connection.Close()
	}
}

func (c *Connection) Close() {
	fmt.Println("Connection " + c.File.Name + " has been closed")
}
