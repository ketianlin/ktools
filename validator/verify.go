package validatorT

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// VerifyStruct 验证结构体
func (v *Verify[T]) VerifyStruct() error {
	//fmt.Printf("%#v\n", v.structure)
	err := validate.Struct(v.structure)
	if err != nil {
		return errors.New(v.Translate(err))
	}
	return nil
}

// VerifyMap 验证Map verifyRuleTag：验证规则Tag字符串
func (v *Verify[T]) VerifyMap(verifyRuleTag string) error {
	fmt.Printf("%#v\n", v.structure)
	err := validate.Var(v.structure, verifyRuleTag)
	if err != nil {
		return errors.New(v.Translate(err))
	}
	return nil
}

// Translate 检验并返回检验错误信息
func (v *Verify[T]) Translate(err error) string {
	errs := err.(validator.ValidationErrors)
	return errs[0].Translate(trans)
}
