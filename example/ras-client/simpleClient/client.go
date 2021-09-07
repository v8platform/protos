package simpleClient

import (
	"context"
	"fmt"
	clientv1 "github.com/v8platform/protos/gen/ras/client/v1"
	protocolv1 "github.com/v8platform/protos/gen/ras/protocol"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var defaultVersion = "10.0"

var _ clientv1.ClientServiceServer = (*Client)(nil)

type Client struct {
	ClientOptions

	host       string
	conn       net.Conn
	onceInit   *sync.Once
	usedAt     uint32 // atomic
	_closed    uint32 // atomic
	_connected uint32 // atomic
	_inited    uint32 // atomic

	endpoints *sync.Map
	stats     Stats
	mu        *sync.Mutex // Блокировка всего клиента
	connMu    *sync.Mutex // Блокировка только соединения
}

type ClientOptions struct {
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
	NegotiateMessage   *protocolv1.NegotiateMessage
	ConnectMessage     *protocolv1.ConnectMessage
	OpenEndpoint       *protocolv1.EndpointOpen
}

type Stats struct {
	Recv  uint32
	Send  uint32
	Wrong uint32
	Ping  uint32
}

var defaultClientOptions = ClientOptions{
	IdleTimeout:        30 * time.Minute,
	IdleCheckFrequency: 5 * time.Minute,
	NegotiateMessage:   protocolv1.NewNegotiateMessage(),
	ConnectMessage:     &protocolv1.ConnectMessage{},
	OpenEndpoint: &protocolv1.EndpointOpen{
		Service: "v8.service.Admin.Cluster",
		Version: defaultVersion,
	},
}

func NewClient(host string, options ...ClientOptions) *Client {

	opt := defaultClientOptions
	if len(options) > 0 {
		opt = options[0]
	}

	return &Client{
		host:          host,
		ClientOptions: opt,
		endpoints:     &sync.Map{},
		mu:            &sync.Mutex{},
		connMu:        &sync.Mutex{},
		onceInit:      &sync.Once{},
	}
}

func (c *Client) reset() error {

	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.Close()
	if err != nil {
		return err
	}

	c.endpoints = &sync.Map{}
	c.onceInit = &sync.Once{}
	c.ClientOptions = *(&defaultClientOptions) // FIX Copy if need

	return nil
}

func (c *Client) getEndpoint(id int32) (*protocolv1.Endpoint, bool) {

	if val, ok := c.endpoints.Load(id); ok {
		return val.(*protocolv1.Endpoint), ok
	}
	return nil, false
}

func (c *Client) addEndpoint(endpoint *protocolv1.Endpoint) (loaded bool) {

	if _, ok := c.endpoints.LoadOrStore(endpoint.GetId(), endpoint); ok {
		return true
	}

	return false

}

func (c *Client) Init(ctx context.Context, request *clientv1.InitRequest) (*clientv1.StatusInfo, error) {

	c.IdleTimeout = time.Duration(request.GetIdleTimeout())

	if request.GetLazy() != nil {

		if val := request.Lazy.GetOpenEndpoint(); val != nil {
			c.OpenEndpoint = val
		}
		if val := request.Lazy.GetNegotiate(); val != nil {
			c.NegotiateMessage = val
		}
		if val := request.Lazy.GetConnect(); val != nil {
			c.ConnectMessage = val
		}
	}

	err := c.init(ctx)
	if err != nil {
		return nil, err
	}

	return new(clientv1.StatusInfo), nil
}

func (c *Client) Status(ctx context.Context, _ *emptypb.Empty) (*clientv1.StatusInfo, error) {
	panic("implement me")
}

func (c *Client) NewEndpoint(ctx context.Context, request *clientv1.NewEndpointRequest) (*protocolv1.Endpoint, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return nil, nil
}

func (c *Client) CloseEndpoint(ctx context.Context, endpoint *protocolv1.Endpoint) (*emptypb.Empty, error) {
	panic("implement me")
}

func (c *Client) Requests(server clientv1.ClientService_RequestsServer) error {
	panic("implement me")
}

func (c *Client) Request(ctx context.Context, request *clientv1.EndpointRequest) (*clientv1.EndpointResponse, error) {
	panic("implement me")
}

func (c *Client) Negotiate(ctx context.Context, message *protocolv1.NegotiateMessage) (*emptypb.Empty, error) {

	if err := c.initConn(ctx); err != nil {
		return nil, err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if err := message.Formatter(c.conn, 0); err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (c *Client) Connect(ctx context.Context, message *protocolv1.ConnectMessage) (resp *protocolv1.ConnectMessageAck, err error) {

	if err := c.initConn(ctx); err != nil {
		return nil, err
	}

	if err := c.request(ctx, message, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Disconnect(ctx context.Context, message *protocolv1.DisconnectMessage) (resp *emptypb.Empty, err error) {

	if err := c.initConn(ctx); err != nil {
		return nil, err
	}

	if err := c.request(ctx, message, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) EndpointOpen(ctx context.Context, open *protocolv1.EndpointOpen) (resp *protocolv1.EndpointOpenAck, err error) {

	if err := c.lazyInit(ctx); err != nil {
		return nil, err
	}

	if err := c.request(ctx, open, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) EndpointClose(ctx context.Context, endpointClose *protocolv1.EndpointClose) (resp *emptypb.Empty, err error) {

	if err := c.initConn(ctx); err != nil {
		return nil, err
	}

	if err := c.request(ctx, endpointClose, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) EndpointMessage(ctx context.Context, message *protocolv1.EndpointMessage) (resp *protocolv1.EndpointMessage, err error) {

	if err := c.lazyInit(ctx); err != nil {
		return nil, err
	}

	if err := c.request(ctx, message, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) lazyInit(ctx context.Context) (err error) {
	c.onceInit.Do(func() {
		err = c.init(ctx)
	})
	return
}

func (c *Client) inited() bool {
	return atomic.LoadUint32(&c._inited) == 1
}

func (c *Client) connected() bool {
	return atomic.LoadUint32(&c._connected) == 1
}

func (c *Client) init(ctx context.Context) (err error) {

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if atomic.LoadUint32(&c._inited) == 1 {
		return
	}

	if atomic.LoadUint32(&c._connected) == 0 {
		err = c.initConn(ctx)
		if err != nil {
			return
		}
	}

	err = c.NegotiateMessage.Formatter(c.conn, 0)
	if err != nil {
		return
	}

	err = c.request(ctx, c.ConnectMessage, &protocolv1.ConnectMessageAck{})
	if err != nil {
		return
	}

	atomic.StoreUint32(&c._inited, 1)

	return err

}
func (c *Client) lazyConn(ctx context.Context) (err error) {

	if atomic.LoadUint32(&c._connected) == 0 {
		err = c.initConn(ctx)
		if err != nil {
			return
		}
	}

	return nil
}

func (c *Client) initConn(ctx context.Context) (err error) {

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if len(c.host) == 0 {
		return fmt.Errorf("host is not set")
	}

	if c.conn != nil && !c.closed() {
		return nil
	}

	err = c.populateConn(ctx)
	if err != nil {
		return err
	}

	go c.reaper(c.IdleCheckFrequency)

	atomic.StoreUint32(&c._connected, 1)

	return nil
}

func (p *Client) reaper(frequency time.Duration) {

	if p.IdleTimeout == 0 {
		return
	}

	ticker := time.NewTicker(frequency)
	defer ticker.Stop()

	for range ticker.C {
		if p.closed() {
			break
		}

		now := time.Now()
		if now.Sub(p.UsedAt()) >= p.IdleTimeout {
			err := p.Close()
			if err != nil {
				log.Println("reaper err", err)
				return
			}
		}
	}
}

func (c *Client) populateConn(_ context.Context) (err error) {

	conn, err := net.Dial("tcp", c.host)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}

func (c *Client) send(packetMessage protocolv1.PacketMessageFormatter) error {
	packet, err := protocolv1.NewPacket(packetMessage)
	if err != nil {
		return err
	}
	if _, err := packet.WriteTo(c.conn); err != nil {
		return err
	}
	return nil
}

func (c *Client) request(ctx context.Context, req protocolv1.PacketMessageFormatter, resp interface{}) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.send(req)
	atomic.AddUint32(&c.stats.Send, 1)
	if err != nil {
		atomic.AddUint32(&c.stats.Wrong, 1)
		return err
	}

	switch typed := resp.(type) {
	case protocolv1.PacketMessageParser:

		packet, err := c.recv(ctx)
		if err != nil {
			return err
		}

		atomic.AddUint32(&c.stats.Recv, 1)
		err = packet.Unpack(typed)
		if err != nil {
			atomic.AddUint32(&c.stats.Wrong, 1)
			return err
		}

	default:
	}

	return nil
}

func (c *Client) recv(ctx context.Context) (*protocolv1.Packet, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	err := c.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return nil, err
	}

	return protocolv1.NewPacket(c.conn)
}

func (c *Client) UsedAt() time.Time {
	unix := atomic.LoadUint32(&c.usedAt)
	return time.Unix(int64(unix), 0)
}

func (c *Client) SetUsedAt(tm time.Time) {
	atomic.StoreUint32(&c.usedAt, uint32(tm.Unix()))
}

func (c *Client) Close() error {

	if !atomic.CompareAndSwapUint32(&c._closed, 0, 1) {
		return nil
	}

	ctx := context.Background()
	var err error
	c.endpoints.Range(func(key, value interface{}) bool {

		err = c.request(ctx, &protocolv1.EndpointClose{EndpointId: key.(int32)}, nil)
		if err != nil {
			return false
		}

		return true
	})

	if atomic.CompareAndSwapUint32(&c._inited, 0, 1) {

		err = c.request(ctx, &protocolv1.DisconnectMessage{}, nil)
		if err != nil {
			return err
		}

	}

	if c.closed() {
		return nil
	}

	return c.conn.Close()
}

func (c *Client) closed() bool {

	if atomic.LoadUint32(&c._closed) == 1 {
		return true
	}
	_ = c.conn.SetReadDeadline(time.Now())
	_, err := c.conn.Read(make([]byte, 0))
	var zero time.Time
	_ = c.conn.SetReadDeadline(zero)

	if err == nil {
		return false
	}

	netErr, _ := err.(net.Error)
	if err != io.EOF && !netErr.Timeout() {
		atomic.StoreUint32(&c._closed, 1)
		return true
	}
	return false
}

// func (c *Client) Open(version string) (endpoint clientv1.EndpointServiceImpl, err error) {

// EndpointOpenAck, err := c.client.EndpointOpen(&protocolv1.EndpointOpen{
// 	Service: "v8.service.Admin.Cluster",
// 	Version: version,
// })
//
// if err != nil {
// 	if version := c.client.DetectSupportedVersion(err); len(version) == 0 {
// 		return nil, err
// 	}
// 	if EndpointOpenAck, err = c.client.EndpointOpen(&protocolv1.EndpointOpen{
// 		Service: "v8.service.Admin.Cluster",
// 		Version: version,
// 	}); err != nil {
// 		return nil, err
// 	}
// }
//
// end, err := c.client.NewEndpoint(EndpointOpenAck)
// if err != nil {
// 	return nil, err
// }
//
// return clientv1.NewEndpointService(c.client, end), nil
// }
