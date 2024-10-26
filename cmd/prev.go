package cmd

import (
	"context"
	"fmt"
	"github.com/gremislaw/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var prevCmd = &cobra.Command{
	Use:   "prev",
	Short: "play prev song in playlist",
	Long:  ``,
	Run:   prev,
}

func prev(cmd *cobra.Command, args []string) {
	host, port := getHostPort()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)
	response, err := client.Prev(context.Background(), &api.Empty{})
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(response.Response)
}

func init() {
	rootCmd.AddCommand(prevCmd)
}
