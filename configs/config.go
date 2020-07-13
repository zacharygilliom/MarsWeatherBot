package configs

func GetNasaCredentials() string {
	apiKey := "YMbaetD9A4IuqGT6YskPjzYR1WTb3mz8owdTaa9O"
	return apiKey
}

func GetTwitterCredentials() (string, string, string, string) {
	twitterConsumerApiKey := "rBix0Quh40OTSQyJgy3t3mYGb"
	twitterConsumerApiSecretKey := "FR3CtfrzSBY7ITJuGq8aghvi7qjvhTFKzakASZJ8ugkop4tw55"
	twitterAccessToken := "1264750796391972864-4gFu3B8s6ajeqanwNNmq15CKtVa1JT"
	twitterAccessTokenSecret := "o9aJOa9eKLAy8yefo7LhRAf0hV4qmEWjbhSiZrlIlYWv7"
	return twitterConsumerApiKey, twitterConsumerApiSecretKey, twitterAccessToken, twitterAccessTokenSecret
}
