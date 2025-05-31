package kvstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_create_get(t *testing.T) {
	var kv_database Kv_store = Kv_store_fake_new()

	_, _ = kv_database.Create(&Key{partition: "key1"}, "value1")

	key_values, _ := kv_database.Get(&Key{partition: "key1"})
	assert.Equal(t, 1, len(key_values))
	assert.Equal(t, "key1", key_values[0].Get_key().Get_partition_key())
	assert.Equal(t, "", key_values[0].Get_key().Get_sort_key())
	assert.Equal(t, "value1", key_values[0].Get_value())
}

func Test_create_not_exist(t *testing.T) {
	var kv_database Kv_store = Kv_store_fake_new()

	kv, _ := kv_database.Create_not_exist(&Key{partition: "Key"}, "value")
	assert.Equal(t, "value", kv.Get_value())

	_, _ = kv_database.Create_not_exist(&Key{partition: "Key"}, "?")

	kvs, _ := kv_database.Get(&Key{partition: "Key"})
	assert.Equal(t, 1, len(kvs))
	assert.Equal(t, "value", kvs[0].Get_value())
}
