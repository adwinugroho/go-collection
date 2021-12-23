package config

import (
	"os"
)

var (
	AWS_ACCESS_ID  = os.Getenv("AWS_ACCESS_ID")
	AWS_SECRET_KEY = os.Getenv("AWS_SECRET_KEY")
	AWS_REGION     = os.Getenv("AWS_REGION")
	BUCKET_NAME    = os.Getenv("BUCKET_NAME")
)
