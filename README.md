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

### Пример работы с плейлистами и песнями:
- `./music_service add --songName "Hello" --songAuthor "Adele" --songDuration 31`
  
- `./music_service addSongToPlaylist --songName "Hello" --playlistName "Pop"`
  
- `./music_service getPlaylist --playlistName "Pop"`
  
- `./music_service play`
  
- `./music_service deletePlaylist --playlistName "Pop"`

<details>
  <summary><strong> 
    Пример использования cobra-cli
  </strong></summary>

  ```golang
  package cmd

import (
	"context"
	"fmt"
	"github.com/gremislaw/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "play current song",
	Long: ``,
	Run: play,
}


func play(cmd *cobra.Command, args []string) {
	host, port := getHostPort()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)
	response, err := client.Play(context.Background(), &api.Empty{})
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(response.Response)
}

func init() {
	rootCmd.AddCommand(playCmd)
}

  ```

</details>

### Команда для очистки ненужных файлов:
- `make clean`

### Команда для пересборки:
- `make rebuild`

### Команда для форматирования кода:
- `make format`