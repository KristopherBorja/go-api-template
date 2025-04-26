package logs

type contextKey string

const (
	ContextKeyRequestID contextKey = "request_id"
	ContextKeyRemoteIP  contextKey = "remote_ip"
	ContextKeyUserAgent contextKey = "user_agent"
	ContextKeyURI       contextKey = "uri"
	ContextKeyMethod    contextKey = "method"
	ContextKeyHost      contextKey = "host"
)
