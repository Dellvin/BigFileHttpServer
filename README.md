#   Big file http server
Для запуска сервера нужно сделать несколько вещей:

1) Создать файл конфигурации в папке /MainApplication/config (пример находится в файле "config_example.txt").
В этом файле нужно указать следующие переменные:

       DbUser = "имя пользователя PostgreSQL"
    	DbPassword   = "пароль пользователя PostgreSQL"
    	DbDB         = "название базы данных"
    	Port         = ":порт сервера"
    	ReadTimeout  = тайм аут на чтение
    	WriteTimeout = тайс аут на запись

2) Синхронизировать все зависимости:

        go get ...
        
3) Ну и последнее:

        go rum ./cmd/main.go
        

    
    
    	
 