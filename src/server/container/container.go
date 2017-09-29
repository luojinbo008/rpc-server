package container

type Container interface {
	GetDb(msg interface{})
}