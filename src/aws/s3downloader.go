// author: Umesh Patil
// All copyrights reserved with NeoSemantix, Inc.
// April 2020
//
// The program download email objects received on a bucket and
// deletes those from the bucket upon successful download.
//
// Appropriate credentials are needed in ~/.aws directory.
// Contents of the file credentials are of the format:
//
//	[default]
//	aws_access_key_id = <your access key id>
//	aws_secret_access_key = <AWS generated access key>
//

package main

import (
	"../util"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
	"path/filepath"
)

type BucketToFetch struct {
	AwsBucketName			string
	LocalDestinationDir		string
}

type BucketList struct {
	Bl	[]BucketToFetch
}

type downloader struct {
	*s3manager.Downloader
	bucket, dir string
	svc *session.Session
}

var (
	Prefix         = ""    // Using this key prefix
)

var bucketList BucketList


func main() {

	// Read the configuration - meaning AWS bucket names and local destination
	err := util.ReadCfg(&bucketList, "s3downloader.json")
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatal("Could not read bucket list")
	}

	// Invoke one by one each bucket
	for i := 0; i < len(bucketList.Bl); i++ {
		bucketName := bucketList.Bl[i].AwsBucketName
		handleBucket(bucketName, bucketList.Bl[i].LocalDestinationDir)
		fmt.Printf("Done %s\n", bucketName)
	}
}

// Download all accumulated emails / objects in the given bucket to the specified local directory.
func handleBucket(bckt string, ld string) {
	// The session the S3 Downloader will use
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))

	// Create a downloader with the session and default options
	manager := s3manager.NewDownloader(sess)

	d := downloader{bucket: bckt, dir: ld, Downloader: manager, svc: sess}

	client := s3.New(session.New(&aws.Config{
		Region: aws.String("us-east-1"),
		CredentialsChainVerboseErrors: aws.Bool(true),
		LogLevel: aws.LogLevel(aws.LogDebug)},
	))
	params := &s3.ListObjectsInput{Bucket: &bckt, Prefix: &Prefix}
	err := client.ListObjectsPages(params, d.eachPage)
	fmt.Println(err)
}

// Get page to download every object in the page
func (d *downloader) eachPage(page *s3.ListObjectsOutput, more bool) bool {
	for _, obj := range page.Contents {
		fmt.Printf("Object: %s modified on: %v\n", *obj.Key, *obj.LastModified)
		d.downloadToFile(*obj.Key)
	}
	return true
}

// Download the given file.
func (d *downloader) downloadToFile(key string) {
	// Create the directories in the path
	file := filepath.Join(d.dir, key+".eml")
	if err := os.MkdirAll(filepath.Dir(file), 0775); err != nil {
		panic(err)
	}

	// Set up the local file
	fd, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	// Download the file using the AWS SDK for Go
	fmt.Printf("Downloading s3://%s/%s to %s...\n", d.bucket, key, file)
	params := &s3.GetObjectInput{Bucket: &d.bucket, Key: &key}
	numBytes, err := d.Download(fd, params)
	if err != nil {
		fmt.Printf("Error downloading file %s: %v\n", key, err)
	} else {
		fmt.Printf("Downloaded %d bytes, error: %v\n", numBytes, err)
		delete(d.bucket, key)
	}
}

// Delete an object of given key in the given bucket
func delete(bucket string, key string) {
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String("us-east-1")}))
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	_, err := svc.DeleteObject(input)
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
		return
	}
	fmt.Printf("Deleted %s successfully\n", key)
}