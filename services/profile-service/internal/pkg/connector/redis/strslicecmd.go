package redis

type StrSliceCmd struct {
	val []interface{}
	err error
}

// Val Получаем значения
func (c *StrSliceCmd) Val() []interface{} {
	return c.val
}

// Err Получаем ошибку
func (c *StrSliceCmd) Err() error {
	return c.err
}

// SetErr Для установки ошибки вручную (как в хуках)
func (c *StrSliceCmd) SetErr(err error) {
	c.err = err
}
