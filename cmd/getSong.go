package cmd

import (
	"context"
	"fmt"
	"github.com/gremislaw/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var getSongCmd = &cobra.Command{
	Use:   "getSong",
	Short: "get specific song by name",
	Long: ``,
	Run: getSong,
}

func getSong(cmd *cobra.Command, args []string) {
	host, port := getHostPort()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)
	response, err := client.GetSong(context.Background(), &api.Song{Name: songName})
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(response)
}

func init() {
	rootCmd.AddCommand(getSongCmd)
	getSongCmd.Flags().StringVar(&songName, "songName", "unknown", "")
}
