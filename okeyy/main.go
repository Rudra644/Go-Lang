package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	masterPlaylistURL := "https://video.bsky.app/watch/did%3Aplc%3Auo2fna47c4v6zcnklxfhcvjb/bafkreidqnefttrnlft5jaxjqh7stfs3lzwqyekizhvqhn7cvxqqrjkyuzq/playlist.m3u8"
	baseURL := "https://video.bsky.app/watch/" // Default base URL

	// Fetch the master playlist
	resp, err := http.Get(masterPlaylistURL)
	if err != nil {
		fmt.Println("Failed to fetch master playlist:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the master playlist to find sub-playlists
	var selectedSubPlaylist string
	scanner := bufio.NewScanner(resp.Body)

	fmt.Println("Master playlist lines:")
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println(line) // Debugging: Print each line

		// Look for sub-playlists containing "720p"
		if strings.Contains(line, "720p") && strings.Contains(line, ".m3u8") {
			selectedSubPlaylist = line
			break
		}
	}

	if selectedSubPlaylist == "" {
		fmt.Println("720p sub-playlist not found in master playlist.")
		return
	}

	// Debug: Try different base URLs if needed
	subPlaylistURL := baseURL + selectedSubPlaylist
	fmt.Println("Trying sub-playlist URL:", subPlaylistURL)

	// Fetch the sub-playlist
	resp, err = http.Get(subPlaylistURL)
	if err != nil {
		fmt.Println("Failed to fetch sub-playlist:", err)
		return
	}
	defer resp.Body.Close()

	// Debug: Print the full content of the sub-playlist
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read sub-playlist content:", err)
		return
	}

	fmt.Println("Sub-playlist content:")
	fmt.Println(string(body))

	// Parse the sub-playlist to extract .ts files
	scanner = bufio.NewScanner(strings.NewReader(string(body)))
	fmt.Println("Video segments:")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasSuffix(line, ".ts") { // Identify .ts segment files
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading sub-playlist:", err)
	}
}
