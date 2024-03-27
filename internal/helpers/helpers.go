package helpers

import (
	"fmt"
	"log"
	"runtime/debug"
)

// WrapErr оборачивает ошибку в text
func WrapErr(text string, err error) error {
	return fmt.Errorf("%s%w", text, err)
}

// WrapErrNotNil оборачивает ошибку в text если ошибка не nil
func WrapErrNotNil(text string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s%w", text, err)
}

// ServerError выводит ошибку и ее стэк в консоль, а также возвращает обернутую ошибку
// если сама ошибка не nil
func ServerError(text string, err error) error {
	if err == nil {
		return nil
	}

	log.Printf("[ERR]: %s\n%s", err.Error(), debug.Stack())
	return WrapErr(text, err)
}
