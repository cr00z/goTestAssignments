# Задание

Написать скрипт, который успешно проходит опрос, расположенный по адресу: http://185.204.3.165

Скрипт должен быть написан на языке программирования GO.

Не нужно тестировать указанный сервер! Мы будем анализировать только код вашего скрипта.

Скрипт должен быть помещён на любой публичный GIT репозиторий. Тестовое задание присылается на проверку виде URL ссылки на GIT репозиторий.

# Решение

Запуск: `make run` или `go run main.go`

```
➜  goTestAssignments git:(main) ✗ make run
go run github.com/cr00z/parser
2022/10/19 13:41:44 auth completed
2022/10/19 13:41:44 question 1 answer: map[UslskpBX2eh48FkO:test g4kd8tWwHReP5iUI:hKQG6KNd8yIRi]
2022/10/19 13:41:44 question 2 answer: map[IX6DGcPMyPDxyJ1Y:oJjUmqeqcz7KZrSk]
2022/10/19 13:41:44 question 3 answer: map[M1WcIEZQ5IOOvwv7:s8xn7BttX9eUrGAL fK9iyB1yCyz8ktmo:8CXUMQcPx7n8ueh]
2022/10/19 13:41:44 question 4 answer: map[4alW8JbCeNp5SMkJ:test P1PRyTVhNQdGBXko:jY2vf1c2rvdEj]
2022/10/19 13:41:45 question 5 answer: map[ArhDvfNm61mqJ2z5:test aYMZUO9D4KE6VTKx:ufhvvdp2zQe4g1 q9EL444zylRSH4oB:5a5uQSjLpN8ZPC rWW14brDiLRWy0tr:test]
2022/10/19 13:41:45 question 6 answer: map[H4AVdfUuVL8IW9nv:test LV33rYS0qXxs1v5y:6Jy2FcmcoUdMinAZ]
2022/10/19 13:41:45 question 7 answer: map[jgITUBgGsKKBS9jL:tqnT2W00AmAOa oCwIBAhGR2aDadxA:test]
2022/10/19 13:41:45 question 8 answer: map[I1OKb6I6eSbUgTVb:GyC0BZG1Cfjb58x MWeIEA0hpxZXi3Vz:C7TXx4GqwCLCG bo19zdlTlOwKX3pG:test shSfeiWsyb3N9um8:test]
2022/10/19 13:41:45 test successfully passed
```