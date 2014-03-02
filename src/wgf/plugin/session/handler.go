package session

import (
	"bytes"
	"hash/crc32"

	"wgf/sapi"
)

type Handler interface {

	//called when server started
	ServerInit(pServer *sapi.Server) error

	//called when server shutdown
	ServerShutdown() error

	//write session data into storage, expired in timeout seconds
	Write(sessionid string, data []byte, timeout int) error

	//set the session expired in timeout seconds
	Expire(sessionid string, timeout int) error

	//read session data
	Read(sessionid string) ([]byte, error)

	//delete all session data
	Del(sessionid string) error
}

type mapCacheBucket struct {
	Lock *sync.RWMutex
	Data map[string][]byte
	ExpireTime map[string]time.Time
}

func newMapCacheBucket() *mapCacheBucket{
	var re *mapCacheBucket
	re.Lock = &sync.RWMutex
	re.Data = make(map[string][]byte)
	re.ExpireTime = make(map[string]time.Time)
}

//default session storage handler
type MapCache struct {
	buckets []*mapCacheBucket
	bucketNum uint
}

func (c *MapCache) ServerInit(pServer *sapi.Server) error {
	c.bucketNum = 1000
	if c.bucketNum <= 0 {
		c.bucketNum = 1000
	}
	c.buckets = make([]mapCacheBucket, c.bucketNum)

	for key, _ := range c.buckets {
		c.buckets[key] = newMapCacheBucket()
	}
	return nil
}

func (c *MapCache) ServerShutdown() error {
	return nil
}

func (c *MapCache) Write(sessionid string, data []byte, timeout int) error {
	index := c.sidToIndex(sessionid)
	c.buckets[index].Lock.Lock()
	defer c.buckets[index].Lock.Unlock()

	c.buckets[index].Data[sessionid] = data
	c.buckets[index].ExpireTime[sessionid] = time.Now().Add(timeout * time.Second)
	return nil
}

func (c *MapCache) Expire(sessionid string, timeout int) error {
	index := c.sidToIndex(sessionid)
	c.buckets[index].Lock.Lock()
	defer c.buckets[index].Lock.Unlock()
	c.buckets[index].ExpireTime[sessionid] = time.Now().Add(timeout * time.Second)
	return nil
}

func (c *MapCache) Read(sessionid string) ([]byte, error) {
	index := c.sidToIndex(sessionid)
	c.buckets[index].Lock.RLock()
	defer c.buckets[index].Lock.RUlock()
	return c.buckets[index].Data[sessionid], nil
}

func (c *MapCache) Del(sessionid string) error {
	index := c.sidToIndex(sessionid)
	c.buckets[index].Lock.Lock()
	defer c.buckets[index].Lock.Unlock()
	delete(c.buckets[index], sessionid)
	return nil
}

func (c *MapCache) sidToIndex(sessionid string) uint32 {
	return crc32.ChecksumIEEE([]byte(sessionid))%c.bucketNum
}

