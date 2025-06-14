package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"learn-go/v2/internal/common"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var _ UploadProvider = (*s3Provider)(nil)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secretKey  string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName, region, apiKey, secretKey, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secretKey:  secretKey,
		domain:     domain,
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(provider.apiKey, provider.secretKey, ""),
	})

	if err != nil {
		log.Fatalln(err)
	}

	provider.session = sess

	return provider
}

func (p *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	_, err := s3.New(p.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(p.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(http.DetectContentType(data)),
	})

	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", p.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}
