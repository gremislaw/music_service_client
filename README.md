# music_service_client

## Клиент реализован на фреймворке cobra-cli. Клиент дает возможность использовать функционал сервиса через соответствующие команды.

* Реализация сервера - https://github.com/gremislaw/music_service

### Команда сборки:
- `make`

### Перейти в папку `bin`
- `cd bin`

### Список доступных комманд:
- `./music_service --help`

### Просмотр необходимых опций комманд:
- `./music_service add --help`

### Пример:
- `./music_service add --songName "Hello" --songAuthor "Adele" --songDuration 31`
  
- `./music_service addSongToPlaylist --songName "Hello" --playlistName "Pop"`
  
- `./music_service getPlaylist --playlistName "Pop"`
  
- `./music_service play`
  
- `./music_service deletePlaylist --playlistName "Pop"`
