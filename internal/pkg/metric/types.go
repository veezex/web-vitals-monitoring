package metric

type Metric interface {
	GetName() string
	GetValue() float64
	GetDelta() float64
	GetAttribution() string
	GetRating() string
	GetUri() string
	GetClient() string
	GetID() string
}
