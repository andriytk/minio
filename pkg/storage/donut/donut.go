package donut

import "io"

// INTERFACES

// Donut interface
type Donut interface {
	CreateBucket(bucket string) error
	ListBuckets() ([]string, error)
	GetObjectWriter(bucket, object string) (ObjectWriter, error)
	GetObject(bucket, object string) (io.ReadCloser, error)
	GetObjectMetadata(bucket, object string) (map[string]string, error)
	//	GetObjectWriter(bucket, object string) (ObjectWriter, error)
	//	ListObjects(bucket string) ([]string, error)
}

// Bucket interface
type Bucket interface {
	GetNodes() ([]string, error)
}

// Node interface
type Node interface {
	GetBuckets() ([]string, error)
	GetWriter(bucket, object string) (Writer, error)
	GetReader(bucket, object string) (io.ReadCloser, error)
	GetMetadata(bucket, object string) (map[string]string, error)
	GetDonutMetadata(bucket, object string) (map[string]string, error)
}

// ObjectWriter interface
type ObjectWriter interface {
	Write([]byte) (int, error)
	Close() error
	CloseWithError(error) error
	SetMetadata(map[string]string) error
	GetMetadata() (map[string]string, error)
}

// Writer interface
type Writer interface {
	ObjectWriter

	GetDonutMetadata() (map[string]string, error)
	SetDonutMetadata(map[string]string) error
}
