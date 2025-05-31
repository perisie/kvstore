package kvstore

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_mouse_create_get(t *testing.T) {
	test_before(t)

	var kv_database Kv_store = Kv_store_mouse_new("data")

	_, _ = kv_database.Create(Key_new("key", ""), "value")

	key_values, _ := kv_database.Get(&Key{partition: "key"})
	require.Equal(t, 1, len(key_values))
	require.Equal(t, "key", key_values[0].Get_key().Get_key())
	require.Equal(t, "value", key_values[0].Get_value())
}

func Test_mouse_create_get_sort_key(t *testing.T) {
	test_before(t)

	var kv_database Kv_store = Kv_store_mouse_new("data")

	_, _ = kv_database.Create(Key_new("u/key", "2025-06-16"), "value_2")
	_, _ = kv_database.Create(Key_new("u/key", "2025-06-15"), "value_1")

	key_values, _ := kv_database.Get(Key_new("u/key", ""))
	require.Equal(t, 2, len(key_values))

	require.Equal(t, "u/key〰2025-06-15", key_values[0].Get_key().Get_key())
	require.Equal(t, "u/key〰2025-06-16", key_values[1].Get_key().Get_key())

	require.Equal(t, "value_1", key_values[0].Get_value())
	require.Equal(t, "value_2", key_values[1].Get_value())
}

func Test_mouse_create_not_exist(t *testing.T) {
	test_before(t)

	var kv_database Kv_store = Kv_store_mouse_new("data")

	kv, _ := kv_database.Create_not_exist(Key_new("key", ""), "value")
	require.Equal(t, "value", kv.Get_value())

	_, _ = kv_database.Create_not_exist(&Key{partition: "key"}, "?")

	kvs, _ := kv_database.Get(&Key{partition: "key"})
	require.Equal(t, 1, len(kvs))
	require.NotEqual(t, "?", kvs[0].Get_value()[0])
}

func test_before(t *testing.T) {
	err := os.RemoveAll("data")
	require.Nil(t, err)
}
