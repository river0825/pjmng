// package cache

// import (
// 	"time"

// 	// "nikeyi/beiqilai-server/module/account/domain/entity"
// 	// "nikeyi/beiqilai-server/module/account/domain/value_object"

// 	"github.com/jellydator/ttlcache/v3"
// )

// type InMemCache struct {
// 	cache  *ttlcache.Cache
// }

// // NewInMemCache creates a new instance of InMemCache
// func NewInMemCache() *InMemCache {
// 	cache := ttlcache.New[value_object.UserID, *entity.User](
// 		ttlcache.WithTTL[value_object.UserID, *entity.User](30 * time.Minute),
// 	)

// 	return &InMemCache{
// 		userCache: cache,
// 	}
// }

// // GetUser retrieves a user from the cache by userID
// func (c *InMemCache) GetUser(userID value_object.UserID) (user *entity.User) {
// 	if !c.userCache.Has(userID) {
// 		return nil
// 	}

// 	u := c.userCache.Get(userID)
// 	if u == nil {
// 		return nil
// 	}

// 	return u.Value()
// }

// // SetUser stores a user in the cache with the given userID
// func (c *InMemCache) SetUser(userID value_object.UserID, user *entity.User) {
// 	c.userCache.Set(userID, user, ttlcache.DefaultTTL)
// }

// func (c *InMemCache) DeleteUser(userID value_object.UserID) {
// 	c.userCache.Delete(userID)
// }