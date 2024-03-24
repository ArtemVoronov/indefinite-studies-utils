package entities

const (
	USER_ROLE_OWNER    string = "OWNER"
	USER_ROLE_RESIDENT string = "RESIDENT"
	USER_ROLE_GI       string = "GI"
)

func GetPossibleUserRoles() []string {
	return []string{USER_ROLE_OWNER, USER_ROLE_RESIDENT, USER_ROLE_GI}
}
