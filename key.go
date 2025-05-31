package kvstore

import "strings"

type Key struct {
	partition string
	sort      string
}

func (k *Key) Get_partition_key() string {
	return strings.TrimSpace(k.partition)
}

func (k *Key) Get_sort_key() string {
	return strings.TrimSpace(k.sort)
}

func (k *Key) Get_key() string {
	if k.partition == "" {
		return ""
	}
	if k.Get_sort_key() == "" {
		return k.Get_partition_key()
	}
	return k.Get_partition_key() + sort_split + k.Get_sort_key()
}

func Key_new(partition_key string, sort_key string) *Key {
	return &Key{
		partition: partition_key,
		sort:      sort_key,
	}
}

func Key_from(s string) *Key {
	chunks := strings.Split(s, sort_split)
	if len(chunks) > 1 {
		return &Key{
			partition: chunks[0],
			sort:      chunks[1],
		}
	} else {
		return &Key{
			partition: chunks[0],
			sort:      "",
		}
	}
}
