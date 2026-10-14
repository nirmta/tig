package UserEntity

type UserDraft struct {
	Alias        string
	Name         string
	Email        string
	Host         string
	IdentityFile string
}

func Constructor(userDraft UserDraft) *User {
	return &User{
		alias:        userDraft.Alias,
		name:         userDraft.Name,
		email:        userDraft.Email,
		host:         userDraft.Host,
		identityFile: userDraft.IdentityFile,
		inGlobal:     false,
	}
}
