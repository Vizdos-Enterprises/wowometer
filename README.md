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

## Loading configuration from environment variables

You can also create a Wowometer endpoint from environment variables using `WowometerFromEnv`.

This is useful when you do not want to hardcode Google Form IDs or field IDs in code.

```go
func StartHTTP() {
	http.Handle("/submit", wowometer_http.WowometerFromEnv(
		"InternalFormID",
		func(r *http.Request) (string, error) {
			return "test user id", nil
		},
		func(r *http.Request, rating wowometer_http.WowometerBody, forUserID string) {
			log.Printf("New rating from %s: %d", forUserID, rating.Rating)
		},
	))

	log.Printf("Starting HTTP server..")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}
```

Given the key `InternalFormID`, Wowometer will look for these environment variables:

```env
InternalFormID_ForAppName="DEMO_APP"
InternalFormID_FormID="google-form-id"
InternalFormID_FieldID_AppName="google-form-entry-id"
InternalFormID_FieldID_UserID="google-form-entry-id"
InternalFormID_FieldID_Rating="google-form-entry-id"
InternalFormID_FieldID_Feedback="google-form-entry-id"
```

These values map directly to the same fields used by `WowometerEndpoint`.

### Optional SkipSend mode

You can disable sending to Google Forms for a specific environment-based endpoint by setting:

```env
InternalFormID_SkipSend="true"
```

When `SkipSend` is `true`, Wowometer will not POST to Google Forms. Instead, it will print the submitted rating and feedback to the console.

This is useful for local development, staging, or testing `PostAction` without needing a real Google Form.

```env
InternalFormID_ForAppName="DEMO_APP"
InternalFormID_SkipSend="true"
```

When `SkipSend` is enabled, the Google Form-specific fields are optional:

```env
InternalFormID_FormID
InternalFormID_FieldID_AppName
InternalFormID_FieldID_UserID
InternalFormID_FieldID_Rating
InternalFormID_FieldID_Feedback
```

If `SkipSend` is not enabled, all Google Form fields are required. Missing fields will cause startup to fail with an error showing which environment variable is missing.

## Testing

If you would just like to test your PostAction, you can run with `-tags skip_wowometer_send` to disable POSTing to the Google Form. **It will still trigger the internal PostAction if specified**
