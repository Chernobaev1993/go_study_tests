package main

import (
	"fmt"
	"github.com/Chernobaev1993/go_study_tests/basics"
)

func main() {
	start()
	basics.Syntax()
}

func start() {
	start_message := "Приветствую в интерактивном обучении по Go!\nВыбери номер категории обучения:\n"

	cat_1 := "1. БАЗА\n"

	subcat_1 := "1. Синтаксис. Переменные и их объявление. Типы данных. Приведение\n"
	subcat_2 := "2. Циклы. Условные конструкции. Ошибки. Паника\n"
	subcat_3 := "3. Массивы. Слайсы. Карты. Структуры\n"
	subcat_4 := "4. Пакеты. Импорт. Функции\n"
	fmt.Println(start_message, cat_1, subcat_1, subcat_2, subcat_3, subcat_4)
}

