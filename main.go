package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const POSTS_URL = "https://jsonplaceholder.typicode.com/posts"

func getCommentsUrl(postId int) string {
	return fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d/comments", postId)
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	posts, err := getPosts()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total %d posts are received\n", len(posts))

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	totalComments := 0
	for _, post := range posts {
		wg.Add(1)
		go getComments(post.Id, &totalComments, wg, mutex)
	}

	wg.Wait()
	fmt.Printf("Total %d comments are received\n", totalComments)
}

func getPosts() ([]Post, error) {
	response, err := http.Get(POSTS_URL)
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var posts []Post
	err = json.Unmarshal(responseData, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func getComments(postId int, totalComments *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	response, err := http.Get(getCommentsUrl(postId))
	if err != nil {
		fmt.Printf("Error while getting comments for post %d\n", postId)
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error while getting comments for post %d\n", postId)
		return
	}
	var comments []Comment
	err = json.Unmarshal(responseData, &comments)
	if err != nil {
		fmt.Printf("Error while getting comments for post %d\n", postId)
		return
	}

	mutex.Lock()
	*totalComments = *totalComments + len(comments)
	mutex.Unlock()
	fmt.Printf("Total %d comments are received for post %d\n", len(comments), postId)
}
