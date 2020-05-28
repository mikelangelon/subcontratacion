package attachmentrepo

import (
	"io"
	"time"
)

// Store allows us to persist, retrieve and manage attachments.
type Store interface {
	Save(id string, meta UploadMeta, attachment io.ReadCloser) (Meta, error)
	ListAttachments(id string) ([]Meta, error)
}

// Meta holds metadata about an attachment
type Meta struct {
	Name         string    // The name of the file as it will appear to the user
	Description  string    // A user defined description
	LastModified time.Time // When the attachment was save/last modified.
	MD5Hex       string    // The hex encoded MD5 sum of the attachment file contents
	SizeBytes    uint64    // The size in bytes of the attachment
	OfferID      string
}

type UploadMeta struct {
	Name        string // The name of the file as it will appear to the user
	Description string // A user defined description
	OfferID     string //
}
