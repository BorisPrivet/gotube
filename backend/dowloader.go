package main

import (
	"fmt"
	"io"

	"github.com/kkdai/youtube/v2"
)

func DownloadFromUrl(url string) (io.ReadCloser, string, error) {
	client := youtube.Client{}
	video, err := client.GetVideo(url)

	if err != nil {
		fmt.Println("client.GetVideo error")
		return nil, "", err
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])

	fmt.Println("DownloadFromUrl function completed")
	fmt.Println(video.Title)

	return stream, video.Title, nil
}
