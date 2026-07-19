package storage

import "sync"

var globalTokenStore TokenStore
var createGlobalTokenStoreOnce sync.Once

func GetTokenStore() *TokenStore {
	createGlobalTokenStoreOnce.Do(func() {
		globalTokenStore = TokenStore{}
	})
	return &globalTokenStore
}

type TokenStore struct {
	mu sync.RWMutex
	token string
}

func NewTokenStore() *TokenStore{
	return &TokenStore{}
}

func (ts *TokenStore) Set(token string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.token = token
}

func (ts *TokenStore) Get() string {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	return ts.token
}
