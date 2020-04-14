package service

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/api"
	"nakama-golang/model/event"
	"nakama-golang/protocol"
	"nakama-golang/util"

	"github.com/heroiclabs/nakama-common/runtime"
)

func (s *Service) Hello(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, req protocol.Request) (v interface{}, err error) {
	userId, err := util.GetUserId(ctx)
	if err != nil {
		return v, err
	}
	logger.Info("userId:%v", userId)
	s.Notify(ctx, logger, db, nk, userId, &protocol.NotifyHelloMsg{})
	return req, nil
}

func (s *Service) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, req *protocol.ReqMatchJoin) (v interface{}, err error) {
	userId, err := util.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	if err := nk.Event(ctx, &api.Event{
		Name: event.EventMatchJoin.String(),
		Properties: map[string]string{
			"topic":   req.Topic,
			"aid":     util.ToString(req.Aid),
			"user_id": userId,
		},
		External: false,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Service) MatchReady(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, req *protocol.ReqMatchReady) (v interface{}, err error) {
	userId, err := util.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	sessionId, err := util.GetSessionId(ctx)
	if err != nil {
		return nil, err
	}
	if err := nk.Event(ctx, &api.Event{
		Name: event.EventMatchReady.String(),
		Properties: map[string]string{
			"match_id":   req.MatchId,
			"aid":        util.ToString(req.MatchId),
			"user_id":    userId,
			"session_id": sessionId,
		},
		External: false,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Service) GameTick(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, req *protocol.ReqGameTick) (v interface{}, err error) {
	jstr, _ := json.Marshal(req.Frame)
	userId, err := util.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	sessionId, err := util.GetSessionId(ctx)
	if err != nil {
		return nil, err
	}
	matchId, err:=util.GetMatchId(ctx)
	if err!=nil{
		return nil,err
	}
	if err := nk.Event(ctx, &api.Event{
		Name: event.EventGameRun.String(),
		Properties: map[string]string{
			"data":       string(jstr),
			"user_id":    userId,
			"session_id": sessionId,
			"match_id":   matchId,
		},
		External: false,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Service) GameReady(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, req *protocol.ReqGameReady) (v interface{}, err error) {
	userId, err := util.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	sessionId, err := util.GetSessionId(ctx)
	if err != nil {
		return nil, err
	}
	matchId, err:=util.GetMatchId(ctx)
	if err!=nil{
		return nil,err
	}
	if err := nk.Event(ctx, &api.Event{
		Name: event.EventGameReady.String(),
		Properties: map[string]string{
			"user_id":    userId,
			"session_id": sessionId,
			"match_id":   matchId,
		},
		External: false,
	}); err != nil {
		return nil, err
	}
	return nil, nil
}
