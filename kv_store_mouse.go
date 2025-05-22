package kvstore

import (
	"github.com/perisie/mouse"
)

type Kv_store_mouse struct {
	m mouse.Mouse
}

func Kv_store_mouse_new(data_dir_path string) *Kv_store_mouse {
	m := mouse.Mouse_fs_new(data_dir_path)
	return &Kv_store_mouse{
		m: m,
	}
}

func (k *Kv_store_mouse) Create_not_exist(key string, value string) (*Key_value, error) {
	kv, err := k.Get(key)
	if err != nil {
		return nil, err
	}
	if kv == nil || !kv.Exist() {
		return k.Create(key, value)
	} else {
		return kv, nil
	}
}

func (k *Kv_store_mouse) Create(key string, value string) (*Key_value, error) {
	err := k.m.Put(key, []byte(value))
	if err != nil {
		return nil, err
	}
	kv := Key_value_new(key, value)
	return kv, nil
}

func (k *Kv_store_mouse) Get(key string) (*Key_value, error) {
	bytes, err := k.m.Get(key)
	if err != nil {
		return nil, err
	}
	kv := Key_value_new(key, string(bytes))
	return kv, nil
}

func (k *Kv_store_mouse) Get_many(keys []string) ([]*Key_value, error) {
	var kvs []*Key_value
	for _, key := range keys {
		kv, _ := k.Get(key)
		kvs = append(kvs, kv)
	}
	return kvs, nil
}
