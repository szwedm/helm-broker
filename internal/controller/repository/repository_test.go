package repository

import (
	"testing"

	"errors"

	"github.com/kyma-project/helm-broker/pkg/apis/addons/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	// Given
	tr := NewAddonsRepository("http://example.com/index.yaml")

	// Then
	assert.Equal(t, v1alpha1.RepositoryStatusReady, tr.Repository.Status)
}

func TestRepository_IsFailed(t *testing.T) {
	// Given
	tr := NewAddonsRepository("http://example.com/index.yaml")

	// When
	tr.Failed()

	// Then
	assert.True(t, tr.IsFailed())
}

func TestRepository_FetchingError(t *testing.T) {
	// Given
	tr := NewAddonsRepository("http://example.com/index.yaml")

	// When
	err := errors.New("bug")
	tr.FetchingError(err)

	// Then
	assert.Equal(t, v1alpha1.RepositoryStatusFailed, tr.Repository.Status)
	assert.Equal(t, v1alpha1.RepositoryURLFetchingError, tr.Repository.Reason)
	assert.Equal(t, "Fetching repository failed due to error: 'bug'", tr.Repository.Message)
}
