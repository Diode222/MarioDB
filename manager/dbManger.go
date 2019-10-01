package manager

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"sync"
)

type dbManger struct {
	cache *lru.Cache
}

var manager *dbManger
var once sync.Once

func NewDBManger(maxCacheCount uint) *dbManger {
	lru, err := lru.New(int(maxCacheCount))
	if err != nil {
		log.Printf("Init cache cache failed, maxCacheCount: %d", maxCacheCount)
		panic(err)
	}
	once.Do(func() {
		manager = &dbManger{
			cache: lru,
		}
	})
	return manager
}

func (m *dbManger) Clear() {
	m.cache.Purge()
}

func (m *dbManger) Get(dbName string) (*leveldb.DB, bool) {
	var db *leveldb.DB
	dbInterface, ok := m.cache.Get(dbName)
	if ok {
		db = dbInterface.(*leveldb.DB)
	}
	return db, ok
}

// Contains will not update recent use info of cache
func (m *dbManger) Contains(dbName string) bool {
	return m.cache.Contains(dbName)
}

// Peek will not update recent use info of cache
func (m *dbManger) Peek(dbName string) (*leveldb.DB, bool) {
	var db *leveldb.DB
	dbInterface, ok := m.cache.Peek(dbName)
	if ok {
		db = dbInterface.(*leveldb.DB)
	}
	return db, ok
}

// Add db if not contains
func (m *dbManger) Add(dbName string, db *leveldb.DB) (bool, bool) {
	return m.cache.ContainsOrAdd(dbName, db)
}

// return: if dbName exists and being removed
func (m *dbManger) Remove(dbName string) bool {
	db, ok := m.Peek(dbName)
	ok = (ok && m.cache.Remove(dbName))
	if ok {
		db.Close()
	}
	return ok
}

func (m *dbManger) Len() int {
	return m.cache.Len()
}
