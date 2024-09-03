package tokenbucket_test

import (
	"testing"

	"github.com/mukhtar-husnain/rate-limiter/tokenbucket"
	"github.com/stretchr/testify/assert"
)

func TestCreateTokenBucketRateLimiter(t *testing.T) {
	badBucket, err := tokenbucket.NewBucket(0, 86400, 1000)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), tokenbucket.INVALID_MAX_AMOUNT)
	assert.Nil(t, badBucket)

	badBucket, err = tokenbucket.NewBucket(-10, 86400, 1000)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), tokenbucket.INVALID_MAX_AMOUNT)
	assert.Nil(t, badBucket)

	badBucket, err = tokenbucket.NewBucket(1000, 0, 1000)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), tokenbucket.INVALID_REFILL_TIME)
	assert.Nil(t, badBucket)

	badBucket, err = tokenbucket.NewBucket(1000, -1, 1000)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), tokenbucket.INVALID_REFILL_TIME)
	assert.Nil(t, badBucket)

	badBucket, err = tokenbucket.NewBucket(100, 3600e9, 1000)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), tokenbucket.INVALID_REFILL_AMOUNT)
	assert.Nil(t, badBucket)

	badBucket, err = tokenbucket.NewBucket(100, 3600e9, 0)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), tokenbucket.INVALID_REFILL_AMOUNT)
	assert.Nil(t, badBucket)

	badBucket, err = tokenbucket.NewBucket(100, 3600e9, -1000)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), tokenbucket.INVALID_REFILL_AMOUNT)
	assert.Nil(t, badBucket)

	//Correct token bucket
	tb, err := tokenbucket.NewBucket(1000, 86400, 1000)
	assert.Nil(t, err)
	assert.NotNil(t, tb)
}

func TestAllowRequest(t *testing.T) {
	tokenBucketSize := int64(10)
	tb, err := tokenbucket.NewBucket(tokenBucketSize, 86400, tokenBucketSize)
	assert.Nil(t, err)
	assert.NotNil(t, tb)

	for i := 0; i < int(tokenBucketSize) ; i++ {
		check := tb.AllowRequest() 
		assert.Equal(t, true, check)
	} 
	assert.Equal(t, false, tb.AllowRequest())
}