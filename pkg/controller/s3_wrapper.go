package controller

import (
	"fmt"
	"log"
	"minio-api/api"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetNewS3Client(endpoint, accessKeyID, secretAccessKey string) *s3.S3 {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	s3Session, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatalf("cannot create a new session: %v", err)
	}
	return s3.New(s3Session)
}

func S3CreateBucket(s3Client *s3.S3, bucketName string) *api.ApiResp {
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}

	result, err := s3Client.CreateBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				return &api.ApiResp{Status: "failed", Message: fmt.Sprint(s3.ErrCodeBucketAlreadyExists, aerr.Error())}
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				return &api.ApiResp{Status: "failed", Message: fmt.Sprint(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())}
			default:
				return &api.ApiResp{Status: "failed", Message: fmt.Sprint(aerr.Error())}
			}
		} else {
			return &api.ApiResp{Status: "failed", Message: fmt.Sprint(err.Error())}
		}
	}

	fmt.Print(result)
	return &api.ApiResp{Status: "success", Message: fmt.Sprintf("bucket %v successfully created", bucketName)}
}

func S3ListBuckets(s3Client *s3.S3) *s3.ListBucketsOutput {
	input := &s3.ListBucketsInput{}

	result, err := s3Client.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return nil
	}
	return result
}

func S3ListObjects(s3Client *s3.S3, bucketName string) *s3.ListObjectsV2Output {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		// MaxKeys: aws.Int64(3), // default value is 1000
	}

	result, err := s3Client.ListObjectsV2(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return nil
	}

	return result
}

func S3PutBucketLifeCycle(s3Client *s3.S3, bucketName string) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	ruleID := fmt.Sprintf("%s-expire", bucketName)
	deleteMarkRuleID := "Delete Marker"
	input := &s3.PutBucketLifecycleConfigurationInput{
		Bucket: aws.String(bucketName),
		LifecycleConfiguration: &s3.BucketLifecycleConfiguration{
			Rules: []*s3.LifecycleRule{
				{
					Status: aws.String("Enabled"),
					Expiration: &s3.LifecycleExpiration{
						Days: aws.Int64(2),
					},
					ID: &ruleID,
					Filter: &s3.LifecycleRuleFilter{
						Prefix: aws.String(""),
					},
					NoncurrentVersionExpiration: &s3.NoncurrentVersionExpiration{
						NoncurrentDays: aws.Int64(2),
					},
				},
				{
					Status: aws.String("Enabled"),
					Expiration: &s3.LifecycleExpiration{
						ExpiredObjectDeleteMarker: aws.Bool(true),
					},
					ID: &deleteMarkRuleID,
					Filter: &s3.LifecycleRuleFilter{
						Prefix: aws.String(""),
					},
				},
			},
		},
	}
	fmt.Println("PutBucketLifecycle")
	result, err := s3Client.PutBucketLifecycleConfiguration(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return nil, err
	}
	return result, nil
}

func S3GetBucketLifecycle(s3Client *s3.S3, bucketName string) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	input := &s3.GetBucketLifecycleConfigurationInput{
		Bucket: &bucketName,
	}

	result, err := s3Client.GetBucketLifecycleConfiguration(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	return result, nil
}
