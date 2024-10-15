package cmd

import (
	"context"
	"fmt"
	"github.com/gremislaw/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var deleteSongFromPlaylistCmd = &cobra.Command{
	Use:   "deleteSongFromPlaylist",
	Short: "delete song from playlist",
	Long: ``,
	Run: deleteSongFromPlaylist,
}

func deleteSongFromPlaylist(cmd *cobra.Command, args []string) {
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
	response, err := client.DeleteSongFromPlaylist(context.Background(), songPlaylist)
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
	rootCmd.AddCommand(deleteSongFromPlaylistCmd)
	deleteSongFromPlaylistCmd.Flags().StringVar(&songName, "songName", "unknown", "")
	deleteSongFromPlaylistCmd.Flags().StringVar(&playlistName, "playlistName", "unknown", "")
}
