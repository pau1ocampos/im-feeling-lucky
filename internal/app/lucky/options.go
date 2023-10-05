package lucky

type FileSystem interface {
	ReadFile(path string) ([]byte, error)
}

type Options struct {
	Client      HttpCli
	FileSystem  FileSystem
	BaseUrl     string
	FromYear    int
	UserAgent   string
	ShouldStore bool
	StoreFile   string
}
