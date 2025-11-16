package gstorage

import (
	"context"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/ysfgrl/gcore/gconf"
	"github.com/ysfgrl/gcore/gerror"
)

type IStorage interface {
	Init(conf *gconf.Storage)
	GetInfo(ctx context.Context, key string) (minio.ObjectInfo, *gerror.Error)
	CopyFromKey(ctx context.Context, storage IStorage, key string) (minio.UploadInfo, *gerror.Error)
	GetBucket() string
	GetPrefix() string
	PubHeaderFile(ctx context.Context, fileHeader *multipart.FileHeader) (string, *gerror.Error)
	GetSignedUrl(ctx context.Context, key string, duration time.Duration) (*url.URL, *gerror.Error)
}
