package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/category"
	"goblog/app/requests"
	"goblog/pkg/flash"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

type CategoriesController struct {
	BaseController
}

func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

func (ca *CategoriesController) Edit(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_category, err := category.Get(id)

	if err != nil {
		ca.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   view.D{},
		}, "categories.edit")
	}
}

func (ca *CategoriesController) Update(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_category, err := category.Get(id)

	// 3. 如果出现错误
	if err != nil {
		ca.ResponseForSQLError(w, err)
	} else {
		_category.Name = r.PostFormValue("name")

		errors := requests.ValidateCategoryForm(_category)

		if len(errors) == 0 {
			// 4.2 表单验证通过，更新数据
			rowsAffected, err := _category.Update()

			if err != nil {
				// 数据库错误
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
				return
			}

			if rowsAffected > 0 {
				showURL := route.Name2URL("categories.show", "id", id)
				http.Redirect(w, r, showURL, http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何更改！")
			}
		} else {
			// 4.3 表单验证不通过，显示理由
			view.Render(w, view.D{
				"Category": _category,
				"Errors":   errors,
			}, "categories.edit")
		}
	}
}

func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	errors := requests.ValidateCategoryForm(_category)

	if len(errors) == 0 {
		_category.Create()
		if _category.ID > 0 {
			flash.Success("分类创建成功")
			indexURL := route.Name2URL("articles.index")
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建分类失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errors,
		}, "categories.create")
	}
}

func (ca *CategoriesController) Articles(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_category, err := category.Get(id)

	if err != nil {
		ca.ResponseForSQLError(w, err)
	}

	articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 2)

	if err != nil {
		ca.ResponseForSQLError(w, err)
	} else {
		// ---  2. 加载模板 ---
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}

func (ca *CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_category, err := category.Get(id)

	if err != nil {
		ca.ResponseForSQLError(w, err)
	} else {
		// ---  2. 加载模板 ---
		view.Render(w, view.D{
			"Category": _category,
		}, "categories.show")
	}
}

func (ca *CategoriesController) Delete(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_category, err := category.Get(id)

	if err != nil {
		ca.ResponseForSQLError(w, err)
	} else {
		rowsAffected, err := _category.Delete()

		// 4.1 发生错误
		if err != nil {
			// 应该是 SQL 报错了
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			if rowsAffected > 0 {
				indexURL := route.Name2URL("articles.index")
				http.Redirect(w, r, indexURL, http.StatusFound)
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "404 文章分类未找到")
			}
		}
	}
}
