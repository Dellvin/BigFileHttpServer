 #  Клиент загрузки файла с сервера
 
Для запуска введит команду
    
    go run main.go -addr=localhost:8080 -chunk=1024 -src=/Users/dellvin/Desktop/Space_of_Space.docx
    
Флаги

    -addr:  адрес сервера
    -src:   откуда качать
    -chunk: размер chunk в байтах