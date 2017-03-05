package middleware

import (
	"github.com/labstack/echo"
	"github.com/Sirupsen/logrus"
	"github.com/patrickmn/go-cache"
	"time"
	"MicroFilm/model"
	"github.com/gocraft/dbr"
)

const (
	CacheKey = "Cache"
	CateListKey = "CateList"
)

func CacheHandler(db *dbr.Session) echo.MiddlewareFunc {

	// 应该定义长久缓存、短期缓存
	// 创建一个默认的缓存过期时间5分钟,清洗过期物品每30秒
	cache := cache.New(12 * time.Hour, 1 * time.Hour)
	initCacheData(cache, db)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(ctx echo.Context) error {

			ctx.Set(CacheKey, cache)
			return next(ctx)
		})
	}
}

// shortcut to get Cache
//func Default(ctx *echo.Context) cache.Cache {
//	// return c.MustGet(DefaultKey).(ec.CacheStore)
//	return ctx.Get(CacheKey).(cache.Cache)
//}

func CacheCateList(ctx echo.Context) (interface{}) {

	c := ctx.Get(CacheKey).(*cache.Cache)

	list, found := c.Get(CateListKey)
	if (found) {
		return list
	} else {
		return nil
	}
}

func initCacheData(c *cache.Cache, db *dbr.Session) {

	tx, _ := db.Begin()

	var list model.CategoryList
	if err := list.Load(tx); err != nil {
		tx.Rollback()
		logrus.Debug("电影分类列表查询失败")
	} else {
		logrus.Debug("电影分类列表: ", list)
		c.Set(CateListKey, list, cache.DefaultExpiration)
	}
	tx.Commit()
}