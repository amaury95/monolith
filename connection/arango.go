package connection

import (
	"context"
	"sync"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type IConfig interface {
	ArangoEndpoints(context.Context) ([]string, error)
	ArangoRootPassword(context.Context) (string, error)
}

type arangoClient struct {
	c   *driver.Client
	cnf IConfig

	m sync.RWMutex
}

func NewArangoClient(cnf IConfig) *arangoClient {
	return &arangoClient{m: sync.RWMutex{}}
}

func (c *arangoClient) Init(ctx context.Context) error {
	endpoints, err := c.cnf.ArangoEndpoints(ctx)
	if err != nil {
		return err
	}

	conn, err := http.NewConnection(http.ConnectionConfig{Endpoints: endpoints})
	if err != nil {
		return err
	}

	password, err := c.cnf.ArangoRootPassword(ctx)
	if err != nil {
		return err
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", password),
	})
	if err != nil {
		return err
	}
	
	c.c = &client
	return nil
}

func (c *arangoClient) GetConnection(ctx context.Context) (*driver.Client, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	return c.c, true
}
