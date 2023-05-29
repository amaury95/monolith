package connection

import (
	"context"
	"sync"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type IConfig interface {
	ArangoEndpoints(context.Context) ([]string, error)
	ArangoDatabase(context.Context) (string, error)
	ArangoRootPassword(context.Context) (string, error)
}

type arangoClient struct {
	db  driver.Database
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

	dbName, err := c.cnf.ArangoDatabase(ctx)
	if err != nil {
		return err
	}

	found, err := client.DatabaseExists(ctx, dbName)
	if err != nil {
		return err
	}

	if !found {
		db, err := client.CreateDatabase(ctx, dbName, &driver.CreateDatabaseOptions{})
		if err != nil {
			return err
		}
		return c.setDb(db)
	}

	db, err := client.Database(ctx, dbName)
	if err != nil {
		return err
	}
	return c.setDb(db)
}

func (c *arangoClient) setDb(db driver.Database) error {
	c.m.Lock()
	defer c.m.Unlock()

	c.db = db
	return nil
}

func (c *arangoClient) GetConnection(ctx context.Context) (driver.Database, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	if _, err := c.db.Info(ctx); err != nil {
		return nil, false // connection is not healthy
	}

	return c.db, true
}
