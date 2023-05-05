package booking

import (
	"github.com/gin-gonic/gin"
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

	skip := int64(paginate.Page*paginate.Limit - paginate.Limit)
	fOpt := options.FindOptions{Limit: &paginate.Limit, Skip: &skip}

	return fOpt
}
