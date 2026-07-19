package storage

import (
	"sync"
	"t1"

	"github.com/google/uuid"
)

var globalOrderStore *OrderStore
var initStoreOnce sync.Once

type OrderStore t1.Cache[int, uuid.UUID]

func GetOrderStore() *OrderStore{
	initStoreOnce.Do(func(){
		globalOrderStore = (*OrderStore)(t1.NewCache[int, uuid.UUID]())
	})

	return globalOrderStore
}
