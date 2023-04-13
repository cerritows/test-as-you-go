package leveltwo

import (
	"net/http"
)

const DataSourceURL = "https://mydatasource.com/posts"

func GetData(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	processData(res)
	return nil

}

func processData(res *http.Response) {}
