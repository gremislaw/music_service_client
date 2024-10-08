package cmd

import (
	"context"
	"fmt"
	"github.com/gremislaw/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addPlaylistCmd = &cobra.Command{
	Use:   "addPlaylist",
	Short: "create new playlist",
	Long: ``,
	Run: addPlaylist,
}

var playlistName string

func addPlaylist(cmd *cobra.Command, args []string) {
	host, port := getHostPort()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)
	playlist := new(api.Playlist)
	playlist.Name = playlistName
	response, err := client.AddPlaylist(context.Background(), playlist)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(response.Response)
}

func init() {
	rootCmd.AddCommand(addPlaylistCmd)
	addPlaylistCmd.Flags().StringVar(&playlistName, "playlistName", "unknown", "")
}
