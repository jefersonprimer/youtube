package models

import (
	"time"
)

type Subscription struct {
	SubscriptionID   int       `json:"subscription_id" db:"subscription_id"`
	SubscriberUserID int       `json:"subscriber_user_id" db:"subscriber_user_id"`
	ChannelUserID    int       `json:"channel_user_id" db:"channel_user_id"`
	SubscribeDate    time.Time `json:"subscribe_date" db:"subscribe_date"`
}
