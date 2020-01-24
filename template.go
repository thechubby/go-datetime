package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageVariables struct {
	Date string
	Time string
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()              // Находим время
	HomePageVars := PageVariables{ //Загоняем дату и время в структуру
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("homepage.html") //парсим html файл
	if err != nil {                                // если ошибка
		log.Print("template parsing error: ", err) // выводи её
	}
	err = t.Execute(w, HomePageVars) //запускаем шаблон и пропускаем HomePageVars структуру чтобы заполнить поля
	if err != nil {                  // еслт ошибка
		log.Print("template executing error: ", err) //выводи её
	}
}
