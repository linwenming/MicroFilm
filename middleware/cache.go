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

func DataCacheHandler() echo.MiddlewareFunc {

	// 应该定义长久缓存、短期缓存
	// 创建一个默认的缓存过期时间5分钟,清洗过期物品每30秒
	_cache := cache.New(12 * time.Hour, 1 * time.Hour)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			c.Set(CacheKey, _cache)

			_, found := _cache.Get(CateListKey)
			if !found {
				tx := c.Get("Tx").(*dbr.Tx)
				var cateList model.CategoryList

				if err := cateList.Load(tx); err == nil {
					logrus.Debug("电影分类列表: ",cateList)
					_cache.Set(CateListKey, cateList, cache.DefaultExpiration)
				}
			}

			return next(c)
		})
	}
}