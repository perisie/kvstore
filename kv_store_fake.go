package kvstore

import (
	"sort"
	"strings"
)

type Kv_store_fake struct {
	m map[string]map[string]string
}

func Kv_store_fake_new() *Kv_store_fake {
	return &Kv_store_fake{
		m: map[string]map[string]string{},
	}
}

func (k *Kv_store_fake) Create(key *Key, value string) (*Key_value, error) {
	if _, ok := k.m[key.Get_partition_key()]; !ok {
		k.m[key.Get_partition_key()] = map[string]string{}
	}
	k.m[key.Get_partition_key()][key.Get_sort_key()] = value
	kv := Key_value_new(key, k.m[key.Get_partition_key()][key.Get_sort_key()])
	return kv, nil
}

func (k *Kv_store_fake) Create_not_exist(key *Key, value string) (*Key_value, error) {
	if _, ok := k.m[key.Get_partition_key()]; ok {
		return nil, Err_already_exist
	}
	if _, ok := k.m[key.Get_partition_key()][key.Get_sort_key()]; ok {
		return nil, nil
	}
	return k.Create(key, value)
}

func (k *Kv_store_fake) Get(key *Key) ([]*Key_value, error) {
	if _, ok := k.m[key.Get_partition_key()]; !ok {
		return []*Key_value{}, nil
	}
	var kvs []*Key_value
	for _, v := range k.m[key.Get_partition_key()] {
		kvs = append(kvs, Key_value_new(key, v))
	}
	sort.Slice(kvs, func(a, b int) bool {
		va := kvs[a]
		vb := kvs[b]
		return strings.Compare(va.key.Get_sort_key(), vb.key.Get_sort_key()) < 0
	})
	return kvs, nil
}
