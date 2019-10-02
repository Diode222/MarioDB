package client

import (
	"bytes"
)

type ClientType int

const (
	NORMAL             ClientType = iota // Means a normal client which will get/put kv data
	MONITOR                                // Means a monitor client which will monitor the server's stats
)

type Client struct {
	Address string
	Buffer *bytes.Buffer
	Type ClientType
}

func (c *Client) GetAddress() string {
	return c.Address
}

func (c *Client) GetBuffer() *bytes.Buffer {
	return c.Buffer
}

func (c *Client) GetType() ClientType {
	return c.Type
}

// Discard buffer has consumed
func (c *Client) DiscardConsumedBuffer(consumeBytesLength int) error {
	var consumed []byte = make([]byte, consumeBytesLength, consumeBytesLength)
	_, err := c.Buffer.Read(consumed)
	return err
}

func (c *Client) Write(message []byte) error {
	_, err := c.Buffer.Write(message)
	return err
}

func (c *Client) ClearBuffer() {
	c.Buffer = bytes.NewBuffer(nil)
}