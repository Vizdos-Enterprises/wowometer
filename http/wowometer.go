package wowometer_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type wowometerBody struct {
	Rating   int    `json:"rating"`
	Feedback string `json:"feedback"`
}

func (b wowometerBody) Validate() bool {
	if b.Rating <= 0 || b.Rating > 5 {
		return false
	}

	return true
}

type WowometerFormEntryIDs struct {
	AppName  string
	UserID   string
	Rating   string
	Feedback string
}

type WowometerEndpoint struct {
	ForAppName     string
	FieldIDs       WowometerFormEntryIDs
	FormID         string
	DiscoverUserID func(r *http.Request) (string, error)
	PostAction     func(r *http.Request, rating wowometerBody, forUserID string)
}

func (wow WowometerEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userID, err := wow.DiscoverUserID(r)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		// handle the error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var reviewBody wowometerBody
	err = json.Unmarshal(body, &reviewBody)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !reviewBody.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	formData := url.Values{
		fmt.Sprintf("entry.%s", wow.FieldIDs.UserID):   {userID},
		fmt.Sprintf("entry.%s", wow.FieldIDs.Rating):   {fmt.Sprintf("%d", reviewBody.Rating)},
		fmt.Sprintf("entry.%s", wow.FieldIDs.Feedback): {reviewBody.Feedback},
		fmt.Sprintf("entry.%s", wow.FieldIDs.AppName):  {wow.ForAppName},
	}

	payload := bytes.NewBufferString(formData.Encode())

	url := fmt.Sprintf("https://docs.google.com/forms/u/0/d/e/%s/formResponse", wow.FormID)
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Referer", fmt.Sprintf("https://docs.google.com/forms/d/e/%s/viewform", wow.FormID))
	req.Header.Add("Referrer-Policy", "strict-origin-when-cross-origin")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("NEW FALLBACK REVIEW (%s): %s %d %s", wow.ForAppName, userID, reviewBody.Rating, reviewBody.Feedback)
	}

	wow.PostAction(r, reviewBody, userID)
}
