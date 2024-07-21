package ngrpc

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type Client struct {
	Address string
	Port    int

	grpcConn *grpc.ClientConn
}

// Создание клиента
func CreateClient(addr string, port int) *Client {
	client := &Client{}

	client.Init(addr, port)

	return client
}

// Генерация URL адреса
func (c *Client) CreateURL() string {
	return fmt.Sprintf("%s:%v", c.Address, c.Port)
}

// Инициализация параметров
func (c *Client) Init(addr string, port int) {
	c.Address = addr
	c.Port = port
}

// Соединение с сервисом gRPC
func (c *Client) Connect(opts ...grpc.DialOption) (err error) {
	c.grpcConn, err = grpc.NewClient(c.CreateURL(), opts...)
	if err != nil {
		log.Println(
			ErrClientConnect.Error(),
			err,
		)
	}

	return err
}

// Оснобождение ресурсов
func (c *Client) Release() {
	defer c.grpcConn.Close()
}

func (c *Client) GetConn() *grpc.ClientConn {
	return c.grpcConn
}
