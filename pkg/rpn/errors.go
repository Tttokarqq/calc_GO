package rpn

import (
	"errors"
)

var(
	err_skobk = errors.New("ошибка в записи скобок")
	err_symbl = errors.New("ошибка - непредвиденный сивол")
	err_znak = errors.New("ошибка в записи знаков")
	err_float = errors.New("ошибка при обработке дробных значений")
)