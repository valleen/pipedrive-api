package pipedrive

import "time"

type MailMessage struct {
	ID   int `json:"id"`
	From []struct {
		ID                 int    `json:"id"`
		EmailAddress       string `json:"email_address"`
		Name               string `json:"name"`
		LinkedPersonID     any    `json:"linked_person_id"`
		LinkedPersonName   any    `json:"linked_person_name"`
		MailMessagePartyID int    `json:"mail_message_party_id"`
	} `json:"from"`
	To []struct {
		ID                 int    `json:"id"`
		EmailAddress       string `json:"email_address"`
		Name               string `json:"name"`
		LinkedPersonID     int    `json:"linked_person_id"`
		LinkedPersonName   string `json:"linked_person_name"`
		MailMessagePartyID int    `json:"mail_message_party_id"`
	} `json:"to"`
	Cc                          []any     `json:"cc"`
	Bcc                         []any     `json:"bcc"`
	BodyURL                     string    `json:"body_url"`
	AccountID                   string    `json:"account_id"`
	UserID                      int       `json:"user_id"`
	MailThreadID                int       `json:"mail_thread_id"`
	Subject                     string    `json:"subject"`
	Snippet                     string    `json:"snippet"`
	MailTrackingStatus          any       `json:"mail_tracking_status"`
	MailLinkTrackingEnabledFlag int       `json:"mail_link_tracking_enabled_flag"`
	ReadFlag                    int       `json:"read_flag"`
	Draft                       any       `json:"draft"`
	DraftFlag                   int       `json:"draft_flag"`
	SyncedFlag                  int       `json:"synced_flag"`
	DeletedFlag                 int       `json:"deleted_flag"`
	ExternalDeletedFlag         int       `json:"external_deleted_flag"`
	ExpungedFlag                int       `json:"expunged_flag"`
	HasBodyFlag                 int       `json:"has_body_flag"`
	SentFlag                    int       `json:"sent_flag"`
	SentFromPipedriveFlag       int       `json:"sent_from_pipedrive_flag"`
	SmartBccFlag                int       `json:"smart_bcc_flag"`
	MessageTime                 time.Time `json:"message_time"`
	AddTime                     time.Time `json:"add_time"`
	UpdateTime                  time.Time `json:"update_time"`
	HasAttachmentsFlag          int       `json:"has_attachments_flag"`
	HasInlineAttachmentsFlag    int       `json:"has_inline_attachments_flag"`
	HasRealAttachmentsFlag      int       `json:"has_real_attachments_flag"`
	LastRepliedAt               time.Time `json:"last_replied_at"`
}

type MailMessageResponse struct {
	Success bool `json:"success"`
	Data    []struct {
		Data MailMessage `json:"data"`
	} `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}
