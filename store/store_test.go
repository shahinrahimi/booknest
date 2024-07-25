package store

import "testing"

// Test NewSqliteStore
func TestNewSqliteStore(t *testing.T) {
	logger := setupTestLogger()
	store := NewSqliteStore(logger)

	defer store.Close()

	if store.db == nil {
		t.Fatal("Expected db to be initialized, got nil")
	}
}
