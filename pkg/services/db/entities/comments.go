package entities

const (
	COMMENT_STATE_NEW           string = "NEW"
	COMMENT_STATE_ON_MODERATION string = "ON_MODERATION"
	COMMENT_STATE_PUBLISHED     string = "PUBLISHED"
	COMMENT_STATE_BLOCKED       string = "BLOCKED"
	COMMENT_STATE_DELETED       string = "DELETED"
)

func GetPossibleCommentStates() []string {
	return []string{COMMENT_STATE_NEW, COMMENT_STATE_ON_MODERATION, COMMENT_STATE_PUBLISHED, COMMENT_STATE_BLOCKED, COMMENT_STATE_DELETED}
}