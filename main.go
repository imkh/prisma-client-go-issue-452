package main

import (
	"context"
	"encoding/json"

	"github.com/imkh/prisma-client-go-issue-452/prisma/db"
)

type PostInfo struct {
	Content string `json:"content"`
}

func main() {
	// Connect to database
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	ctx := context.Background()

	// Old post info
	oldPostInfo := &PostInfo{
		Content: "Old content",
	}
	oldPostInfoBytes, err := json.Marshal(oldPostInfo)
	if err != nil {
		panic(err)
	}

	// Create Post
	createdPost, err := client.Post.CreateOne(
		db.Post.Title.Set("Post title"),
		db.Post.Info.Set(oldPostInfoBytes),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}

	// New post info
	newPostInfo := &PostInfo{
		Content: "New content",
	}
	newPostInfoBytes, err := json.Marshal(newPostInfo)
	if err != nil {
		panic(err)
	}

	// Update post
	_, err = client.Post.FindUnique(
		db.Post.ID.Equals(createdPost.ID),
	).Update(
		db.Post.Info.Set(newPostInfoBytes),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
