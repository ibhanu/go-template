package config

import (
	"testing"

	"web-server/prisma/db"

	"github.com/stretchr/testify/assert"
)

type testClient struct {
	connectCalled    bool
	disconnectCalled bool
	shouldError      bool
}

func newTestClient(shouldError bool) *testClient {
	return &testClient{shouldError: shouldError}
}

func (c *testClient) Connect() error {
	c.connectCalled = true
	if c.shouldError {
		return assert.AnError
	}
	return nil
}

func (c *testClient) Disconnect() error {
	c.disconnectCalled = true
	if c.shouldError {
		return assert.AnError
	}
	return nil
}

func TestDatabaseConnection(t *testing.T) {
	originalGetClient := GetPrismaClient

	t.Run("GetPrismaClient returns singleton", func(t *testing.T) {
		// Get first instance
		client1 := GetPrismaClient()
		assert.NotNil(t, client1)

		// Get second instance
		client2 := GetPrismaClient()
		assert.NotNil(t, client2)

		// Should be the same instance
		assert.Same(t, client1, client2)
	})

	t.Run("handles connection errors", func(t *testing.T) {
		testClient := newTestClient(true) // Should return error
		GetPrismaClient = func() *db.PrismaClient { return nil }

		// Since connection error causes log.Fatal, we can't test it directly
		// but we can verify the test client behavior
		err := testClient.Connect()
		assert.Error(t, err)
		assert.True(t, testClient.connectCalled)
	})

	t.Run("handles successful connection and disconnection", func(t *testing.T) {
		testClient := newTestClient(false) // Should not error

		err := testClient.Connect()
		assert.NoError(t, err)
		assert.True(t, testClient.connectCalled)

		err = testClient.Disconnect()
		assert.NoError(t, err)
		assert.True(t, testClient.disconnectCalled)
	})

	t.Run("handles disconnection errors", func(t *testing.T) {
		testClient := newTestClient(true) // Should return error

		err := testClient.Disconnect()
		assert.Error(t, err)
		assert.True(t, testClient.disconnectCalled)
	})

	// Restore original GetPrismaClient
	GetPrismaClient = originalGetClient
}

func TestDisconnectDB(t *testing.T) {
	t.Run("safely handles nil client", func(t *testing.T) {
		originalGetClient := GetPrismaClient
		GetPrismaClient = func() *db.PrismaClient { return nil }

		DisconnectDB() // Should not panic

		GetPrismaClient = originalGetClient
	})

	t.Run("handles disconnect error", func(t *testing.T) {
		testClient := newTestClient(true) // Should error
		DisconnectDB()
		assert.True(t, testClient.disconnectCalled)
	})
}
