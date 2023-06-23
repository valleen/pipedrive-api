package pipedrive

import (
	"context"
	"net/http"
	"time"
)

// DealService handles mailbox related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/Mailbox
type MailAttachmentService service

type MailAttachment struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	Name           string    `json:"name"`
	FileName       string    `json:"file_name"`
	RemoteID       string    `json:"remote_id"`
	RemoteLocation string    `json:"remote_location"`
	S3Bucket       string    `json:"s3_bucket"`
	FileSize       int       `json:"file_size"`
	FileType       string    `json:"file_type"`
	AddTime        time.Time `json:"add_time"`
	UpdateTime     time.Time `json:"update_time"`
	InlineFlag     bool      `json:"inline_flag"`
	MailTemplateID int       `json:"mail_template_id"`
	URL            string    `json:"url"`
}

func (d MailAttachment) String() string {
	return Stringify(d)
}

type MailAttachmentsResponse struct {
	Success        bool             `json:"success"`
	Data           []MailAttachment `json:"data"`
	AdditionalData AdditionalData   `json:"additional_data"`
}

type MailAttachmentsFilterOptions struct {
	MailTemplateID string `url:"mail_template_ids"`
	StrictMode     string `url:"strict_mode"`
}

// List email templates attachments.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals
func (s *MailAttachmentService) List(ctx context.Context, mailTemplateID string) (*MailAttachmentsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/mailbox/mailAttachments", &MailAttachmentsFilterOptions{
		MailTemplateID: mailTemplateID,
		StrictMode:     "true",
	}, nil)
	if err != nil {
		return nil, nil, err
	}

	var record *MailAttachmentsResponse
	resp, err := s.client.Do(ctx, req, &record)
	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
