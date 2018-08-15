package cab_info_manager

type CabInfoRetriever interface {
	UpdateCache()
	GetCabTripCount(medallion string, date string) (tripCount int, err error)
	GetCabTripCountFromDB(medallion string, date string) (tripCount int, err error)
}
