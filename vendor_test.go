package pagerduty

import (
	"net/http"
	"testing"
)

// ListVendors
func TestVendor_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/vendors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"vendors": [{"id": "1"}]}`))
	})

	var listObj = APIListObject{Limit: 0, Offset: 0, More: false, Total: 0}
	var client = &Client{apiEndpoint: server.URL, authToken: "foo", HTTPClient: defaultHTTPClient}
	var opts = ListVendorOptions{
		APIListObject: listObj,
		Query:         "foo",
	}
	res, err := client.ListVendors(opts)

	want := &ListVendorResponse{
		APIListObject: listObj,
		Vendors: []Vendor{
			{
				APIObject: APIObject{
					ID: "1",
				},
			},
		},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

// Get Vendor
func TestVendor_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/vendors/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"vendor": {"id": "1"}}`))
	})

	var client = &Client{apiEndpoint: server.URL, authToken: "foo", HTTPClient: defaultHTTPClient}
	venID := "1"

	res, err := client.GetVendor(venID)

	want := &Vendor{
		APIObject: APIObject{
			ID: "1",
		},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}
