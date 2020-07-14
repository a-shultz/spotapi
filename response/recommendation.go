package response

type Recommendations struct {
	seeds []RecommendationSeed
	tracks []TrackSimple
}

type RecommendationSeed struct {
	afterFilteringSize int
	afterRelinkingSize int
	href string
	id string
	initialPoolSize int
	entityType string // artist, track, genre
}
