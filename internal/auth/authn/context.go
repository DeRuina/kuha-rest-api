package authn

import "context"

type ctxKey string

const (
	clientKey ctxKey = "client_name"
	rolesKey  ctxKey = "roles"
)

func WithClientMetadata(ctx context.Context, name string, roles []string) context.Context {
	ctx = context.WithValue(ctx, clientKey, name)
	ctx = context.WithValue(ctx, rolesKey, roles)
	return ctx
}

func GetClientName(ctx context.Context) string {
	if val, ok := ctx.Value(clientKey).(string); ok {
		return val
	}
	return ""
}

func GetClientRoles(ctx context.Context) []string {
	if val, ok := ctx.Value(rolesKey).([]string); ok {
		return val
	}
	return nil
}
