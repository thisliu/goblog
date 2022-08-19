package category

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
)

func Get(idstr string) (Category, error) {
	var category Category
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&category, id).Error; err != nil {
		return category, err
	}

	return category, nil
}

func GetAll(r *http.Request, perPage int) ([]Category, pagination.ViewData, error) {
	db := model.DB.Model(Category{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("home"), perPage)

	viewData := _pager.Paging()

	var categories []Category
	_pager.Results(&categories)

	return categories, viewData, nil
}

func (category *Category) Create() (err error) {
	if err = model.DB.Create(&category).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func (category *Category) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&category)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

func (category *Category) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&category)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}
