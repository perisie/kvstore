package kvstore

type Kv_store interface {
	Create(key string, value string) (*Key_value, error)
	Create_not_exist(key string, value string) (*Key_value, error)
	Get(key string) (*Key_value, error)
	Get_many(keys []string) ([]*Key_value, error)
}
