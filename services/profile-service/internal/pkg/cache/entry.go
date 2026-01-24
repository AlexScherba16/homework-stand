package cache

import (
	"time"

	"github.com/samber/lo"
)

// Entries ...
type Entries[TValue Marshaller, TValuePtr UnMarshaller] []Entry[TValue, TValuePtr]

// Entry запись в кеше
type Entry[TValue Marshaller, TValuePtr UnMarshaller] struct {
	Key        string
	Value      TValue
	Expiration time.Duration
}

// From создает новый объект
func From[TValue Marshaller, TValuePtr UnMarshaller](key string, rawValue []byte) (Entry[TValue, TValuePtr], error) {
	var result TValue
	if err := unmarshal[TValuePtr](rawValue, &result); err != nil {
		return Entry[TValue, TValuePtr]{}, err
	}

	return Entry[TValue, TValuePtr]{
		Key:   key,
		Value: result,
	}, nil
}

func (e *Entry[TValue, TValuePtr]) TTL() time.Duration {
	return e.Expiration
}

// Marshall конвертирует значение в байты
func (e *Entry[TValue, TValuePtr]) marshall() []byte {
	bytes, _ := marshall[TValue](e.Value)
	return bytes
}

// Map возвращает ключи в виде map
func (e Entries[TValue, TValuePtr]) Map() map[string]Entry[TValue, TValuePtr] {
	return lo.SliceToMap(e, func(entry Entry[TValue, TValuePtr]) (string, Entry[TValue, TValuePtr]) {
		return entry.Key, entry
	})
}
