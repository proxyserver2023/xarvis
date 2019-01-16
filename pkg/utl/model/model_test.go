package model_test

import (
	"testing"

	"github.com/alamin-mahamud/gapi/pkg/utl/model"
	"github.com/ribice/gorsk/pkg/utl/mock"
)

func TestBeforeInsert(t *testing.T) {
	base := &model.Base{
		ID: 1,
	}
	base.BeforeInsert(nil)
	if base.CreatedAt.IsZero() {
		t.Error("CreatedAt was not changed")
	}
	if base.UpdatedAt.IsZero() {
		t.Error("UpdatedAt was not changed")
	}
}

func TestBeforeUpdate(t *testing.T) {
	base := &model.Base{
		ID:        1,
		CreatedAt: mock.TestTime(2000),
	}

	base.BeforeUpdate(nil)

	if base.UpdatedAt == mock.TestTime(2001) {
		t.Error("UpdatedAt was not changed")
	}
}
