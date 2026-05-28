package wowometer_http

import (
	"fmt"
	"net/http"
	"os"
)

func WowometerFromEnv(
	key string,
	discoverUserID func(r *http.Request) (string, error),
	postAction func(r *http.Request, rating WowometerBody, forUserID string),
) WowometerEndpoint {
	get := func(name string) string {
		fullKey := fmt.Sprintf("%s_%s", key, name)

		value := os.Getenv(fullKey)
		if value == "" {
			panic(fmt.Sprintf(
				"wowometer env missing: %s (prefix=%s)",
				fullKey,
				key,
			))
		}

		return value
	}

	getOptional := func(name string, fallback string) string {
		fullKey := fmt.Sprintf("%s_%s", key, name)

		value := os.Getenv(fullKey)
		if value == "" {
			return fallback
		}

		return value
	}

	skipSend := os.Getenv(fmt.Sprintf("%s_SkipSend", key)) == "true"

	if skipSend {
		return WowometerEndpoint{
			ForAppName: getOptional("ForAppName", "SKIPPED"),
			FormID:     getOptional("FormID", "SKIPPED"),
			FieldIDs: WowometerFormEntryIDs{
				AppName:  getOptional("FieldID_AppName", "SKIPPED"),
				UserID:   getOptional("FieldID_UserID", "SKIPPED"),
				Rating:   getOptional("FieldID_Rating", "SKIPPED"),
				Feedback: getOptional("FieldID_Feedback", "SKIPPED"),
			},
			DiscoverUserID: discoverUserID,
			PostAction:     postAction,
			SkipSend:       true,
		}
	}

	return WowometerEndpoint{
		ForAppName: get("ForAppName"),
		FormID:     get("FormID"),
		FieldIDs: WowometerFormEntryIDs{
			AppName:  get("FieldID_AppName"),
			UserID:   get("FieldID_UserID"),
			Rating:   get("FieldID_Rating"),
			Feedback: get("FieldID_Feedback"),
		},
		DiscoverUserID: discoverUserID,
		PostAction:     postAction,
		SkipSend:       false,
	}
}
