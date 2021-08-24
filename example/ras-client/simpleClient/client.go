package simpleClient

import (
	"context"
	"github.com/v8platform/protos/extra"
	"net"
)

var defaultVersion = "10.0"

type Client struct {
	host string
	net.Conn
	endpoint *extra.Endpoint
	version  string
}

func NewClient(host string) *Client {

	return &Client{
		host:    host,
		version: defaultVersion,
	}

}

func (c *Client) Connect(ctx context.Context) (err error) {

	if c.Conn, err = c.dial(ctx); err != nil {
		return err
	}

	err = extra.Connect(c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Open(version string) (endpoint *extra.Endpoint, err error) {
	c.endpoint, err = extra.OpenEndpoint(c, version)
	if err != nil {
		return nil, err
	}

	return c.endpoint, nil
}

func (c *Client) dial(ctx context.Context) (net.Conn, error) {

	_, err := net.ResolveTCPAddr("tcp", c.host)
	if err != nil {
		return nil, err
	}

	var dialer net.Dialer

	conn, err := dialer.DialContext(ctx, "tcp", c.host)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
