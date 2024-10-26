package cmd

import (
	"context"
	"fmt"
	"github.com/gremislaw/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var printPlaylistCmd = &cobra.Command{
	Use:   "printPlaylist",
	Short: "print all songs in playlist",
	Long:  ``,
	Run:   printPlaylist,
}

func printPlaylist(cmd *cobra.Command, args []string) {
	host, port := getHostPort()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)
	response, err := client.PrintPlaylist(context.Background(), &api.Playlist{Name: playlistName})
	if err != nil {
		fmt.Println("Error:", err)
	}
	if response == nil {
		fmt.Println("Error:", "not found")
	} else {
		for i, e := range response.Songs {
			fmt.Printf("№%d. Song: %s\n", i+1, e)
		}
	}
}

func init() {
	rootCmd.AddCommand(printPlaylistCmd)
	printPlaylistCmd.Flags().StringVar(&playlistName, "playlistName", "unknown", "")
}
