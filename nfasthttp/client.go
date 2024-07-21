package nfasthttp

import (
	"fmt"

	"github.com/neosy/gofw/nbasic"
	"github.com/valyala/fasthttp"
)

type Client struct {
	Address      string
	Port         int
	fasthttpReq  *fasthttp.Request
	fasthttpResp *fasthttp.Response
}

func CreateClient(addr string, port int) *Client {
	client := &Client{}

	client.Init(addr, port)

	return client
}

// Генерация URI
func (c *Client) CreateURI(cmd string) string {
	return fmt.Sprintf("%s:%v%s", c.Address, c.Port, cmd)
}

func (c *Client) Response() *fasthttp.Response {
	return c.fasthttpResp
}

func (c *Client) Init(addr string, port int) {
	c.Address = addr
	c.Port = port
}

func (c *Client) Release() {
	defer fasthttp.ReleaseRequest(c.fasthttpReq)
	defer fasthttp.ReleaseResponse(c.fasthttpResp)
}

// Отправка POST запроса
func (c *Client) SendRequest(reqUri string, method string, dataReq interface{}) error {
	var err error

	c.fasthttpReq = fasthttp.AcquireRequest()
	c.fasthttpReq.Header.SetMethod(method)
	c.fasthttpReq.SetRequestURI(reqUri)
	c.fasthttpResp = fasthttp.AcquireResponse()

	dataReqJSON, err := nbasic.StructToJSON(dataReq)
	if err != nil {
		return err
	}

	req := c.fasthttpReq
	resp := c.fasthttpResp

	req.SetBody(dataReqJSON)

	if err = fasthttp.Do(req, resp); err != nil {
		return err
	}

	return err
}
