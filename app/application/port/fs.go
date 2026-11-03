package port

type FsService interface {
	CreateDirectory(path string) error
}
