package testing_http

import (
	"log"
	"net/http"

	wowometer_http "github.com/Vizdos-Enterprises/wowometer/http"
)

func StartHTTP() {
	http.Handle("/submit", wowometer_http.WowometerEndpoint{
		ForAppName: "Attendance",
		FieldIDs: wowometer_http.WowometerFormEntryIDs{
			AppName:  "1473614318",
			UserID:   "120048923",
			Rating:   "1639522657",
			Feedback: "124761547",
		},
		FormID: "1FAIpQLSc2i7W5ZBkT-v0L7CpRg-gVJm_rhc26IG5UZBQbi8O0-nvEfA",
		DiscoverUserID: func(r *http.Request) (string, error) {
			return "test user id", nil
		},
	})

	log.Printf("Starting HTTP server..")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}
