package calender

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
	key            = os.Getenv("GOOGLE_CALENDER_KEY")
	calenderID     = os.Getenv("GOOGLE_CALENDER_ID")
	eventID        = os.Getenv("GOOGLE_CALENDER_EVENT_ID")
	eventStartDate = os.Getenv("GOOGLE_CALENDER_EVENT_START_DATE")
	eventEndDate   = os.Getenv("GOOGLE_CALENDER_EVENT_END_DATE")
)

var _ = Describe("Get Event List with invalid param", func() {

	os.Setenv("KEY", key)
	calender := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/eventList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(EventList)
	handler.ServeHTTP(recorder, request)

	Describe("List Event", func() {
		Context("list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Event List with invalid key", func() {

	os.Setenv("KEY", "mockKey")
	calender := CalenderArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/eventList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(EventList)
	handler.ServeHTTP(recorder, request)

	Describe("List Event", func() {
		Context("list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Event List with valid param", func() {

	os.Setenv("KEY", key)
	calender := CalenderArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/eventList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(EventList)
	handler.ServeHTTP(recorder, request)

	Describe("List Event", func() {
		Context("list", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Get Event by ID with invalid param", func() {

	os.Setenv("KEY", key)
	calender := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEventByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Event", func() {
		Context("event by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Event by ID with invalid key", func() {

	os.Setenv("KEY", "mockKey")
	calender := CalenderArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEventByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Event", func() {
		Context("event by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Event by ID with invalid param", func() {

	os.Setenv("KEY", key)
	calender := CalenderArguments{CalenderID: "mockCalenderID", EventID: "mockEventID"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEventByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Event", func() {
		Context("event by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Event by ID with valid param", func() {

	os.Setenv("KEY", key)
	calender := CalenderArguments{CalenderID: calenderID, EventID: eventID}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEventByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Event", func() {
		Context("event by id", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Get Calender By ID with invalid param", func() {

	os.Setenv("KEY", key)
	calender := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getCalender", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCalenderByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Calender", func() {
		Context("Calender by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Calender By ID with invalid key", func() {

	os.Setenv("KEY", "mockKey")
	calender := CalenderArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getCalender", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCalenderByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Calender", func() {
		Context("Calender by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Calender By ID with invalid param", func() {

	os.Setenv("KEY", key)
	calender := CalenderArguments{CalenderID: "mockCalenderID"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getCalender", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCalenderByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Calender", func() {
		Context("Calender by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Calender by ID with valid param", func() {

	os.Setenv("KEY", key)
	calender := CalenderArguments{CalenderID: calenderID}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getCalender", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCalenderByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Calender", func() {
		Context("Calender by id", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Create Calender with invalid param", func() {

	os.Setenv("KEY", key)
	calender := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createCalender", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCalender)
	handler.ServeHTTP(recorder, request)

	Describe("Create Calender", func() {
		Context("create calender", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Calender List with invalid key", func() {

	os.Setenv("KEY", "mockKey")
	calender := CalenderArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createCalender", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCalender)
	handler.ServeHTTP(recorder, request)

	Describe("Create Calender", func() {
		Context("create calender", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create Calender with invalid param", func() {

	os.Setenv("KEY", key)
	calender := CalenderArguments{EventID: "mockCalender"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createCalender", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCalender)
	handler.ServeHTTP(recorder, request)

	Describe("Create Calender", func() {
		Context("create calender", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create Calender with valid param", func() {

	os.Setenv("KEY", key)
	calender := CalenderArguments{Summary: "Test Calender", Description: "Test Description", Location: "Test Location"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createCalender", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCalender)
	handler.ServeHTTP(recorder, request)

	Describe("Create Calender", func() {
		Context("create calender", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Create Event with invalid param", func() {

	os.Setenv("KEY", key)
	calender := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEvent)
	handler.ServeHTTP(recorder, request)

	Describe("Create Event", func() {
		Context("create event", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create Event with invalid key", func() {

	os.Setenv("KEY", "mockKey")
	calender := CalenderArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEvent)
	handler.ServeHTTP(recorder, request)

	Describe("Create Event", func() {
		Context("create event", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create Event with invalid param", func() {

	os.Setenv("KEY", key)
	calender := CalenderArguments{EventID: "mockCalender"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEvent)
	handler.ServeHTTP(recorder, request)

	Describe("Create Event", func() {
		Context("create event", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create Event with valid param", func() {

	os.Setenv("KEY", key)
	calender := CalenderArguments{Summary: "Test Calender", Description: "Test Description", CalenderID: calenderID, StartDate: eventStartDate, EndDate: eventEndDate}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calender)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEvent)
	handler.ServeHTTP(recorder, request)

	Describe("Create Event", func() {
		Context("create event", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
