# Wowometer

Used to easily collect ratings on a 1-5 scale, and optionally collect feedback. This system will automatically POST to a Google Form, and optionally run a PostAction function.

## Usage

You'll need to first create a Google Form.

Done?

Cool. Next, go to fill out a form, enable developer tools, preserve network logs, and then submit a test form. Use the field names as entries (recommended).

After that, search the Network Logs for those entries. You should see something in a POST body that resembles `entry.<id>`. Use those to attach in the correct spots below (JUST the id, not including entry.)

```
func StartHTTP() {
	http.Handle("/submit", wowometer_http.WowometerEndpoint{
		ForAppName: "Example App",
		FieldIDs: wowometer_http.WowometerFormEntryIDs{
			AppName:  "google-form-entry-id",
			UserID:   "google-form-entry-id",
			Rating:   "google-form-entry-id",
			Feedback: "google-form-entry-id",
		},
		FormID: "google-form-id",
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
```

## Use the data in your app

Optionally, you can use the PostAction parameter in the Wowometer structure to do what you wish with the info:

```
PostAction     func(r *http.Request, rating wowometerBody, forUserID string)
```
