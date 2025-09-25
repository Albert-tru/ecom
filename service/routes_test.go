package user

import (
	"testing"

	"github.com/Albert-tru/ecom/types"
)

// MockUserStore is a mock implementation of UserStore for testing
type MockUserStore struct{}

func (m *MockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{ID: 1, Username: "test", Email: email}, nil
}

func (m *MockUserStore) GetUserByID(id int) (*types.User, error) {
	return &types.User{ID: id, Username: "test", Email: "test@example.com"}, nil
}

func (m *MockUserStore) CreateUser(user types.User) error {
	return nil
}

func TestNewHandler(t *testing.T) {
	mockStore := &MockUserStore{}
	handler := NewHandler(mockStore)
	
	if handler == nil {
		t.Fatal("NewHandler returned nil")
	}
	
	if handler.store == nil {
		t.Fatal("Handler store is nil")
	}
}

func TestHandlerWithMockStore(t *testing.T) {
	mockStore := &MockUserStore{}
	handler := NewHandler(mockStore)
	
	// Test that we can access the store from the handler
	user, err := handler.store.GetUserByID(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	if user.ID != 1 {
		t.Fatalf("Expected user ID 1, got %d", user.ID)
	}
}