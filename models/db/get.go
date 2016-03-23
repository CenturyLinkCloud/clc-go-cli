package db

type Get struct {
	SubscriptionId string `json:"-" URIParam:"yes" valid:"required"`
}
