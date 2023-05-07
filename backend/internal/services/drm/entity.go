package drm

const SubscriptionPrice = 100

type Proof struct {
	V     int64  `json:"v"`
	R     string `json:"r"`
	S     string `json:"s"`
	Hash  string `json:"hash"`
	Date  string `json:"date"`
	Value int64  `json:"value"`
}
