package api

import (
	"strconv"
	"txmeeting/lib"
	"txmeeting/models"

	"github.com/ivpusic/neo"
)

var userAPI *neo.Region

func loadUserAPI() {

	userAPI = lib.App.Region().Prefix("/api/user")

	userAPI.Get("/all", func(ctx *neo.Ctx) (int, error) { //  直接写查询
		page, _ := strconv.Atoi(ctx.Req.FormValue("page"))
		pageSize, _ := strconv.Atoi(ctx.Req.FormValue("page_size"))

		result := models.NewResult(page, pageSize)
		users := make([]models.CcUser, result.PageSize)
		models.Orm.Offset(result.OffSet()).Limit(result.PageSize).Find(&users).Count(&result.AllCount)
		result.Data = users
		return 200, ctx.Res.Json(result)
	})

	userAPI.Get("/all_dao", func(ctx *neo.Ctx) (int, error) { //  写DAO
		page, _ := strconv.Atoi(ctx.Req.FormValue("page"))
		pageSize, _ := strconv.Atoi(ctx.Req.FormValue("page_size"))
		result := models.GetAllUser(page, pageSize)
		return 200, ctx.Res.Json(result)
	})

}
