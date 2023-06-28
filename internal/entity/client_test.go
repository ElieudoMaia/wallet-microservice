package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "any@mail.com")
	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "any@mail.com", client.Email)
}

func TestNewClientEmptyName(t *testing.T) {
	client, err := NewClient("", "any@mail.com")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestNewClientEmptyEmail(t *testing.T) {
	client, err := NewClient("John Doe", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, err := NewClient("John Doe", "any@mail.com")
	if err != nil {
		t.Error(err)
	}
	err = client.Update("John Doe updated", "updated@mail.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe updated", client.Name)
	assert.Equal(t, "updated@mail.com", client.Email)
}

func TestUpdateClientEmptyName(t *testing.T) {
	client, err := NewClient("John Doe", "any@mail.com")
	if err != nil {
		t.Error(err)
	}
	err = client.Update("", "other@mail.com")
	assert.NotNil(t, err)
	assert.Equal(t, "name is required", err.Error())
}

func TestUpdateClientEmptyEmail(t *testing.T) {
	client, err := NewClient("John Doe", "any@mail.com")
	if err != nil {
		t.Error(err)
	}
	err = client.Update("John Doe", "")
	assert.NotNil(t, err)
	assert.Equal(t, "email is required", err.Error())
}
