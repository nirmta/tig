package constants

const (
	BITBUCKET = "bitbucket.org"
	GITHUB    = "github.com"
	GITLAB    = "gitlab.com"
)

var RepositoryHosts = map[string]string{
	"b": BITBUCKET,
	"g": GITHUB,
	"l": GITLAB,
}
