package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ikalkali/belajar-grpc/blog/blogpb"
	"google.golang.org/grpc"
)


func main() {

	fmt.Println("Hello I'm a client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("127.0.0.1:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	blog := &blogpb.Blog{
		AuthorId: "Bambang",
		Title: "My first blog",
		Content: "First blog content",
	}
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error : %v", err)
	}
	fmt.Printf("Blog has been created: %v", res)

	// Read blog
	fmt.Println("Reading the blog")

	res2, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "61396cc362471c69e96ba136"})
	if err2 != nil {
		fmt.Print("Error happened while reading:%v",err2)
	}
	fmt.Printf("Blog was read: %v", res2)
}

