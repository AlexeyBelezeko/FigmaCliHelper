package figma

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Client figma api client.
type Client struct {
	token string
	baseURL string
	httpClient *http.Client
}

func NewClient(token string, baseURL string, httpClient *http.Client) *Client {
	return &Client{
		token:      token,
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

// CountComments returns map of comments count with user name as key.
func (c *Client) CountComments(fileKey string) (map[string]int, error) {
	comments, err := c.getAllComments(fileKey)
	if err != nil {
		return nil, err
	}
	result := make(map[string]int, 0)
	for _, comment := range comments {
		counter, ok := result[comment.User.Handle]
		if !ok {
			result[comment.User.Handle] = 1
		} else {
			result[comment.User.Handle] = counter + 1
		}
	}
	return result, nil
}

// DeleteAllComments deletes all userName comments from file fileKey and returns count of deleted comments.
func (c *Client) DeleteAllComments(fileKey, userName string) (int, error) {
	comments, err := c.getAllComments(fileKey)
	if err != nil {
		return 0, err
	}
	counter := 0
	for _, comment := range comments {
		if comment.User.Handle == userName {
			err := c.deleteComment(fileKey, comment.Id)
			if err != nil {
				return 0, fmt.Errorf("failed to delete comment %s: %w", comment.Id, err)
			}
			counter++
		} else {
			log.Printf("comment will be skipped cause his author %s", comment.User.Handle)
		}
	}
	return counter, nil
}

type commentsResponse struct {
	Comments []Comment `json:"comments"`
}

func (c *Client) getAllComments(fileKey string) ([]Comment, error) {
	request, err := http.NewRequest("GET", c.baseURL + "v1/files/" + fileKey + "/comments", nil)
	if err != nil {
		return nil, err
	}
	request.Header["X-FIGMA-TOKEN"] = []string{c.token}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	var comments commentsResponse
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&comments)
	if err != nil {
		return nil, err
	}
	return comments.Comments, nil
}

func (c *Client) deleteComment(fileKey, commentID string) error {
	request, err := http.NewRequest("DELETE", c.baseURL + "v1/files/" + fileKey + "/comments/" + commentID , nil)
	if err != nil {
		return err
	}
	request.Header["X-FIGMA-TOKEN"] = []string{c.token}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("delete comemnt response status code is %d", response.StatusCode))
	}
	return nil
}