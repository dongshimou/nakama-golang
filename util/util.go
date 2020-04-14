package util

import (
	"context"
	"fmt"
	"strconv"

	"github.com/heroiclabs/nakama-common/runtime"
	"nakama-golang/model/code"
)

func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
func ToInt64(str string) int64 {
	v, _ := strconv.ParseInt(str, 10, 64)
	return v
}

func GetUserId(ctx context.Context) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		// User ID not found in the context.
		return "", code.UserIdNotFound
	}
	return userId, nil
}
func GetSessionId(ctx context.Context) (string, error) {
	sessionID, ok := ctx.Value(runtime.RUNTIME_CTX_SESSION_ID).(string)
	if !ok {
		// If session ID is not found, RPC was not called over a connected socket.
		return "", code.SessionIdNotFound
	}
	return sessionID, nil
}

func GetMatchId(ctx context.Context)(string,error){
	matchId, ok := ctx.Value(runtime.RUNTIME_CTX_MATCH_ID).(string)
	if !ok {
		return "", code.MatchIdNotFound
	}
	return matchId,nil
}