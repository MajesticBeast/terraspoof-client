package terraspoof

import (
	"encoding/json"
	"net/http"
)

type Bucket struct {
	ID               string
	Name             string
	Tags             string
	BucketDomainName string
}

// GetS3Bucket - Get S3 Bucket
func (c *Client) GetS3Bucket(bucketName string) (Bucket, error) {
	req, err := http.NewRequest("GET", c.HostURL+"/s3/get", nil)
	if err != nil {
		return Bucket{}, err
	}

	res, err := c.doRequest(req, nil)
	if err != nil {
		return Bucket{}, err
	}

	bucket := Bucket{}
	err = json.Unmarshal(res, &bucket)
	if err != nil {
		return Bucket{}, err
	}

	return bucket, nil
}

// CreateS3Bucket - Create S3 Bucket
func (c *Client) CreateS3Bucket(bucketName string) (Bucket, error) {
	req, err := http.NewRequest("POST", c.HostURL+"/s3/create", nil)
	if err != nil {
		return Bucket{}, err
	}

	res, err := c.doRequest(req, nil)
	if err != nil {
		return Bucket{}, err
	}

	bucket := Bucket{}
	err = json.Unmarshal(res, &bucket)
	if err != nil {
		return Bucket{}, err
	}

	return bucket, nil
}

// DeleteS3Bucket - Delete S3 Bucket
func (c *Client) DeleteS3Bucket(bucketName string) (Bucket, error) {
	req, err := http.NewRequest("DELETE", c.HostURL+"/s3/delete", nil)
	if err != nil {
		return Bucket{}, err
	}

	res, err := c.doRequest(req, nil)
	if err != nil {
		return Bucket{}, err
	}

	bucket := Bucket{}
	err = json.Unmarshal(res, &bucket)
	if err != nil {
		return Bucket{}, err
	}

	return bucket, nil
}
