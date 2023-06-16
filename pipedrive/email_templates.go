package pipedrive

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// DealService handles mailbox related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/Mailbox
type EmailTemplateService service

// EmailTemplate represents a Pipedrive email template.
type EmailTemplate struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Content           string    `json:"content"`
	UserID            int       `json:"user_id"`
	AddTime           time.Time `json:"add_time"`
	UpdateTime        time.Time `json:"update_time"`
	DeletedFlag       int       `json:"deleted_flag"`
	SharedFlag        int       `json:"shared_flag"`
	OrderNr           int       `json:"order_nr"`
	Subject           string    `json:"subject"`
	DefaultTemplateID string    `json:"default_template_id"`
}

func (d EmailTemplate) String() string {
	return Stringify(d)
}

// EmailTemplatesResponse represents multiple deals response.
type EmailTemplatesResponse struct {
	Success bool            `json:"success,omitempty"`
	Data    []EmailTemplate `json:"data,omitempty"`
}

// EmailTemplateResponse represents single deal response.
type EmailTemplateResponse struct {
	Success bool          `json:"success,omitempty"`
	Data    EmailTemplate `json:"data,omitempty"`
}

type EmailTemplateFilterOptions struct {
	Language   string `url:"language"`
	StrictMode string `url:"strict_mode"`
}

// Get a email template.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons_id
func (s *EmailTemplateService) Get(ctx context.Context, emailTemplateID int, language string) (*EmailTemplateResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("/mailbox/mailTemplates/%d", emailTemplateID), &EmailTemplateFilterOptions{
		Language:   language,
		StrictMode: "true",
	}, nil)
	if err != nil {
		return nil, nil, err
	}

	var record *EmailTemplateResponse

	resp, err := s.client.Do(ctx, req, &record)
	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// List email templates.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals
func (s *EmailTemplateService) List(ctx context.Context) (*EmailTemplatesResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/mailbox/mailTemplates", nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var record *EmailTemplatesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// EmailTemplateCreateOptions specifices the optional parameters to the
// MailboxService.Create method.
type EmailTemplateCreateOptions struct {
	Title               string    `json:"title"`
	Value               string    `json:"value"`
	Currency            string    `json:"currency"`
	UserID              uint      `json:"user_id"`
	PersonID            uint      `json:"person_id"`
	OrgID               uint      `json:"org_id"`
	StageID             uint      `json:"stage_id"`
	Status              string    `json:"status"`
	Probability         uint      `json:"probability"`
	LostReason          string    `json:"lost_reason"`
	AddTime             Timestamp `json:"add_time"`
	VisibleTo           VisibleTo `json:"visible_to"`
	RequirementAnalysis string    `json:"56d3d40c37c0db60fff576ae73ba2fea0d58dc09"`
	WantedStartTime     Timestamp `json:"a3114acce61bb930180af173b395d76f42af8794"`
	TemporaryLink       string    `json:"4fe88fad67d8dcbc17d18d9ee1faac55122249fd,omitempty"`
	LeadSource          uint      `json:"5d4fbabc9b032aeb3df515d9c66994d6892ee062,omitempty"`
}

// Create a new email template.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals
func (s *EmailTemplateService) Create(ctx context.Context, opt *EmailTemplateCreateOptions) (*EmailTemplateResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/deals", nil, struct {
		Title               string    `json:"title"`
		Value               string    `json:"value"`
		Currency            string    `json:"currency"`
		UserID              uint      `json:"user_id"`
		PersonID            uint      `json:"person_id"`
		OrgID               uint      `json:"org_id"`
		StageID             uint      `json:"stage_id"`
		Status              string    `json:"status"`
		Probability         uint      `json:"probability"`
		LostReason          string    `json:"lost_reason"`
		AddTime             string    `json:"add_time"`
		VisibleTo           VisibleTo `json:"visible_to"`
		RequirementAnalysis string    `json:"56d3d40c37c0db60fff576ae73ba2fea0d58dc09"`
		WantedStartTime     string    `json:"a3114acce61bb930180af173b395d76f42af8794"`
		TemporaryLink       string    `json:"4fe88fad67d8dcbc17d18d9ee1faac55122249fd,omitempty"`
		LeadSource          uint      `json:"5d4fbabc9b032aeb3df515d9c66994d6892ee062,omitempty"`
	}{
		opt.Title,
		opt.Value,
		opt.Currency,
		opt.UserID,
		opt.PersonID,
		opt.OrgID,
		opt.StageID,
		opt.Status,
		opt.Probability,
		opt.LostReason,
		opt.AddTime.FormatFull(),
		opt.VisibleTo,
		opt.RequirementAnalysis,
		opt.WantedStartTime.Format(),
		opt.TemporaryLink,
		opt.LeadSource,
	})

	if err != nil {
		return nil, nil, err
	}

	var record *EmailTemplateResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
