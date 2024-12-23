calc_go
# http-сервер для расчета значений многочлена
## Описание
веб приложение, находящие значение арифмитеческого выражения, полученного из post-запроса
## Установка и использование 
открыть подходящую (т.е. удобную для пользователя) папку через редактор кода (я, конечно, использовал VSCode) и ввести в терминал:
`git clone https://github.com/MaksaNeNegr/calc_go.git`

После скачивания репозитория перейти в появивщуюся папку calc_go:

`cd calc_go`

Затем ввести команду, которая запустит сервер:

`go run cmd/main.go`

дальше можно отправить выражение на сервер
1. Через Git Bach:
 
`curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'`

получим ответ: `{"result":"6"}`

2.  Через веб-сервисы, к примеру: https://reqbin.com/
на странице сайта введем адрес: `http://localhost:8080/api/v1/calculate`
запрос(json): `{"expression": "2+2*2"}` и получим аналогичный ответ

## Кода и виды запросов
доступна запись дробных чисел, но ошибка для них не проверяются
на любой не пост запрос сервер ответим кодом 500 

| json запроса | json ответ | код | пояснение|
|-----------------|---------|------|--------|
| {"expression": "2.99 * 5.3 / 0.3"}| {"result": "52.823333333"} | 200 | операция с дробными |
| {"expression": "2+2*2"}    | {"result": "6"}  | 200 | - |
| {"expression": "(2+2)*2"}    | {"result": "8"}  | 200 | - |
| {"expression": "-+++---+(-1)"} | {"result": "-1"} | 200 | можно записывать перед числом сколько угодно плюсов и минусов, резлуьтат будет высчитан по правилам математики |
|{"expression": "2     + 2*   2 "}| {"result": "6"} | 200 | программа игнорирует наличие\отсутсвие пробелов |
| {"expression": "1 * "}| {"error": "ошибка в записи знаков"} | 422 | попытка записать знак ( любой ) в конце выражения приведет к ошибке|
| {"expression": "5.3 / 0"}| {"error": "деление на 0"} | 422 | деление на ноль |
| {"expression": "2+2*2)"}| { "error": "ошибка в записи скобок"} | 422 | на закрытая правая скобка |
| {"expression": "2+(2*2"}| { "error": "ошибка в записи скобок"} | 422 | на закрытая левая скобка |
| {"expression": "A1 + 2"} | { "error": "ошибка - непредвиденный сивол"} | 422 | символ, не входящий в `.0123456789-+*/()` |