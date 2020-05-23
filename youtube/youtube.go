package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Response struct
type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

//Item struct
type Item struct {
	Kind  string `json:"kind"`
	ID    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

// Stats struct
type Stats struct {
	View        string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
}

// GetSubscribers func
func GetSubscribers() (Item, error) {

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	q := req.URL.Query()

	q.Add("key", "AIzaSyCNWTThbDNf_N8rHMhsyhUlrjIZHaGJ7e0")
	q.Add("id", "UCQNv-8L2eyFlAD7AC2djeVQ")
	q.Add("part", "statistics")

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	defer res.Body.Close()

	fmt.Println("Response Status: ", res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	return response.Items[0], nil
}

// AIzaSyCNWTThbDNf_N8rHMhsyhUlrjIZHaGJ7e0
