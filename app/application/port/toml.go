package port

type TomlService interface {
	AddExtension(fileName string) string
	CreateFile(data interface{}, path string, fileName string) error
	ReadFile(path string, fileName string) ([]byte, error)
	ParseFile(data []byte, v interface{}) error
}
