package subscribe

type SubscribeFormatter struct {
	Email string `json:"email"`
}

func SubscribeFormat(subscribe Subscribe) SubscribeFormatter {
	formatter := SubscribeFormatter{
		Email: subscribe.Email,
	}
	return formatter
}

func SubscriberFormaters(subscribe []Subscribe) []SubscribeFormatter {
	if len(subscribe) == 0 {
		return []SubscribeFormatter{}
	}
	var subscriber []SubscribeFormatter

	for _, subsubscribes := range subscribe {
		subscribes := SubscribeFormat(subsubscribes)
		subscriber = append(subscriber, subscribes)
	}
	return subscriber
}
