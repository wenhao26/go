package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
)

var domain = "https://coinsky.s3.us-west-1.amazonaws.com/"
var id = "AKIA3DRJB3RGEERNBKG6"
var secret = "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"

func UploadBase64imgToS3(imageBase64, contentType, filename string) (url string) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.UsWest1RegionID),
		Credentials: credentials.NewStaticCredentials(id, secret, ""),
	})
	if err != nil {
		return "Err1:" + err.Error()
	}

	decodeStr := Base64Decode(imageBase64)
	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String("coinsky"),
		Key:         aws.String(filename),
		Body:        bytes.NewReader([]byte(decodeStr)),
		ContentType: aws.String(contentType),
		ACL:         aws.String("public-read"),
	})
	if err != nil {
		return "Err2:" + err.Error()
	}
	return domain + filename
}

func Base64Decode(str string) string {
	reader := strings.NewReader(str)
	decoder := base64.NewDecoder(base64.RawStdEncoding, reader)
	buf := make([]byte, 1024)
	dst := ""
	for {
		n, err := decoder.Read(buf)
		dst += string(buf[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return dst
}

func main() {
	imageDataUri := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEEAAABBCAIAAAABlV4SAAAABnRSTlMAAAAAAABupgeRAAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAf0lEQVRoge3ZQQ2AMBAAQYoyJCAJSUhAGhZ4NGEgOwLabu53Hcd5LR+3vv2ACWow1GCowVCDoQZDDYYaDDUYajDUYKjBUIPhDw1j7nHPN5/Hvs269A9zqMFQg6EGQw2GGgyjv11CDYYaDDUYajDUYKjBUIOhBkMNhhoMNRhqMNzRjwdg8MNWIgAAAABJRU5ErkJggg=="
	imageParts := strings.Split(imageDataUri, ";base64,")
	imageTypeAux := strings.Split(imageParts[0], "image/")
	imageType := imageTypeAux[1]
	imageBase64 := imageParts[1]
	contentType := "image/" + imageType
	filename := "test/test001." + imageType
	resp := UploadBase64imgToS3(imageBase64, contentType, filename)
	fmt.Println(resp)
}
