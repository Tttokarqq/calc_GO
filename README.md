calc_go
# http-сервер для параллельного расчета значений многочлена
## Описание
веб приложение, находящие значение арифмитеческого выражения, полученного из post-запроса
мой тг в группе лицея @Makszq
## Установка и использование 
открыть подходящую (т.е. удобную для пользователя) папку через редактор кода (я, конечно, использовал VSCode) и ввести в терминал:
`git clone https://github.com/Tttokarqq/calc_GO.git`

После скачивания репозитория перейти в появивщуюся папку calc_go:

`cd calc_go`

Затем ввести команду, которая запустит сервер:

`go run cmd/main.go`

дальше можно отправить выражение на сервер
1. Через Git Bach:
 
`curl --location 'localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'`

получим ответ: (номер нашего выражения на сервере)
{
    "id": "0"
}

2.  Через веб-сервисы, к примеру: https://reqbin.com/
на странице сайта введем адрес: `http://localhost:8080/api/v1/calculate`
запрос(json): `{"id": "0"}`
 и получим аналогичный ответ

получение списка выражений:

http://localhost:8080/api/v1/expressions
ответ:
{"expressions": [{
        "id": "0",
        "status": "wait",
        "result": "none"
        ]}
}
(т.е. ищем пример по идентификатору)


**изменение точности вычислений**

чтобы изменить количество знаков после запятой, которые будут сохраняться во вреся вычисления, нужно сделать запрос по адресу. 

`curl http://localhost:8080/api/v1/calculate/acc?accuracy=`

и в качестве параметра accuracy укзать необходимое число (по умолчанию 7, хранится в rpn.go в переменной Tochnost).
Обычно базового значения точности, хватает для большинства вычислений:

json запроса: 

{"expression": "55.064999 - 23 * -32 / ((29 - 3 * -2 -2) + 23 / 992 * (2 - (2 + 2)) * ( -1 * 2 - 8.32))"}

числовая часть ответа (accaccuracy = 7): `77.0492258`

Ожиадемая по факту: `77,0492257797868` или `77,0492258`

## Кода и виды запросов
доступна запись дробных чисел (через точку), но большинство возможных ошибок для них не проверяется

**405**: 
при попытке передать в параметр accuracy не верное значение (должно быть натуральное число меньше 65):

`curl http://localhost:8080/api/v1/calculate/acc?accuracy=-3` 
или 
`curl http://localhost:8080/api/v1/calculate/acc?accuracy=2.3` 

или 
`curl http://localhost:8080/api/v1/calculate/acc?accuracy=AAA_куда_я_ЖМАл??!`


Ответ:
{"error":"некорректное число точности. Необходимо целое из отрезка: [0;64]"}

**Прочие коды**

| json запроса | json ответ | код | пояснение |
|-----------------|---------|------|--------|
| {"expression": "45"} | {"result": "45"} | 200 | просто число |
| {"expression": "2.99 * 5.3 / 0.3"}| {"result": "52.823333333"} | 200 | операция с дробными |
| {"expression": "45. + 2.5"} | {"result": "47.5"} | 200 | отсутствие чисел знаков после точки не является ошибкой |
| {"expression": "2+2*2"}    | {"result": "6"}  | 200 | приоритет работает |
| {"expression": "(2+2)*2"}    | {"result": "8"}  | 200 | скобки работают |
| {"expression": "((((6)))+ (1))"} | {"result":"7"} | 200 | скобки раскрываются даже при абсурдном расположении |
| {"expression": "-+++---+(-1)"} | {"result": "-1"} | 200 | можно записывать перед числом сколько угодно плюсов и минусов, резлуьтат будет высчитан по правилам математики |
| {"expression": "((1---1)--4+3+23.99)/(1+2) * 10 /-+ -- 3 - (-90 * 2 * (2 - (2 + 2))) - 2"} | {"result": "-396.433333333" }| 200 | возможны более сложные завуалированные примеры|
| {"expression": "2  &emsp;      + 2*   2 "}| {"result": "6"} | 200 | программа игнорирует наличие\отсутсвие пробелов |
| {"expression": "(2 + 0) (1 - 0)"} | {"result": "21"} | 200 | !!! между скобками нельзы упускать знак *, ошибки не будет, но придет не корренктный результат |
| {"expression": "124.2149472323232312 "}| {"result": "124.2149472"}| 200 | округление до указанного значения (см. изменение точности вычисления) |
| {"expression": "124.2000200 "} | {"result": "124.20002"} | 200 | "лишние" нули после запятой автоматически уберутся, независимо от выстсавленной точности|
| {"expression": "5.3 / 0"}| {"error": "деление на 0"} | 422 | деление на ноль |
| {"expression": "2+2*2)"}| { "error": "ошибка в записи скобок"} | 422 | на закрытая правая скобка |
| {"expression": "2+(2*2"}| { "error": "ошибка в записи скобок"} | 422 | на закрытая левая скобка |
| {"expression": "A1 + 2"} | { "error": "ошибка - непредвиденный сивол"} | 422 | символ, не входящий в `.0123456789-+*/()` |
| {"expression": "2+} | {"error": "ошибка в записи знаков"} | 422 | последний символ не может быть знаком |
| {"expression": "2 ** 2"} | {"error": "ошибка в записи знаков"} | 422 | два идущих подярд ** или * / или / * или // |
| {"expression": " * 2"}| {"error": "ошибка в записи знаков"} | 422 | я намеренно сделал ошибкой запись / или * в начале выражения|
| {"expression": "45.0202.223"} | {"error": "ошибка в записи дробных чисел"} | 422 | попытка записать в число большей одной "."|


