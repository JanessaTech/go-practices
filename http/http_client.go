package http

import (
	"bufio"
	"fmt"
	"net/http"
)

func Test_http_client() {
	response, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("response status:", response.Status)

	scanner := bufio.NewScanner(response.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
