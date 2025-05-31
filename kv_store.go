package kvstore

type Kv_store interface {
	Create(key *Key, value string) (*Key_value, error)
	Create_not_exist(key *Key, value string) (*Key_value, error)
	Get(key *Key) ([]*Key_value, error)
}
