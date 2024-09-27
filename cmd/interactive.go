package cmd

import (
	"context"
	"fmt"
	"github.com/gremislaw/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "interactive client",
	Long: ``,
	Run: interactive,
}

func interactive(cmd *cobra.Command, args []string) {
	var act string
	var response *api.Response = nil
	var err error

	host, port := getHostPort()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)

	for act != "stop" {
		scanner  := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)
		scanner.Scan()
		act = scanner.Text()
		if act == "play" {
			response, err = client.Play(context.Background(), &api.Empty{})
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if act == "pause" {
			response, err = client.Pause(context.Background(), &api.Empty{})
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if act == "prev" {
			response, err = client.Prev(context.Background(), &api.Empty{})
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if act == "next" {
			response, err = client.Next(context.Background(), &api.Empty{})
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if strings.Contains(act, "add") {
			params := strings.Split(act, " ")
			if (len(params) != 4 || params[0] != "add") {
				fmt.Println("can't add new song, wrong format. Use: add <songName> <songAuthor> <duration>")
				continue
			}
			dur, err := strconv.Atoi(params[3])
			if err != nil {
				fmt.Println("Error:", err)
			}
			song := &api.Song{Name: params[1], Author: params[2], Duration: int64(dur)}
			response, err = client.AddSong(context.Background(), song)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if strings.Contains(act, "delete") {
			params := strings.Split(act, " ")
			if (len(params) != 2 || params[0] != "delete") {
				fmt.Println("can't delete song, wrong format. Use: delete <songName>")
				continue
			}
			song := &api.Song{Name: params[1]}
			response, err = client.DeleteSong(context.Background(), song)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if strings.Contains(act, "getSong") {
			params := strings.Split(act, " ")
			if (len(params) != 2 || params[0] != "getSong") {
				fmt.Println("can't get song, wrong format. Use: getSong <songName>")
				continue
			}
			song := &api.Song{Name: params[1]}
			res, err := client.GetSong(context.Background(), song)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(res)
		} else if strings.Contains(act, "getPlaylist") {
			res, err := client.GetPlaylist(context.Background(), &api.Empty{})
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(res)
		} else if strings.Contains(act, "update") {
			params := strings.Split(act, " ")
			if (len(params) != 4 || params[0] != "update") {
				fmt.Println("can't update song, wrong format. Use: update <songName> <songAuthor> <duration>")
				continue
			}
			dur, err := strconv.Atoi(params[3])
			if err != nil {
				fmt.Println("Error:", err)
			}
			song := &api.Song{Name: params[1], Author: params[2], Duration: int64(dur)}
			response, err = client.UpdateSong(context.Background(), song)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
		fmt.Println(response.Response)
	}

}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}
