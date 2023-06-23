package integration

import (
	"context"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/valleen/pipedrive-api/pipedrive"
)

func TestEmailTemplatesService_List(t *testing.T) {
	result, _, err := client.MailTemplates.List(context.Background())
	if err != nil {
		t.Errorf("Could not get email templates: %v", err)
	}

	if result.Success != true {
		t.Error("Unsuccessful email templates request")
	}

	expectedCurrency := pipedrive.MailTemplate{
		ID:                1,
		Name:              "Follow up",
		Content:           "",
		UserID:            17403184,
		AddTime:           time.Date(2023, 6, 7, 7, 56, 30, 0, time.UTC),
		UpdateTime:        time.Date(2023, 6, 7, 7, 56, 30, 0, time.UTC),
		DeletedFlag:       0,
		SharedFlag:        0,
		OrderNr:           1,
		Subject:           "",
		DefaultTemplateID: "template_1",
	}

	if diff := deep.Equal(expectedCurrency, result.Data[0]); diff != nil {
		t.Error(diff)
	}
}

func TestEmailTemplatesService_Get(t *testing.T) {
	result, _, err := client.MailTemplates.Get(context.Background(), 6, "fr")
	if err != nil {
		t.Errorf("Could not get email template: %v", err)
	}

	if result.Success != true {
		t.Error("Unsuccessful email template request")
	}

	expectedCurrency := pipedrive.MailTemplate{
		ID:                6,
		Name:              "test",
		Content:           "<div class=\"pipe-mailbody pipe-mailbody-a0b2609b-2627-42d8-a7e9-fe256f71cee9\">test</div>",
		UserID:            17403184,
		AddTime:           time.Date(2023, 6, 16, 16, 22, 29, 0, time.UTC),
		UpdateTime:        time.Date(2023, 6, 16, 16, 22, 29, 0, time.UTC),
		DeletedFlag:       0,
		SharedFlag:        0,
		OrderNr:           6,
		Subject:           "test",
		DefaultTemplateID: "",
	}

	if diff := deep.Equal(expectedCurrency, result.Data); diff != nil {
		t.Error(diff)
	}
}
