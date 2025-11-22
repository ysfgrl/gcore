package gstorage

import (
	"context"
	"mime/multipart"

	"net/url"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/ysfgrl/gcore/gerror"
)

type MinioBase struct {
	Client *minio.Client
	Bucket string
	Prefix string
}

var chars = []string{
	" ",
	"ş",
	"Ş",
	"ç",
	"Ç",
	"ö",
	"Ö",
	"ü",
	"Ü",
	"ğ",
	"Ğ",
	"?",
	",",
	"^",
	"'",
	"!",
}

func (b *MinioBase) exist(ctx context.Context) *gerror.Error {

	if !b.Client.IsOnline() {
		return gerror.StorageNotOnlineError
	}
	exist, err := b.Client.BucketExists(ctx, b.Bucket)
	if err != nil {
		return gerror.GetError(err)
	}
	if !exist {
		return b.create(ctx)
	}
	return nil
}

func (b *MinioBase) create(ctx context.Context) *gerror.Error {
	err := b.Client.MakeBucket(ctx, b.Bucket, minio.MakeBucketOptions{
		Region: "us-east-1",
	})
	if err != nil {
		return gerror.StorageBucketCreateError
	}
	return nil
}

func zip(a1, a2 []string) []string {
	r := make([]string, 2*len(a1))
	for i, e := range a1 {
		r[i*2] = e
		r[i*2+1] = a2[i]
	}
	return r
}

func (b *MinioBase) GetSignedUrl(ctx context.Context, key string, duration time.Duration) (*url.URL, *gerror.Error) {

	if strings.HasPrefix(key, "storage://") {
		key = strings.Replace(key, "storage://", "", 1)
	}
	reqParams := make(url.Values)
	u, err := b.Client.PresignedGetObject(ctx, b.Bucket, key, duration, reqParams)
	if err != nil {
		return nil, gerror.GetError(err)
	}
	return u, nil
}

func (b *MinioBase) GetBucket() string {
	return b.Bucket
}
func (b *MinioBase) GetPrefix() string {
	return b.Prefix
}
func (b *MinioBase) CopyFromKey(ctx context.Context, storage IStorage, key string) (minio.UploadInfo, *gerror.Error) {

	if strings.HasPrefix(key, "tmp://") {
		key = strings.Replace(key, "tmp://", "", 1)
	}
	if strings.HasPrefix(key, "storage://") {
		key = strings.Replace(key, "storage://", "", 1)
	}
	info, err := storage.GetInfo(ctx, key)
	if err != nil {
		return minio.UploadInfo{}, err
	}

	newName := strings.Split(info.Key, "/")
	name := b.Prefix + time.Now().Format("2006_01_02") + "/" + newName[len(newName)-1]
	uploadInfo, err1 := b.Client.CopyObject(ctx, minio.CopyDestOptions{
		Bucket: b.Bucket,
		Object: name,
	}, minio.CopySrcOptions{
		Bucket: storage.GetBucket(),
		Object: key,
	})
	if err1 != nil {
		return minio.UploadInfo{}, gerror.GetError(err1)
	}
	return uploadInfo, nil
}

func (b *MinioBase) GetInfo(ctx context.Context, key string) (minio.ObjectInfo, *gerror.Error) {

	if strings.HasPrefix(key, "tmp://") {
		key = strings.Replace(key, "tmp://", "", 1)
	}
	if strings.HasPrefix(key, "storage://") {
		key = strings.Replace(key, "storage://", "", 1)
	}
	info, err := b.Client.GetObject(ctx, b.Bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return minio.ObjectInfo{}, gerror.GetError(err)
	}
	stat, err := info.Stat()
	if err != nil {
		return minio.ObjectInfo{}, gerror.GetError(err)
	}
	if stat.Err != nil {
		return minio.ObjectInfo{}, gerror.GetError(stat.Err)
	}
	return stat, nil
}

func (b *MinioBase) DeleteByKey(ctx context.Context, key string) *gerror.Error {
	err := b.Client.RemoveObject(ctx, b.Bucket, key, minio.RemoveObjectOptions{
		ForceDelete: true,
	})
	if err != nil {
		return gerror.GetError(err)
	}
	return nil
}

func (b *MinioBase) PubHeaderFile(ctx context.Context, fileHeader *multipart.FileHeader) (string, *gerror.Error) {

	//if err := b.exist(ctx); err != nil {
	//	return "", err
	//}
	file, err := fileHeader.Open()
	if err != nil {
		return "", gerror.GetError(err)
	}
	newName := strings.NewReplacer(zip(chars, make([]string, len(chars)))...).Replace(fileHeader.Filename)

	name := b.Prefix + time.Now().Format("2006_01_02") + "/" + time.Now().Format("15_04") + "_" + newName
	cType := fileHeader.Header.Get("Content-Type")

	info, err := b.Client.PutObject(ctx,
		b.Bucket,
		name,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{
			ContentType: cType,
			Expires:     time.Now().UTC().Add(time.Second * 10),
		})
	if err != nil {
		return "", gerror.GetError(err)
	}
	return info.Key, nil
}
