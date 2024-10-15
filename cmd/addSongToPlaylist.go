package cmd

import (
	"context"
	"fmt"
	"github.com/gremislaw/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addSongToPlaylistCmd = &cobra.Command{
	Use:   "addSongToPlaylist",
	Short: "add new song to playlist",
	Long: ``,
	Run: addSongToPlaylist,
}

func addSongToPlaylist(cmd *cobra.Command, args []string) {
	host, port := getHostPort()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)
	songPlaylist := new(api.SongPlaylist)
	songPlaylist.Song = &api.Song{Name: songName}
	songPlaylist.Playlist = &api.Playlist{Name: playlistName}
	response, err := client.AddSongToPlaylist(context.Background(), songPlaylist)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if response == nil {
		fmt.Println("Error:", "not found")
	} else {
		fmt.Println(response.Response)
	}
}

func init() {
	rootCmd.AddCommand(addSongToPlaylistCmd)
	addSongToPlaylistCmd.Flags().StringVar(&songName, "songName", "unknown", "")
	addSongToPlaylistCmd.Flags().StringVar(&playlistName, "playlistName", "unknown", "")
}
