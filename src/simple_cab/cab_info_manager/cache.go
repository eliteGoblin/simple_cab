package cab_info_manager

import (
	"fmt"
	"simple_cab/data_model"
	"simple_cab/logging"
	log "simple_cab/logging"
	"simple_cab/toolkit"
	"time"
)

var _ = logging.Info

type cabInfoCache struct {
	minTime  time.Time
	daySlots []map[string]int
}

func (selfPtr *cabInfoCache) rebuild() error {
	allCabData, err := data_model.GetAllCabData()
	if err != nil {
		return err
	}
	selfPtr.minTime, err = toolkit.GetTimeFromDateTimeString(allCabData[0].PickupDatetime)
	if err != nil {
		return err
	}
	maxTime, errEndDay := toolkit.GetTimeFromDateTimeString(allCabData[len(allCabData)-1].PickupDatetime)
	if errEndDay != nil {
		return fmt.Errorf("invalid datetime found: %+v", *allCabData[len(allCabData)-1])

	}
	selfPtr.daySlots = make([]map[string]int, toolkit.GetDiffByDays(&selfPtr.minTime, &maxTime)+16)
	for _, cabTripInfo := range allCabData {
		dayOfTrip, err := toolkit.GetTimeFromDateTimeString(cabTripInfo.PickupDatetime)
		if err != nil {
			log.Warnf("invalid datetime found: %+v", dayOfTrip)
			continue
		}
		selfPtr.incrementTripCountByOne(cabTripInfo.Medallion, &dayOfTrip)
	}
	log.Infof("rebuild cache completed! cache: %v", selfPtr.daySlots)
	return nil
}

func (selfPtr *cabInfoCache) incrementTripCountByOne(medallion string, tm *time.Time) {
	slot := toolkit.GetDiffByDays(&selfPtr.minTime, tm)
	if slot > len(selfPtr.daySlots) || slot < 0 {
		log.Errorf("invalid slot got: %s %+v", medallion, *tm)
		return
	}
	if selfPtr.daySlots[slot] == nil {
		selfPtr.daySlots[slot] = make(map[string]int)
	}
	selfPtr.daySlots[slot][medallion]++
}

func (selfPtr *cabInfoCache) getTripCount(medallion string, tm *time.Time) (tripCount int) {
	slot := toolkit.GetDiffByDays(&selfPtr.minTime, tm)
	if slot > len(selfPtr.daySlots) || slot < 0 {
		return 0
	}
	if mp := selfPtr.daySlots[slot]; mp == nil {
		return 0
	} else {
		return mp[medallion]
	}
}
