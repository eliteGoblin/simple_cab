package cab_info_manager

import (
	"os"
	"simple_cab/data_model"
	log "simple_cab/logging"
	"simple_cab/toolkit"
	"sync"
)

type CabInfoManager struct {
	cacheRWMutex sync.RWMutex
	cache        cabInfoCache
}

var managerInit sync.Once
var manager *CabInfoManager

func GetManagerInstance() *CabInfoManager {

	managerInit.Do(func() {
		manager = &CabInfoManager{}
		if err := manager.cache.rebuild(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	})
	return manager
}

func (selfPtr *CabInfoManager) UpdateCache() {
	newCache := cabInfoCache{}
	finishChn := make(chan int)
	go func(cache *cabInfoCache) {
		cache.rebuild()
		finishChn <- 1
	}(&newCache)
	<-finishChn
	selfPtr.cacheRWMutex.Lock()
	defer selfPtr.cacheRWMutex.Unlock()
	selfPtr.cache = newCache
}

func (selfPtr *CabInfoManager) GetCabTripCount(medallion string, date string) (tripCount int, err error) {
	selfPtr.cacheRWMutex.RLock()
	currentCache := selfPtr.cache
	selfPtr.cacheRWMutex.RUnlock()
	// make date complete as date time string
	tm, err := toolkit.GetTimeFromDateTimeString(date + "T00:17:00+11:00")
	if err != nil {
		return 0, err
	}
	return currentCache.getTripCount(medallion, &tm), nil
}

func (selfPtr *CabInfoManager) GetCabTripCountFromDB(medallions []string, date string) (
	medallionCount map[string]int, err error) {
	return data_model.GetCabTripDataInDay(medallions, date)
}
