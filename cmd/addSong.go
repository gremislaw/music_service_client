package cmd

import (
	"context"
	"fmt"
	"github.com/gremislaw/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addSongCmd = &cobra.Command{
	Use:   "addSong",
	Short: "addSong new song",
	Long:  ``,
	Run:   addSong,
}

var songName string
var songAuthor string
var duration int

func addSong(cmd *cobra.Command, args []string) {
	host, port := getHostPort()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)
	song := new(api.Song)
	song.Author = songAuthor
	song.Name = songName
	song.Duration = int64(duration)
	response, err := client.AddSong(context.Background(), song)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(response.Response)
}

func init() {
	rootCmd.AddCommand(addSongCmd)
	addSongCmd.Flags().StringVar(&songName, "songName", "unknown", "")
	addSongCmd.Flags().StringVar(&songAuthor, "songAuthor", "unknown", "")
	addSongCmd.Flags().IntVar(&duration, "songDuration", 30, "")
}
