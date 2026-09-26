# Анализатор размера диска (Go)

[![hexlet-check](https://github.com/sabudaye/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/sabudaye/go-project-242/actions)

Программа для анализа размера диска, которая позволяет получать информацию о занятом и свободном пространстве.

Учебный проект Хекслета: https://ru.hexlet.io/programs/go


## Стек

- Go

## Установка
```bash
git clone https://github.com/sabudaye/go-project-242.git
cd go-project-242
make build
```

## Использование
Help:
```bash
bin/hexlet-path-size -h
```

Usage
```bash
./bin/hexlet-path-size testdata/dir1
./bin/hexlet-path-size -r -H testdata/dir1
./bin/hexlet-path-size -r -a -H /var/log
```
---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
