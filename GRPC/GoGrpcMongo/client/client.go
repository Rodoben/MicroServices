package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"GoGrpcMongo/blogpb"

	"google.golang.org/grpc"
)

var blogID string

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error Connecting to server: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)
	CreateBlogClient(c)
	ReadBlogClient(c)
	UpdateBlogClient(c)
	DeleteBlogClient(c)
	ListBlogClient(c)
}

func CreateBlogClient(c blogpb.BlogServiceClient) {
	fmt.Println("Hi i'll create a blog")
	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			Author:  "fonalf",
			Title:   "Live in peace",
			Content: "Living in harmoney and",
		},
	}

	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling the Server: %v", err)
	}
	log.Printf("Blog Created with values: %v", res.GetBlog())
	blogID = res.GetBlog().GetId()
	fmt.Println("*******************************************")
	fmt.Println("Blog ID:", blogID)
}

func ReadBlogClient(c blogpb.BlogServiceClient) {
	fmt.Println("Read blog client")

	req := &blogpb.ReadBlogRequest{
		BlogId: "6186de86cc933e05d3a2b9c4",
	}
	req2 := &blogpb.ReadBlogRequest{
		BlogId: blogID,
	}
	_ = req2
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling the server: %v", err)
	}
	log.Printf("Response: %v", res)

}

func UpdateBlogClient(c blogpb.BlogServiceClient) {
	newBlog := &blogpb.Blog{
		Id:      blogID,
		Author:  "Changed Author",
		Title:   "Changed title",
		Content: "Content of the blog,with some awesome additions!",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v", updateErr)
	}
	fmt.Printf("Blog was Updated: %v\n", updateRes)

}

func DeleteBlogClient(c blogpb.BlogServiceClient) {
	deleteres, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})
	if deleteErr != nil {
		fmt.Printf("Error happened while deleting :%v\n", deleteErr)
	}
	fmt.Printf("Blog was deleted:%v\n", deleteres.GetBlogId())

}

func ListBlogClient(c blogpb.BlogServiceClient) {
	fmt.Println("Listing All the documents!!!!!!!!!!!")
	req := &blogpb.ListBlogRequest{}

	res, err := c.ListBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something Happend: %v", err)
		}
		fmt.Println(res.GetBlog())

	}
}
