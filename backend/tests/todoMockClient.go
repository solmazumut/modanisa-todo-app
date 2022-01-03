package tests

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Todo struct {
	Id    string
	Task  string
}

func NewClient(host string) Client {
	return Client{host}
}

type Client struct {
	host string
}

func (c Client) AddTodo(task string) ( error) {
	client := http.Client{}

	data := url.Values{}
	data.Set("task", task)

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/todo", c.host),
		strings.NewReader(data.Encode()),
	)
	req.Header.Add("Accept", "application/json")

	resp, _ := client.Do(req)

	if resp.StatusCode == http.StatusOK {
		return nil
	}


	return  errors.New(fmt.Sprintf("error"))
}
