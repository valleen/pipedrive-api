package integration

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	"github.com/valleen/pipedrive-api/pipedrive"
)

func TestPersonsMailMessagesService_List(t *testing.T) {
	result, _, err := client.Persons.ListMailMessages(context.Background(), 1, 0)
	if err != nil {
		t.Fatalf("Could not get person's mail messages: %v", err)
	}

	if result.Success != true {
		t.Fatal("Unsuccessful person's mail messages request")
	}

	expectedCurrency := pipedrive.MailMessage{
		ID: 10,
	}

	// if diff := deep.Equal(expectedCurrency, result.Data[0]); diff != nil {
	// 	t.Error(diff)
	// }
	if diff := deep.Equal(expectedCurrency.ID, result.Data[len(result.Data)-1].Data.ID); diff != nil {
		t.Error(diff)
	}
}
