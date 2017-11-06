package api

import (
	"txmeeting/lib"

	"github.com/ivpusic/neo"
)

var meetingAPI *neo.Region

func loadmeetingAPI() {
	meetingAPI = lib.App.Region().Prefix("/api/meeting")

	meetingAPI.Get("/all", func(ctx *neo.Ctx) (int, error) {
		return 200, ctx.Res.Text("meeting say hello")
	})

}
