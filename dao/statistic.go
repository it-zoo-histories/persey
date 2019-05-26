package dao

import "github.com/arangodb/go-driver"

type StatisticDao struct {
	Database driver.Database
}

/*InitStatisticDao - инциализация статистического маршрута*/
func InitStatisticDao(bd driver.Database) *StatisticDao {
	statistic := &StatisticDao{}
	statistic.Database = bd
	return statistic
}

func (dao *StatisticDao) CreateNewStatistic() {}
