package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
)

type StoreService struct {
}

func (s *StoreService) List(search string) (list []app.AppStore, err error) {
	db := global.GVA_DB.Model(&app.AppStore{})

	if search != "" {
		db.Where("store_name LIKE ?", "%"+search+"%")
	}
	var stores []app.AppStore
	err = db.Find(&stores).Error
	return stores, err
}

func (s *StoreService) Detail(id int) (store app.AppStore, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&store).Error
	return
}
