package interfaces

type SPACE string

const (
	DATABASE SPACE = "database"
	SERVER   SPACE = "server"
)

type Environment interface {
	GetEnv(space SPACE, key string) string
}
