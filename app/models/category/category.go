package category

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

func (category Category) Link() string {
	return route.Name2URL("categories.show", "id", category.GetStringID())
}

func (category Category) ArticlesLink() string {
	return route.Name2URL("categories.articles", "id", category.GetStringID())
}
