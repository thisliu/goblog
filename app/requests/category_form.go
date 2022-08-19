package requests

import (
	"goblog/app/models/category"

	"github.com/thedevsaddam/govalidator"
)

func ValidateCategoryForm(data category.Category) map[string][]string {
	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:1", "max_cn:10"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需大于 1",
			"max_cn:名称长度需小于 10",
		},
	}

	// 3. 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 4. 开始验证
	return govalidator.New(opts).ValidateStruct()
}
