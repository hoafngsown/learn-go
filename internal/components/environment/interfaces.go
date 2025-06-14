package environment

type SPACE string

const (
	DATABASE        SPACE = "database"
	SERVER          SPACE = "server"
	UPLOAD_PROVIDER SPACE = "upload_provider"
)

type env interface {
	GetEnv(space SPACE, key string) string
}
