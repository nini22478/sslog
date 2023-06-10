package logics

import (
	"jdudp/models"
)

func NasList(datas []models.UserNas, tableId string) ([]models.UserNas, string, error) {
	dbr := OpenDb(tableId)
	dbr = dbr.Table("user_nas")
	//Count()必须在where之后limit和offset之后
	result := dbr.Debug().Find(&datas)
	return datas, tableId, result.Error
}
