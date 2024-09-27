# music_service_client

## Клиент реализован на фреймворке cobra-cli.


### Команда сборки:
make

### Список доступных комманд:
./music_service --help

### Просмотр необходимых опций комманд:
./music_service add --help

### Примеры:
./music_service play \
./music_service add --songName "Hello" --songAuthor "Adele" --duration 31 

### Интерактивный режим:
./music_service interactive
### Интерактивный режим пример:
./music_service interactive \
play \
add wrong \
add "Hello" "Adele" 90
