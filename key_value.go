package kvstore

type Key_value struct {
	Key   string
	Value string
}

func (k *Key_value) Exist() bool {
	return k.Key != ""
}

func Key_value_new(key string, value string) *Key_value {
	return &Key_value{
		Key:   key,
		Value: value,
	}
}
