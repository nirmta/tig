package constants

const (
	BITBUCKET = "ssh.bitbucket.org"
	GITHUB    = "github.com"
	GITLAB    = "gitlab.com"
)

var RepositoryHosts = map[string]string{
	"b": BITBUCKET,
	"g": GITHUB,
	"l": GITLAB,
}
