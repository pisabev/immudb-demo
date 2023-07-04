package model

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestApp(t *testing.T) {
	dsn := startImmudb(t)
	repo, err := NewMyLogImmudbRepo(context.Background(), dsn)
	assert.NoError(t, err)

	t.Run("Create myLog table", func(t *testing.T) {
		err = repo.CreateLogTable()
		assert.NoError(t, err)
	})

	t.Run("Insert into myLog table", func(t *testing.T) {
		err = repo.CreateLogTable()
		assert.NoError(t, err)

		err = repo.InsertLog(&Log{Data: "some data"})
		assert.NoError(t, err)
	})

	t.Run("Fetch data from myLog", func(t *testing.T) {
		err = repo.CreateLogTable()
		assert.NoError(t, err)

		err = repo.InsertLog(&Log{Data: "random data 1"})
		assert.NoError(t, err)
		err = repo.InsertLog(&Log{Data: "random data 2"})
		assert.NoError(t, err)

		res, err := repo.FetchLog(1)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, "random data 2", res[0].Data)
	})
}

func startImmudb(t *testing.T) string {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "codenotary/immudb",
		ExposedPorts: []string{"3322/tcp"},
		WaitingFor:   wait.ForLog("Web API server enabled"),
	}
	immudbC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		immudbC.Terminate(ctx) // nolint: errcheck
	})

	host, _ := immudbC.Host(ctx)
	port, _ := immudbC.MappedPort(ctx, "3322")

	return fmt.Sprintf("immudb://%s:%s@%s:%s/%s?sslmode=disable",
		"immudb", "immudb", host, port.Port(), "defaultdb")
}
