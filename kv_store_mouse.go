package kvstore

import (
	"github.com/perisie/mouse"
)

type Kv_store_mouse struct {
	m mouse.Mouse
}

func (k *Kv_store_mouse) Create(key *Key, value string) (*Key_value, error) {
	err := k.m.Put(key.Get_key(), []byte(value))
	if err != nil {
		return nil, err
	}
	return Key_value_new(key, value), nil
}

func (k *Kv_store_mouse) Create_not_exist(key *Key, value string) (*Key_value, error) {
	kvs, err := k.Get(key)
	if err != nil {
		return nil, err
	}
	insert := true
	for _, kv := range kvs {
		if kv.Get_key().Get_key() == key.Get_key() {
			insert = false
		}
	}
	if !insert {
		return nil, Err_already_exist
	}
	return k.Create(key, value)
}

func (k *Kv_store_mouse) Get(key *Key) ([]*Key_value, error) {
	keys, err := k.m.Get_keys(key.Get_key())
	if err != nil {
		return nil, err
	}
	kvs := make([]*Key_value, 0)
	for _, key_g := range keys {
		bytes, err_get := k.m.Get(key_g)
		if err_get != nil {
			return nil, err_get
		}
		if key.Get_sort_key() == "" {
			key_s := Key_from(key_g)
			kvs = append(kvs, Key_value_new(key_s, string(bytes)))
		} else {
			if key_g == key.Get_key() {
				kvs = append(kvs, Key_value_new(key, string(bytes)))
			}
		}
	}
	return kvs, nil
}

func Kv_store_mouse_new(data_dir_path string) *Kv_store_mouse {
	m := mouse.Mouse_fs_new(data_dir_path)
	return &Kv_store_mouse{
		m: m,
	}
}
