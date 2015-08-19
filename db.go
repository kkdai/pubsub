package pubsub

type db struct {
	id  int
	val interface{}
}

func (b *db) getVal() interface{} {
	return b.val
}
