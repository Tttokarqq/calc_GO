package rpn

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var(
	err_skobk = errors.New("ошибка в записи скобок")
	err_symbl = errors.New("ошибка - непредвиденный сивол")
	err_znak = errors.New("ошибка в записи знаков")
	err_float = errors.New("ошибка при обработке дробных значений")
)

func IsNum(s byte) bool{ // среди цифр есть ".", для работы с десятичными дробями /
// ошибки для них не описаны, т.к. у меня были более важные дела (партеечка в хойку), да и вообще это необязательно
	nums := ".0123456789"
	for i := 0; i < len(nums); i++{
		if s == nums[i]{return true}
	}
	return false
}

func IsZnak(s byte) bool{
	nums := "-+*/"
	for i := 0; i < len(nums); i++{
		if s == nums[i]{return true}
	}
	return false
}

func Calc(expression string) (string, error) { 
	for _, i := range expression{ // удаление пробелов, хотя можно было просто continue :D
		if string(i) == " "{
			ind := strings.Index(expression, string(i))
			expression = expression[0:ind] + expression[ind + 1:]
		}
	}
	// боже как же я намучился со скобками, которые все усложняют и портят
	// цикл, отправляющий в рекурсию найденное выражение в скобках (самое последнее, с максимальным приоритетом)
	// и заменяющий его на возращенную строку (результат вычисления рекусрсии)
	index := strings.Index(expression, "(") + 1 // индекс первой открытой скобки
	index_left := 1 // количество открытых левых скобок
	for strings.Index(expression, "(") != -1 && index < len(expression){
		symbol := expression[index]
		if string(symbol) == "("{
			index_left += 1
		} else if string(symbol) == ")"{
			index_left -= 1
		}
		if index_left == 0{
			m, err := Calc(expression[strings.Index(expression, "(") + 1 :index])
			if err != nil{
				return "0", err
			}
			expression = expression[0:strings.Index(expression, "(")] + m + expression[index + 1:]
			index = strings.Index(expression, "(") + 1
			index_left = 1
			// fmt.Println(expression[strings.Index(expression, "(") + 1 :index])
		}  
		if index_left > 1 && index + 1 == len(expression){ // когда последний символ, но есть не закрытые скобки
			fmt.Println(expression, index_left)
			return "0", err_skobk
		}
		index++
	}
	// работа со знаками 
	for i := 0; i < len(expression); i++{
		symbol := expression[i]
		if i + 1 == len(expression) && IsZnak(symbol){ // если в конце выражения стоит знак 
			// fmt.Println(expression)
			return "-1", err_znak
		}
		if IsZnak(symbol){
			if i + 1 < len(expression){  // очистка "лишних" плюсов и минусов
				if string(symbol) == "+" && string(expression[i + 1]) == "-" || // когда подяряд -+ или +- -> -
					string(symbol) == "-" && string(expression[i + 1]) == "+" {
					expression = expression[0:i] + "-" +  expression[i + 2:]
					i = -1
					continue
				} else if string(symbol) == "-" && string(expression[i + 1]) == "-" { // когда подряд -- -> +
					expression = expression[0:i] + "+" +  expression[i + 2:] 
					i = -1
					continue
				} else if string(symbol) == "+" && string(expression[i + 1]) == "+" { // когда подряд ++ -> +
					expression = expression[0:i] + "+" + expression[i + 2:]
					i = -1
					continue
				} else if ((string(symbol) == "*" || string(symbol) == "/") &&
					string(expression[i + 1]) == "+") { // когда после * или / идет + -> ""
					expression = expression[0:i + 1] +  expression[i + 2:] 
					i = -1
					continue
				} else if ((string(symbol) == "-" || string(symbol) == "+") &&
					(string(expression[i + 1]) == "*") || string(expression[i + 1]) == "/") { // когда после - или + идет / или * -> ""
					return "-1", err_znak
				} else if (i == 0 && string(symbol) == "+"){ // если начинается с +
					expression = expression[1:]
				}
			} 
		}
	}
	// рассчет выражений, в котором не осталось скобок и повторяющихся корректно записаны знаки
	num1, num2 := "", "" // числа хранятся как строки
	znak1, znak2 := "", "" // знаки между числами (минус перед числом не записывается, а идет  в num1_znak)
	num1_znak, num2_znak := 1.0, 1.0 // знаки num1 и num2
	ind1, ind2 := 0, 0 // индексы начала num1 и num2 
	for i := 0; i < len(expression); i++{
		symbol := expression[i]
		if IsZnak(symbol){
			if string(symbol) == "-" && num1_znak == 1 && num1 == ""{
				num1_znak = -1.0 // смена знака
			} else if string(symbol) == "-" && num2 == "" && (znak1 == "*" || znak1 == "/"){
				num2_znak = -1.0
			} else if znak1 == ""{
				znak1 = string(symbol)
			} else if znak2 == ""{
				znak2 = string(symbol)
			}
		} else if IsNum(symbol){
			if znak1 == ""{
				if num1 == ""{
					ind1 = i
				}
				num1 += string(symbol)
			} else if znak1 != "" && znak2 == ""{
				// fmt.Println(string(symbol))
				if num2 == ""{
					ind2 = i
				}
				num2 += string(symbol)
			}
		} else {
			if string(symbol) == ")" { // найдена правая не парная скобка
				return "0", err_skobk
			} else {
				return "0", err_symbl
			}
		}
		if i + 1 == len(expression) || znak2 != ""{ 
			// промежуточное вычисление, когда индекс дошел до конца выражения или найдем второй знак
			if znak1 == ""{
				if num1_znak > 0{
					return num1, nil
				} else {
					return "-" + num1, nil
				}
			}
			if (znak1 == "+" || znak1 == "-") && (znak2 == "*" || znak2 == "/"){
				m, err := Calc(expression[ind2:]) // отправление в рекурсию, если второй знак с большим приоритетом
				if err != nil{
					return "0", err
				}
				if num1_znak == -1.0{
					expression = expression[0:ind2 - 1] + m // изменение выражения, с учетом "ответа рекурсии"
				} else {
					expression = expression[0:ind2] + m
				}
				i, num1, num2, znak1, znak2, num1_znak, num1_znak, ind1, ind2 = -1, "", "", "", "", 1.0, 1.0, 0, 0 
				// после высчитывания каждой операции ВСЕ сбравсывается и начинается сначала
			} else {
				num1_, _ := strconv.ParseFloat(num1, 64) // перевод из строки(num1) в дробное(num1_)
				num1_ *= num1_znak
				num2_, _ := strconv.ParseFloat(num2, 64)
				num2_ *= num2_znak
				znach := ""
				if znak1 == "-"{
					znach = fmt.Sprintf("%v", num1_ - num2_) 
				} else if znak1 == "+" {			
					znach = fmt.Sprintf("%v", num1_ + num2_) 
				} else if znak1 == "*" {
					znach = fmt.Sprintf("%v", num1_ * num2_) 
				} else if znak1 == "/" {
					if num2_ == 0 {return "-1", errors.New("деление на 0")}
					znach =  fmt.Sprintf("%.9f", num1_ / num2_)
				}
				if num1_znak == -1.0 { // если не будет этой проверки, то 1 - 2 * - 3 -> 1--6, а не 1+6
					expression = expression[1:ind1] + znach + expression[ind2 + len(num2):]
				} else {
					expression = expression[0:ind1] + znach + expression[ind2 + len(num2):]
				}
				i, num1, num2, znak1, znak2, num1_znak, num1_znak, ind1, ind2, znach = -1, "", "", "", "", 1.0, 1.0, 0, 0, ""
				
			}
			// fmt.Println(expression)
		}
	}
	// fmt.Println("expression:", expression)
	return expression, nil
}