package handle

import (
	"encoding/json"

	"nakama-golang/fantasy"
	"nakama-golang/model"
	"nakama-golang/model/event"
	"nakama-golang/util"
)

func worldEvent(c *fantasy.Claude) {
	switch c.Event() {
	case event.EventMatchJoin:
		info := c.Evt.Properties
		aid := util.ToInt64(info["aid"])
		matchGroup.AddPlayer(aid, info["user_id"])
	case event.EventMatchReady:
		info := c.Evt.Properties
		aid := util.ToInt64(info["aid"])
		matchId := info["match_id"]
		matchGroup.ReadyMatch(aid, matchId, info["user_id"], info["session_id"])
	case event.EventGameReady:
		info := c.Evt.Properties
		msg := &model.GameMsg{
			UserId:    info["user_id"],
			SessionId: info["session_id"],
			MatchId:   info["match_id"],
			Data:      nil,
		}
		gameGroup.Run(msg)
	case event.EventGameRun:
		info := c.Evt.Properties
		data := &model.GamePlayFrame{}
		json.Unmarshal([]byte(info["data"]), data)
		msg := &model.GameMsg{
			UserId:    info["user_id"],
			SessionId: info["session_id"],
			MatchId:   info["match_id"],
			Data:      data,
		}
		gameGroup.Run(msg)
	}
}
