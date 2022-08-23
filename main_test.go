package main

import (
	"context"
	"testing"

	openapi "./pkg/lumosclient/lumosrestapiclient"
)

func TestGetPet(t *testing.T) {
	expectedPet := openapi.Pet{
		Dog: &openapi.Dog{Name: "Steve", Age: 32, Type: "Lizard", Bark: true},
	}
	api := openapi.DefaultApiService{}
	pet, resp, err := api.PetNameGet(context.Background(), expectedPet.Dog.GetName()).Execute()
	if resp.StatusCode != 200 {
		t.Fatal("Bad resp code")
	}
	if err != nil {
		t.Fatal("Request returned an error")
	}
	if pet != expectedPet {
		t.Fatal("Pets not equal")
	}
}
