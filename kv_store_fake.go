package kvstore

type Kv_store_fake struct {
	m map[string]string
}

func Kv_store_fake_new() *Kv_store_fake {
	return &Kv_store_fake{
		m: map[string]string{},
	}
}

func (k *Kv_store_fake) Create_not_exist(key string, value string) (*Key_value, error) {
	kv, _ := k.Get(key)
	if !kv.Exist() {
		return k.Create(key, value)
	} else {
		return kv, nil
	}
}

func (k *Kv_store_fake) Create(key string, value string) (*Key_value, error) {
	k.m[key] = value
	kv := Key_value_new(key, k.m[key])
	return kv, nil
}

func (k *Kv_store_fake) Get(key string) (*Key_value, error) {
	if _, ok := k.m[key]; !ok {
		return Key_value_new("", ""), nil
	}
	kv := Key_value_new(key, k.m[key])
	return kv, nil
}

func (k *Kv_store_fake) Get_many(keys []string) ([]*Key_value, error) {
	var kvs []*Key_value
	for _, key := range keys {
		kv, _ := k.Get(key)
		kvs = append(kvs, kv)
	}
	return kvs, nil
}
