package minioclient

import (
	"io"
	attachments "mikelangelon/m/v2/internal/app/attachmentrepo"
	"strings"
	"time"

	"github.com/minio/minio-go/v6"
	"github.com/pkg/errors"
)

const bucketAttachments = "offers"

var allBuckets = []string{bucketAttachments}

type repo struct {
	client *minio.Client
}

func New(client *minio.Client) *repo {
	return &repo{
		client: client,
	}
}

// Prepare creates the minio buckets needed by the store
func (s *repo) Prepare() error {
	for _, name := range allBuckets {
		exists, err := s.client.BucketExists(name)
		if err != nil {
			return errors.Wrapf(err, "problem checking if bucket '%s' exists", name)
		}
		if !exists {
			err := s.client.MakeBucket(name, "EUROPE-NORTH1")
			if err != nil {
				return errors.Wrapf(err, "bucket '%s'", name)
			}
		}
	}
	return nil
}

// SaveAttachment implements attachment.Store.SaveAttachment
func (s *repo) Save(id string, meta attachments.UploadMeta, attachment io.ReadCloser) (attachments.Meta, error) {
	if id == "" {
		return attachments.Meta{}, errors.New("id can't be empty")
	}
	if meta.Name == "" {
		return attachments.Meta{}, errors.New("name can't be empty")
	}
	_, err := s.client.PutObject(bucketAttachments, s.objectName(id, meta.Name), attachment, -1, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
		UserMetadata: map[string]string{
			"X-Attachment-Context-ID":  id,
			"X-Attachment-Description": meta.Description,
		},
	})
	defer attachment.Close()
	if err != nil {
		return attachments.Meta{}, errors.Wrap(err, "could not put object")
	}

	fullMeta := attachments.Meta{
		Name:         meta.Name,
		LastModified: time.Now(),
		MD5Hex:       "", // TODO: Calculate or obtain from store
		SizeBytes:    0,  // TODO: Obtain from store
		Description:  meta.Description,
	}

	return fullMeta, nil
}

// ListAttachments implements attachment.Store.ListAttachments
func (s *repo) ListAttachments(id string) (res []attachments.Meta, err error) {
	if id == "" {
		return nil, errors.New("id can't be empty")
	}
	res = make([]attachments.Meta, 0, 10)
	doneCh := make(chan struct{})
	defer close(doneCh)

	objectCh := s.client.ListObjectsV2(bucketAttachments, id+"/", false, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			return nil, errors.Wrap(object.Err, "object list error")
		}
		if object.Size < 0 {
			return nil, errors.Errorf("negative size (%d) found", object.Size)
		}

		statInfo, err := s.client.StatObject(bucketAttachments, object.Key, minio.StatObjectOptions{})
		if err != nil {
			return nil, errors.Wrapf(err, "could not stat object: %s", object.Key)
		}

		res = append(res, attachments.Meta{
			Name:         s.stripObjectName(object.Key, id),
			Description:  statInfo.UserMetadata["X-Attachment-Description"],
			LastModified: object.LastModified,
			MD5Hex:       object.ETag,
			SizeBytes:    uint64(object.Size),
		})
	}
	return
}

func (s *repo) stripObjectName(objName, id string) string {
	return strings.Replace(objName, id+"/", "", 1)
}

// objectName prefixes the context ID to the filename so that its easy to search in minio using prefix searching.
func (s *repo) objectName(id, filename string) string {
	return id + "/" + filename
}
