package entities

const (
	POST_STATE_NEW           string = "NEW"
	POST_STATE_ON_MODERATION string = "ON_MODERATION"
	POST_STATE_PUBLISHED     string = "PUBLISHED"
	POST_STATE_BLOCKED       string = "BLOCKED"
	POST_STATE_DELETED       string = "DELETED"
)

func GetPossiblePostStates() []string {
	return []string{POST_STATE_NEW, POST_STATE_ON_MODERATION, POST_STATE_PUBLISHED, POST_STATE_BLOCKED, POST_STATE_DELETED}
}
