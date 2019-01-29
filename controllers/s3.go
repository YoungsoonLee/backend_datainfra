package controllers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Controller ...
type S3Controller struct {
	BaseController
}

const (
	S3_REGION = "ap-northeast-2"
	S3_BUCKET = "ndc-pm-resources"
)

//Access key ID: AKIAIOSFODNN7EXAMPLE
//Secret access key: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

func (s *S3Controller) UploadImage() {
	fmt.Println("upload image ... ")

	//var Buf bytes.Buffer
	fileHeader, err := s.GetFiles("files")
	if err != nil {
		beego.Debug(err)
	}
	beego.Info(fileHeader)
	beego.Info(fileHeader[0].Filename)

	f, err := fileHeader[0].Open()
	//sess := session.Must(session.NewSession())
	//creds := credentials.NewSharedCredentials("", "default")
	creds := credentials.NewSharedCredentials("../conf/aws/credentials", "default")

	s3, err := session.NewSession(&aws.Config{
		Region:      aws.String(S3_REGION),
		Credentials: creds,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = AddFilesToS3(s3, fileHeader[0].Filename, fileHeader[0].Size, f)
	if err != nil {
		log.Fatal(err)
	}

}

func AddFilesToS3(s *session.Session, fileName string, size int64, r io.Reader) error {

	buffer := make([]byte, size)
	r.Read(buffer)

	rObj, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String("myfolder" + "/" + fileName),
		ACL:    aws.String("private"),
		Body:   bytes.NewReader(buffer),
		//ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	fmt.Println("rObj: ", rObj)
	fmt.Println("err: ", err)

	return err

}