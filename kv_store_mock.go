package kvstore

import "github.com/stretchr/testify/mock"

type Kv_store_mock struct {
	mock.Mock
}

func (k *Kv_store_mock) Create(key string, value string) (*Key_value, error) {
	args := k.Called(key, value)
	return args.Get(0).(*Key_value), args.Error(1)
}

func (k *Kv_store_mock) Create_not_exist(key string, value string) (*Key_value, error) {
	args := k.Called(key, value)
	return args.Get(0).(*Key_value), args.Error(1)
}

func (k *Kv_store_mock) Get(key string) (*Key_value, error) {
	args := k.Called(key)
	return args.Get(0).(*Key_value), args.Error(1)
}

func (k *Kv_store_mock) Get_many(keys []string) ([]*Key_value, error) {
	args := k.Called(keys)
	return args.Get(0).([]*Key_value), args.Error(1)
}
