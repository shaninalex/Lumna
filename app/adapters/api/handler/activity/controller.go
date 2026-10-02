package activity

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("", handleActivityList(resolve))
}

func handleActivityList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		entityId, err := strconv.Atoi(c.Query("entity_id"))
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		entityType := c.Query("entity_type")
		activities, err := contract.AskActivityList(c.Request.Context(), resolve(), contract.ActivityList{
			EntityId:   entityId,
			EntityType: entityType,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		items := make([]activityDTO, len(activities))
		for i, activity := range activities {
			items[i] = toActivityDTO(activity)
		}
		httpx.Success(c, items)
	}
}
