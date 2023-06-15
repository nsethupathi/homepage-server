package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func GetMusicUrl(c *gin.Context) {
	songName := c.Param("songName")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create S3 service client
	svc := s3.New(sess)

	songKey := fmt.Sprintf("music/%s.wav", songName)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("niv-homepage-bucket"),
		Key:    aws.String(songKey),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		exitErrorf("Failed to sign request", err)
	}

	c.IndentedJSON(http.StatusOK, urlStr)
}
