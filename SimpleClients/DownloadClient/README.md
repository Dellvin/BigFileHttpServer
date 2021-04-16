 #  Клиент скачивания файла
 
Для запуска введит команду
    
    go run main.go -addr=http://localhost:8080 -id=5 -seek=0 -chunk=1024 -dest=/Users/dellvin/Desktop/word.docx
    
Флаги

    -addr:  адрес сервера
    -id:    номер файла для скачивания
    -seek:  смещение от начала файла
    -dest:  куда скачать
    -chunk: размер chunk в байтах