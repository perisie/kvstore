package kvstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mouse_create_get(t *testing.T) {
	var kv_database Kv_store = Kv_store_mouse_new("data")
	_, _ = kv_database.Create("key1", "value1")
	key_value, _ := kv_database.Get("key1")
	assert.Equal(t, "key1", key_value.Key)
	assert.Equal(t, "value1", key_value.Value)
}

func Test_mouse_create_get_many(t *testing.T) {
	var kv_database Kv_store = Kv_store_mouse_new("data")
	keys := []string{"k1", "k2"}
	for _, key := range keys {
		_, _ = kv_database.Create(key, key+key)
	}
	kvs, _ := kv_database.Get_many(keys)
	assert.Equal(t, 2, len(kvs))
	assert.Equal(t, keys[0]+keys[0], kvs[0].Value)
	assert.Equal(t, keys[1]+keys[1], kvs[1].Value)
}

func Test_mouse_create_not_exist(t *testing.T) {
	var kv_database Kv_store = Kv_store_mouse_new("data")
	kv, _ := kv_database.Create_not_exist("key", "value")
	assert.Equal(t, "value", kv.Value)

	_, _ = kv_database.Create_not_exist("key", "?")
	kv, _ = kv_database.Get("key")
	assert.NotEqual(t, "?", kv.Value)
}
