package types

const (
	// ModuleName is the name of the module
	ModuleName = "telepathy"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

const (
	ThoughtPrefix      = "thought-value-"
	ThoughtCountPrefix = "thought-count-"
)

const (
	CommentPrefix      = "comment-value-"
	CommentCountPrefix = "comment-count-"
)

const (
	UserPrefix      = "user-value-"
	UserCountPrefix = "user-count-"
)
