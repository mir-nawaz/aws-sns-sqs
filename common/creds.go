package common

import "os"

type Creds struct {
	AccessKey string
	SecretKey string
	Region    string
	Token     string
}

func GetCredentials() Creds {
	return Creds{
		AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		SecretKey: os.Getenv("AWS_SECRET_KEY"),
		Region:    os.Getenv("AWS_REGION"),
		Token:     os.Getenv("AWS_TOKEN"),
	}
}
