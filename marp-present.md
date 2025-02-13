---
theme: cern
---

# Валидация на Golang

Игорь Лазарев
техлид istock.info
https://t.me/strider2038

---

## О себе

* Работаю в IT более 7 лет, около 3 лет пишу на Golang. Изначально писал на PHP.
* Интересы: микросервисная архитектура, Domain Driven Design, Golang.
* Последние 4 года работаю над проектом в промышленной сфере istock.info.
* Начинали с командой около 5 человек, сейчас нас 50+ (три продуктовых команды и менеджемент).

---

## Предпосылки

* Изначально прототип системы делался на основе PHP и фреймворка Symfony.
* Для валидации использовался компонент Symfony Validator.
    * Приятный и расширяемый синтаксис.
    * Есть встроенная система переводов.
    * Сложен в отладке, неудобен для corner cases (например, если нужно заинжектить зависимость).
* При миграции кода на Golang появилась необходимость:
    * использования инструмента валидации, похожего по функциональности;
    * плавного перехода для API;
    * валидации больших сложных структур, в том числе с рекурсией.

---

## Поиск готовых решений

* github.com/asaskevich/govalidator
* github.com/go-ozzo/ozzo-validation

---

## asaskevich/govalidator

```golang
type Document struct {
	Title   string   `valid:"required"`
	Keyword string   `valid:"stringlength(5|10)"`
	Tags    []string // нет встроенной проверки на длину слайса
}

func main() {
	document := Document{
		Title:   "",
		Keyword: "book",
	}

	result, err := govalidator.ValidateStruct(document)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
}
```

---

## asaskevich/govalidator

* Достоинства:
    * валидация на основе тегов;
    * для простых случаев - минимально и просто;
    * большое количество готовых правил (в виде функций).
* Недостатки:
    * теги используют reflection;
    * теги сложно тонко настроить (проблемы с экранированием строк);
    * отладка - сплошной ад ("почему не работает?");
    * сложно с переводами.

---

## github.com/go-ozzo/ozzo-validation

.code ozzo-validation.go /START/,/END/

---

## github.com/go-ozzo/ozzo-validation

* Достоинства:
    * настройка на основе языковых конструкций, а не тегов;
    * => код более гибкий и настраиваемый;
    * структурированные ошибки (код, сообщение, можно составить путь);
    * поддержка переводов;
    * есть много готовых правил, легко добавлять свои.
* Недостатки:
    * местами есть завязка на reflection => трудности с отладкой;
    * местами есть свободные типы `interace{}` => runtime errors;
    * систему переводов надо дорабатывать самостоятельно.

---

## Какие задачи хотелось решить

* Обеспечить полную статическую типизацию без `interface{}`
* Максимальную совместимость по API с Symfony Validator
* Более простую систему переводов
* Стиль, похожий на ozzo validation + структуру как в Symfony Validator для более простой миграции
* Гибкую систему для управления условной валидацией (запускать все, последовательно, одно из правил, группы валидации)
* Передавать контекстные параметры вглубь для сложных вложенных структур и с рекурсией
* Тонкий контроль над формированием путей к ошибкам

---

## github.com/muonsoft/validation

Особенности библиотеки

* Go version >= 1.18
* Гибкое и расширяемое API, созданое с учетом преимуществ статической типизации и дженериков
* Декларативный стиль описания процесса валидации
* Валидация различных типов: логические, числа любого типа, строки, слайсы, карты, `time.Time`
* Валидация собственных типов, удовлетворяющих интерфейсу Validatable
* Гибкая система ошибок с поддержкой переводов и склонений
* Простой способ описания собственных правил с передачей контекста и переводами ошибок
* Подробная документация с большим количеством примеров

---

## Простейший пример

.code basic-example.go /START/,/END/

---

## Валидация структуры

.code struct-validation.go /START/,/END/ HL00

---

## Валидация структуры, атрибут title

.code struct-validation.go /START/,/END/ HL01

---

## Валидация структуры, длина атрибута keywords

.code struct-validation.go /START/,/END/ HL02

---

## Валидация структуры, уникальные значения keywords

.code struct-validation.go /START/,/END/ HL03

---

## Валидация структуры, каждое значение keywords

.code struct-validation.go /START/,/END/ HL04

---

## Валидация структуры, перебор структурированных ошибок

.code struct-validation.go /START/,/END/ HL05

---

## Встроенные переводы

Внутри используется **`golang.org/x/text`**

.code json-marshal.go /START TRANSLATION/,/END TRANSLATION/ HL01

---

## Маршалинг в JSON (1)

.code json-marshal.go /START JSON/,/END JSON/ HL02

---

## Маршалинг в JSON (2)

```json
[
  {
    "error": "is blank",
    "message": "Значение не должно быть пустым.",
    "propertyPath": "title"
  },
  {
    "error": "too few elements",
    "message": "Эта коллекция должна содержать 5 элементов или больше.",
    "propertyPath": "keywords"
  },
  {
    "error": "is not unique",
    "message": "Эта коллекция должна содержать только уникальные элементы.",
    "propertyPath": "keywords"
  },
  {
    "error": "is blank",
    "message": "Значение не должно быть пустым.",
    "propertyPath": "keywords[0]"
  }
]
```

---

## Типовое использование, реализация интерфейса Validatable

.code typical-usage.go /START BASE/,/END BASE/ HL01

---

## Типовое использование, валидация слайса []Validatable

.code typical-usage.go /START BASE/,/END BASE/ HL02

---

## Типовое использование, запуск на типе Validatable

.code typical-usage.go /START EXEC/,/END EXEC/ HL03

---

## Немного экзотики, управляющие конструкции

* `validation.AtLeastOneOf()`
* `validation.Sequentially()`
* `validation.All()`
* `validation.When()`
* `validation.Async()`

.code at-least.go /START/,/END/

---

## Немного экзотики, группы валидации

.code groups.go /START USER/,/END USER/
.code groups.go /START VALIDATE/,/END VALIDATE/

---

## Результат

* Получился гибкий инструмент для решения всех поставленных задач.
    * Высокий уровень кастомизации.
    * Простые переводы.
    * Понятный стиль описания.
    * Можно использовать для описания сложных процедур валидации.
* На проекте используется уже несколько лет, эволюционировала вместе с запросами.
* Версия библиотеки near stable (на 95%), но пока еще `v0`.

---

## Ссылки

Спасибо за внимание!
Буду рад любой обратной связи.

.link https://github.com/muonsoft/validation github.com/muonsoft/validation
.image qr-code.gif
