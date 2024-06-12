package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Address struct {
	Street string
	Number int
}

type Post struct {
	Id int
	Title string
	Content string
}

func newPost(title, content string) (p Post){
	id, err := rand.Int(rand.Reader, big.NewInt(1000000))
	
	if err == nil {
		p.Id = int(id.Int64())
	}

	p.Title = title
	p.Content = content

	return p
}

func (p Post) showPost() {
	fmt.Printf("------ Post %d ------\n", p.Id)
	fmt.Println("Title: ", p.Title)
	fmt.Println("Content: ", p.Content)
}

type Client struct {
	Name string
	Age int
	IsActive bool
	Address
	Posts []Post
}

func (c *Client) newPost(post Post) {
	c.Posts = append(c.Posts, post)
}

func (c Client) allPosts () {
	for _, post := range c.Posts {
		post.showPost()
	}

}

func main() {

	user1 := Client {
		Name: "Odeassis",
		Age: 25,
		IsActive: true,
		Address: Address {
			Number: 10,
			Street: "Av. Gustavo",
		},
	}

	post1 := newPost("My first post", "This is my first post")
	post2 := newPost("My second post", "This is my second post")
	//user1.Posts = append(user1.Posts, newPost("My second post", "This is my second post"))
	
	user1.newPost(post1)
	user1.newPost(post2)

	user1.allPosts()
}