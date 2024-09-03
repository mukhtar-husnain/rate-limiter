package tokenbucket

import (
	"fmt"
	"time"

	U "github.com/mukhtar-husnain/rate-limiter/util"
)

const (
	INVALID_MAX_AMOUNT    string = "invalid_max_amount"
	INVALID_REFILL_AMOUNT string = "invalid_refill_amount"
	INVALID_REFILL_TIME   string = "invalid_refill_time"
)

type TokenBucket struct {
	Key           string `json:"key"`
	MaxAmount     int64  `json:"max_amount"`
	Value         int64  `json:"value"`
	RefillTime    int64  `json:"refill_time"`
	RefillAmount  int64  `json:"refill_amount"`
	LastUpdatedAt int64  `json:"last_updated_at"`
}

func NewBucket(maxAmount, refillTime, refillAmount int64) (*TokenBucket, error) {
	if maxAmount <= 0 {
		return nil, fmt.Errorf(INVALID_MAX_AMOUNT)
	}
	if refillAmount > maxAmount || refillAmount <= 0 {
		return nil, fmt.Errorf(INVALID_REFILL_AMOUNT)
	}
	if refillTime <= 0 {
		return nil, fmt.Errorf(INVALID_REFILL_TIME)
	}

	key := U.GetNewBucketKey()

	return &TokenBucket{
		Key:           key,
		MaxAmount:     maxAmount,
		Value:         maxAmount,
		RefillAmount:  refillAmount,
		RefillTime:    refillTime,
		LastUpdatedAt: time.Now().UnixNano(),
	}, nil
}

func (tb *TokenBucket) AllowRequest() bool {
	tb.Value--
	return tb.Value >= 0
}

func (tb *TokenBucket) RefillBucket() {
	currentTime := time.Now().UnixNano()
	refillRate := (currentTime - tb.LastUpdatedAt) / tb.RefillTime
	tb.Value = U.MinInt64(tb.MaxAmount, tb.Value+refillRate*tb.RefillAmount)
	tb.LastUpdatedAt = currentTime
}
