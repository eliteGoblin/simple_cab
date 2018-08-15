package data_model

import (
	"fmt"
	"simple_cab/data_model/db"
)

type CabTripData struct {
	Medallion        string  `gorm:"type:text"`
	HackLicense      string  `gorm:"type:text"`
	VendorId         string  `gorm:"type:text"`
	RateCode         *int    `gorm:"default:0"`
	StoreAndFwdFlag  string  `gorm:"type:text"`
	PickupDatetime   string  `gorm:"type:timestamp;"`
	DropoffDatetime  string  `gorm:"type:timestamp;"`
	PassengerCount   *int    `gorm:"default:0"`
	TripTimeInSecs   *int    `gorm:"default:0"`
	TripDistance     float64 `gorm:"default:0"`
	PickupLongitude  float64 `gorm:"default:0"`
	PickupLatitude   float64 `gorm:"default:0"`
	DropoffLongitude float64 `gorm:"default:0"`
	DropoffLatitude  float64 `gorm:"default:0"`
}

func GetAllCabData() (cabTripData []*CabTripData, err error) {
	cabTripData = make([]*CabTripData, 0)
	sql := "select medallion, pickup_datetime from cab_trip_data order by pickup_datetime"
	err = db.MyDB.Raw(sql).Find(&cabTripData).Error
	return cabTripData, err
}

type TripCountInfo struct {
	Medallion string `gorm:"type:text"`
	TripCount int    `gorm:"default:0"`
}

func GetCabTripDataInDay(medallions []string, dateStr string) (medallionCount map[string]int, err error) {
	response := make([]TripCountInfo, 0)
	medallionCount = make(map[string]int)
	startOfDay := dateStr + " 00:00:00"
	endOfDay := dateStr + " 23:59:59"
	dateWhereClause := fmt.Sprintf(" where medallion in '%s' AND pickup_datetime BETWEEN %s AND %s GROUP BY medallion",
		formatInOperator(medallions), startOfDay, endOfDay)
	sql := "select medallion, COUNT(medallion) AS trip_count from cab_trip_data order by pickup_datetime " + dateWhereClause
	err = db.MyDB.Raw(sql).Find(&response).Error
	if err != nil {
		return medallionCount, err
	}
	for _, v := range response {
		medallionCount[v.Medallion] = v.TripCount
	}
	return medallionCount, nil
}

func formatInOperator(set []string) string {
	result := "( "
	for i, v := range set {
		if 0 == i {
			result += v
		} else {
			result += fmt.Sprintf(" , %s", v)
		}
	}
	result += " )"
	return result
}
