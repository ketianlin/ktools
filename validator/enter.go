package validatorT

import (
	cnzh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

type Verify[T any] struct {
	structure T
}

func Enter[T any](structure T) *Verify[T] {
	return &Verify[T]{
		structure: structure,
	}
}

var (
	// Validate/v10 全局验证器
	validate *validator.Validate
	// 初始化Validate/v10国际化
	trans ut.Translator
)

// 初始化Validate/v10国际化
func init() {
	zh := cnzh.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	validate.SetTagName("v")
	//通过label标签返回自定义错误内容
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	zhTranslations.RegisterDefaultTranslations(validate, trans)

	//自定义required_if错误内容
	validate.RegisterTranslation("required_if", trans, func(ut ut.Translator) error {
		return ut.Add("required_if", "{0}为必填字段!", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_if", fe.Field())
		return t
	})
}
