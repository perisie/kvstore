package kvstore

type Key_value struct {
	key   *Key
	value string
}

func (k *Key_value) Get_key() *Key {
	return k.key
}

func (k *Key_value) Get_value() string {
	return k.value
}

func (k *Key_value) Exist() bool {
	return k.key.Get_key() != ""
}

func Key_value_new(key *Key, value string) *Key_value {
	return &Key_value{
		key:   key,
		value: value,
	}
}
