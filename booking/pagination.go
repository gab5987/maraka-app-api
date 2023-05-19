package booking

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

type Page struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}

func Paginate(c *gin.Context) options.FindOptions {
	paginate := Page{
		Limit: 10,
		Page:  1,
	}

	_limit := c.Query("limit")
	if _limit != "" {
		lim, _ := strconv.Atoi(_limit)
		paginate.Limit = int64(lim)
	}

	_page := c.Query("page")
	if _page != "" {
		pg, _ := strconv.Atoi(_page)
		paginate.Page = int64(pg)
	}

	fOpt := options.FindOptions{}

	fOpt.SetSort(bson.D{{"startDate", -1}})
	fOpt.SetLimit(paginate.Limit)
	fOpt.SetSkip(paginate.Page*paginate.Limit - paginate.Limit)

	return fOpt
}
