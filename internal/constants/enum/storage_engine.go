package enum

type StorageEngine int

const (
	Local StorageEngine = iota
	Cos
	Oss
)
