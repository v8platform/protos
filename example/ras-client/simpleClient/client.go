package simpleClient

import (
	"context"
	clientv1 "go.buf.build/v8platform/go-gen-ras/v8platform/rasapis/ras/client/v1"
	protocolv1 "go.buf.build/v8platform/go-gen-ras/v8platform/rasapis/ras/protocol/v1"
)

var defaultVersion = "10.0"

type Client struct {
	host    string
	client  clientv1.ClientServiceImpl
	version string
}

func NewClient(host string) *Client {

	return &Client{
		host:    host,
		version: defaultVersion,
	}

}

func (c *Client) Connect(ctx context.Context) (err error) {

	c.client = clientv1.NewClientService(c.host)
	client := c.client

	_, err = client.Negotiate(protocolv1.NewNegotiateMessage())
	if err != nil {
		return err
	}

	_, err = client.Connect(&protocolv1.ConnectMessage{})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Open(version string) (endpoint clientv1.EndpointServiceImpl, err error) {

	EndpointOpenAck, err := c.client.EndpointOpen(&protocolv1.EndpointOpen{
		Service: "v8.service.Admin.Cluster",
		Version: version,
	})

	if err != nil {
		if version := c.client.DetectSupportedVersion(err); len(version) == 0 {
			return nil, err
		}
		if EndpointOpenAck, err = c.client.EndpointOpen(&protocolv1.EndpointOpen{
			Service: "v8.service.Admin.Cluster",
			Version: version,
		}); err != nil {
			return nil, err
		}
	}

	end, err := c.client.NewEndpoint(EndpointOpenAck)
	if err != nil {
		return nil, err
	}

	return clientv1.NewEndpointService(c.client, end), nil
}

func (c *Client) Close() error {
	_, err := c.client.Disconnect(&protocolv1.DisconnectMessage{})
	if err != nil {
		return err
	}

	return nil
}
