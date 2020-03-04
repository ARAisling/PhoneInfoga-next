package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	assert := assert.New(t)

	t.Run("successResponse", func(t *testing.T) {
		t.Run("should return success response", func(t *testing.T) {
			result := successResponse()

			assert.IsType(JsonResponse{}, result)
			assert.Equal(result.Success, true, "they should be equal")
			assert.Equal(result.Error, "", "they should be equal")
		})

		t.Run("should return success response with custom message", func(t *testing.T) {
			result := successResponse("test")

			assert.IsType(JsonResponse{}, result)
			assert.Equal(result.Success, true, "they should be equal")
			assert.Equal(result.Error, "test", "they should be equal")
		})
	})

	t.Run("errorResponse", func(t *testing.T) {
		t.Run("should return error response", func(t *testing.T) {
			result := errorResponse()

			assert.IsType(JsonResponse{}, result)
			assert.Equal(result.Success, false, "they should be equal")
			assert.Equal(result.Error, "An error occurred", "they should be equal")
		})

		t.Run("should return error response with custom message", func(t *testing.T) {
			result := errorResponse("test")

			assert.IsType(JsonResponse{}, result)
			assert.Equal(result.Success, false, "they should be equal")
			assert.Equal(result.Error, "test", "they should be equal")
		})
	})
}
