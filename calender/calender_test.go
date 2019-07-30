package service

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	apiKey    = os.Getenv("AMPLITUDE_API_KEY")
	eventType = os.Getenv("AMPLITUDE_EVENT_TYPE")
)

var _ = Describe("Upload event with invalid param", func() {

	event := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(event)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/event", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadEvent)
	handler.ServeHTTP(recorder, request)

	Describe("Upload Event", func() {
		Context("event", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Upload event with valid param", func() {

	os.Setenv("API_KEY", apiKey)
	event := EventArguments{EventType: eventType}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(event)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/event", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadEvent)
	handler.ServeHTTP(recorder, request)

	Describe("Upload Event", func() {
		Context("event", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
