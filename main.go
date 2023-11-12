package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

func ollama() {
	godotenv.Load()
	value := os.Getenv("CENSYS")

	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Basic "+value).
		Get("https://search.censys.io/api/v2/hosts/search?q=services.http.response.body%3A%20%22Ollama%20is%20running%22&per_page=100&virtual_hosts=INCLUDE&sort=RELEVANCE")

	if err != nil {
		fmt.Println(err)
		return
	}

	var result ApiResponse
	json.Unmarshal([]byte(resp.Body()), &result)

	for i := 0; i < len(result.Result.Hits); i++ {
		var hit Hit = result.Result.Hits[i]

		var location Location = hit.Location

		fmt.Println(strings.Repeat("-", 80))

		fmt.Println(hit.IP)

		for i := 0; i < len(hit.Services); i++ {
			fmt.Print(fmt.Sprint(hit.Services[i].Port), " ", hit.Services[i].ServiceName)

			if i == len(hit.Services)-1 {
				fmt.Println("")
			} else {
				fmt.Print(", ")
			}
		}

		if len(hit.DNS.ReverseDNS.Names) != 0 {
			fmt.Println(hit.DNS.ReverseDNS.Names[0])
		}

		fmt.Println(location.Country, " (", location.Province, ", ", location.City+")")
	}

}

func main() {
	ollama()
}
