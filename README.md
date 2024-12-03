# kafka-go-service
# Описание проекта

Мой проект состоит из двух микросервисов, связанных с помощью **Kafka**. Сборка и запуск проекта производятся с помощью **Docker Compose**.

## Первый микросервис

Первый микросервис работает на портe **8081** и обрабатывает следующие HTTP-запросы:

- **POST** на адрес `/api/v1/set`
- **GET** на адрес `/api/v1`
- **POST** на адрес `/api/v1/set/generate`

Метод `GET` возвращает HTML-страничку, которая иллюстрирует работу сервиса. Создание этой страницы описано в директории `writer/template` в соответствующем файле. 

### Методы POST

1. **Добавление заказа в формате JSON**  
   Пользователь вводит информацию о заказе в формате JSON, которая записывается в тело запроса.

2. **Генерация заказов**  
   Пользователь может выбрать количество заказов, которые будут сгенерированы случайным образом. Это число отправляется на сервис в качестве аргумента запроса.

После обработки запросов необходимая информация о заказе отправляется в **Kafka**, откуда второй сервис считывает структуры для добавления в базу данных. Все действия с базой данных происходят на стороне сервиса **consumer**.

## Второй микросервис

Второй микросервис обрабатывает HTTP-запросы на порту **8080** с следующими маршрутами:

- **GET** по адресу `/api/v1/get`
- **GET** по адресу `/api/v1`
- `hb.GetHtml()`

Последний запрос возвращает HTML-страничку, которая иллюстрирует работу сервиса. Создание этой страницы описано в директории `consumer/template` в соответствующем файле.

### Обработка запросов

При вызове первого запроса пользователь вводит в окно на странице фронта **ID** запроса, информацию о котором он хочет получить. 

- Сервис получает **ID** в аргументах запроса и возвращает всю информацию о заказе в формате **JSON**.

При попытке получения информации о запросе сервис обращается к кэшу. Если запрашиваемый **ID** отсутствует в кэше, происходит обращение к базе данных **PostgreSQL** для поиска нужного **ID**. 

### Архитектура сервиса

В сервисной части **consumer** реализованы методы для работы с кэшем и базой данных. 

- **Consumer** считывает из Kafka записи, которые необходимо добавить в базу данных, и выполняет вставку сначала в кэш, затем и в базу данных. 
- Записи в кэше хранятся **1 час**, после чего удаляются. Проверка времени жизни выполняется каждые **10 минут**.

## Тестирование

Код покрыт тестами на go с использованием mock. Для удобного тестирования в репозиторий загружена **Postman-коллекция**

## Архитектура проекта

В ходе написания проекта придерживалась принципов **чистой архитектуры**, правил **SOLID** и метода нормальных форм в проектировании баз данных. В данном проекте структура базы данных находится в **2 НФ**.

Схема архитектуры проекта:


![image](https://github.com/user-attachments/assets/28d2f410-5fe2-4bce-8e35-a983b1c362c5)




Схема базы данных:


![image](https://github.com/user-attachments/assets/768088b6-c2f6-40a9-8087-d396c561d048)

