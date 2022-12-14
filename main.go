package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func main() {
    lambda.Start(delicioushareappGetUserPostDetail)
}

type post struct {
    Id string `json:"postId"`
}

type eatingPlace struct {
    Name string `json:"name"`
    Address string `json:"address"`
    Website string `json:"website"`
    Id string `json:"id"`
}

type contributor struct {
    UserId string `json:"userId"`
}

type postDetail struct {
    LargeImageUrl string `json:"largeImageUrl"`
    EatingPlace eatingPlace `json:"eatingPlace"`
    Contributor contributor `json:"contributor"`
    PostedTime string `json:"postedTime"`
}

func delicioushareappGetUserPostDetail(post post) (postDetail, error) {
    db := dynamo.New(
        session.New(),
        &aws.Config{
            Region: aws.String(os.Getenv("DYNAMO_REGION")),
        },
    )
    table := db.Table(os.Getenv("TABLE_NAME"))
    var detail postDetail
    err := table.Get("PostId", post.Id).One(&detail)
    return detail, err
}
