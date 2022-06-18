package imagestorage

type awsS3 struct{}

var _ ImageUploader = &awsS3{}

func NewS3Uploader() *awsS3 {
	return &awsS3{}
}

func (s3 *awsS3) Upload(data []byte, name string) (location string, err error) {
	return "hello world", nil
}
