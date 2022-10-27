package samplesvc_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/doozer-de/rest"
	"github.com/doozer-de/restgen/samplesvc"
	"github.com/doozer-de/restgen/samplesvc/pb"
)

func init() {
	mock := samplesvc.NewBigTestMock()

	// add test objects
	mock.Small[0] = &pb.Small{
		Id:   0,
		Text: "Small test object, init object",
	}
	mock.Small[1000] = &pb.Small{
		Id:   1000,
		Text: "Small test object for TestDeleteSmall, init object",
	}

	if mock == nil {
		log.Fatalf("Cannot initialize BigTestMock.")
	}

	gen, err := samplesvc.NewGeneratedService(mock, nil)

	if err != nil {
		log.Fatalf("Error in init")
	}

	go func() {
		rest := rest.New(rest.Configuration{
			BaseURI: "",
		}, []rest.HandlerRegistration{gen})

		if rest == nil {
			log.Fatalf("Testinstance 'rest' is not initialized.")
		}

		log.Print("Start GRPCRESTService on 127.0.0.1:54321")
		log.Fatalf("Error during start of GRPCRESTService: %v", http.ListenAndServe("127.0.0.1:54321", rest))
	}()

	time.Sleep(500 * time.Millisecond)
}

func TestCreateHaveAll(t *testing.T) {
	// 1. Prepare
	r := pb.CreateHaveAllRequest{
		HaveAll: &pb.HaveAll{
			Id:          1,
			DoubleField: 1.1,
			Float:       1.2,
			Int32:       2,
			Int64:       3,
			UInt32:      4,
			UInt64:      5,
			SInt64:      6,
			SInt32:      7,
			SFixed:      8,
			Fixed64:     9,
			SFixed32:    10,
			SFixed64:    11,
			BoolField:   true,
			StringField: "test_stringfield",
			BytesField:  []byte("ok"),
			Embedded: &pb.Embedded{
				Id:   100,
				Text: "embedded struct",
			},
			RepeatedString: []string{"this is", "ok"},
			RepeatedBytes:  [][]byte{[]byte("this is"), []byte("ok")},
			RepeatedEmbedded: []*pb.Embedded{
				{
					Id:   101,
					Text: "embedded struct 1",
				},
				{
					Id:   102,
					Text: "embedded struct 2",
				},
			},
		},
	}

	// prepare http method and URL
	method := "PUT"
	url := "http://127.0.0.1:54321/base_uri/haveall"

	bExpected, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("Cannot marshal '%v'. Error: %v", r, err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bExpected))
	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()

	bActual, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Cannot read the responsed body. Error: %v", err)
	}

	if bActual == nil {
		t.Fatalf("Read response body is unexpectedly nil.")
	}

	if string(bExpected) != string(bActual) {
		t.Errorf("Response doesn't match the expected response.\nExpected: %v\nActual:%v", string(bExpected), string(bActual))
	}
}

func TestGetHaveAll(t *testing.T) {
	// 1. Prepare
	v := url.Values{}
	v.Set("key1", "1")
	v.Set("key2", "1.1")
	v.Set("key3", "1.2")
	v.Set("key4", "2")
	v.Set("key5", "3")
	v.Set("key6", "4")
	v.Set("key7", "5")
	v.Set("key8", "6")
	v.Set("key9", "7")
	v.Set("key10", "8")
	v.Set("key11", "9")
	v.Set("key12", "10")
	v.Set("key13", "11")
	v.Set("key14", "true")
	v.Set("key15", "test_stringfield")

	encoded := base64.URLEncoding.EncodeToString([]byte("ok"))
	v.Set("key16", encoded)

	// prepare http method and URL
	method := "GET"
	url := "http://127.0.0.1:54321/base_uri/haveall"

	req, err := http.NewRequest(method, fmt.Sprintf("%s?%s", url, v.Encode()), nil)
	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Cannot read the responsed body. Error: %v", err)
	}

	var r pb.GetHaveAllResponse
	if err = json.Unmarshal(body, &r); err != nil {
		t.Errorf("Cannot unmarshal the responsed body '%v'. Error: %v", string(body), err)
	}

	ha := r.HaveAll
	if ha.Id != 1 {
		t.Errorf("res.ID '%d' != the exp.ID '1'.", ha.Id)
	}

	if ha.DoubleField != 1.1 {
		t.Errorf("res.Double '%f' != exp.Double '1.1'", ha.DoubleField)
	}

	if ha.Float != 1.2 {
		t.Errorf("res.Float '%f' != exp.Float '1.2'", ha.Float)
	}

	if ha.Int32 != 2 {
		t.Errorf("res.Int32 '%d' != exp.Int32 '2'", ha.Int32)
	}

	if ha.Int64 != 3 {
		t.Errorf("res.Int64 '%d' != exp.Int64 '3'", ha.Int64)
	}

	if ha.UInt32 != 4 {
		t.Errorf("res.UInt32 '%d' != exp.UInt32 '4'", ha.UInt32)
	}

	if ha.UInt64 != 5 {
		t.Errorf("res.UInt64 '%d' != exp.UInt64 '5'", ha.UInt64)
	}

	if ha.SInt32 != 7 {
		t.Errorf("res.SInt32 '%d' != exp.SInt32 '7'", ha.SInt32)
	}

	if ha.SInt64 != 6 {
		t.Errorf("res.SInt64 '%d' != exp.SInt64 '6'", ha.SInt64)
	}

	if ha.SFixed != 8 {
		t.Errorf("res.Fixed '%d' != exp.SFixed '8'", ha.SFixed)
	}

	if ha.Fixed64 != 9 {
		t.Errorf("res.Fixed64 '%d' != exp.Fixed64 '9'", ha.Fixed64)
	}

	if ha.SFixed32 != 10 {
		t.Errorf("res.SFixed32 '%d' != exp.SFixed32 '10'", ha.SFixed32)
	}

	if ha.SFixed64 != 11 {
		t.Errorf("res.SFixed64 '%d' != exp.SFixed64 '11'", ha.SFixed64)
	}

	if !ha.BoolField {
		t.Errorf("res.Bool '%t' != exp.Bool 'true'", ha.BoolField)
	}

	if ha.StringField != "test_stringfield" {
		t.Errorf("res.StringField '%s' != exp.StringField 'test_stringfield'", ha.StringField)
	}
}

func TestGetDeepPath(t *testing.T) {
	// 1. Prepare
	v := url.Values{}
	v.Set("key", "1")

	// prepare http method and URL
	method := "GET"
	url := "http://127.0.0.1:54321/base_uri/deeppath/"

	req, err := http.NewRequest(method, fmt.Sprintf("%s?%s", url, v.Encode()), nil)
	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Cannot read the responsed body. Error: %v", err)
	}

	var r pb.DeepHaveAllResponse
	if err = json.Unmarshal(body, &r); err != nil {
		t.Fatalf("Cannot unmarshal the responded body '%v'. Error: %v", string(body), err)
	}

	if r.Id != 1 {
		t.Errorf("res.ID '%d' != exp.Id '1'", r.Id)
	}
}

func TestPutSmall(t *testing.T) {
	// 1. Prepare
	r := &pb.PutSmallRequest{
		Small: &pb.Small{
			Id:   0,
			Text: "Small object from TestPutSmall",
		},
	}

	// prepare http method and URL
	method := "PUT"
	url := "http://127.0.0.1:54321/base_uri/small/"

	b, err := json.Marshal(&r)
	if err != nil {
		t.Fatalf("Cannot marshal the request '%v'. Error: %v", r, err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()
}

func TestPostSmall(t *testing.T) {
	// 1. Prepare
	r := &pb.PostSmallRequest{
		Small: &pb.Small{
			Id:   1,
			Text: "Small object from TestPostSmall",
		},
	}

	// prepare http method and URL
	method := "POST"
	url := "http://127.0.0.1:54321/base_uri/small/"

	b, err := json.Marshal(&r)
	if err != nil {
		t.Fatalf("Cannot marshal the request '%v'. Error: %v", r, err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()
}

func TestGetSmall(t *testing.T) {
	// 1. Prepare
	// prepare http method and URL
	method := "GET"
	url := "http://127.0.0.1:54321/base_uri/small/"
	req, err := http.NewRequest(method, fmt.Sprintf("%s0", url), nil)

	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Cannot read the responsed body. Error: %v", err)
	}

	var r pb.GetSmallResponse
	if err := json.Unmarshal(body, &r); err != nil {
		t.Fatalf("Cannot unmarshal responded body '%v'. Error: %v", string(body), err)
	}

	if r.Small.Text == "" {
		t.Errorf("resp.Small.Text '%s' == nil", r.Small.Text)
	}
}

func TestDeleteSmall(t *testing.T) {
	// 1. Prepare
	// prepare http method and URL
	method := "DELETE"
	url := "http://127.0.0.1:54321/base_uri/small/"

	req, err := http.NewRequest(method, fmt.Sprintf("%s1000", url), nil)
	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()
}

func TestGetLongPath(t *testing.T) {
	// 1. Prepare
	// prepare http method and URL
	method := "GET"
	url := "http://127.0.0.1:54321/base_uri/getmid/"

	req, err := http.NewRequest(method, fmt.Sprintf("%sa/1/b/teststring_Idb/c/2/d/-1/e/teststring_Ide", url), nil)
	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Cannot read the responsed body. Error: %v", err)
	}

	var r pb.GetLongPathResponse
	if err := json.Unmarshal(body, &r); err != nil {
		t.Fatalf("Canot unmarshal the responded body '%s'. Error: '%v'", string(body), err)
	}

	if r.Idf != 1 {
		t.Errorf("resp.Idf '%d' != exp.Idf '1'", r.Idf)
	}
}

func TestGetDeepPathPath(t *testing.T) {
	// 1. Prepare
	// prepare http method and URL
	method := "GET"
	url := "http://127.0.0.1:54321/base_uri/deeppathpath/1/"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Cannot read the responsed body. Error: %v", err)
	}

	if string(body) != `{"Id":1}` {
		t.Errorf("resp '%s' != exp '{\"Id\":1}'", string(body))
	}
}

func TestPageSortFilter(t *testing.T) {
	// 1. Prepare
	// prepare http method and URL
	method := "GET"
	url := "http://127.0.0.1:54321/base_uri/pagesortfilter"
	qs := fmt.Sprintf("?pageLimit=%d&pageOffset=%d&sort=sortField&desc=%t&q=filterQuery", 100, 10, true)

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", url, qs), nil)
	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Cannot read the responsed body. Error: %v", err)
	}

	var r pb.PageSortFilterResponse
	if err = json.Unmarshal(body, &r); err != nil {
		t.Fatalf("Cannot unmarshal the responsed body. Error: %v", err)
	}

	if r.PageLimit != 100 {
		t.Errorf("resp '%d' != exp '100'", r.PageLimit)
	}

	if r.PageOffset != 10 {
		t.Errorf("resp '%d' != exp '10'", r.PageOffset)
	}

	if r.SortField != "sortField" {
		t.Errorf("resp '%s' != exp 'sortField'", r.SortField)
	}

	if !r.SortDesc {
		t.Errorf("resp '%t' != exp 'true'", r.SortDesc)
	}

	if r.FilterQuery != "filterQuery" {
		t.Errorf("resp '%s' != exp 'filterQuery'", r.FilterQuery)
	}
}

func TestReturnDerrorMessage(t *testing.T) {
	// 1. Prepare
	// prepare http method and URL
	method := "GET"
	url := "http://127.0.0.1:54321/base_uri/returnderror"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatalf("Cannot create a new request. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 2. Act
	resp, err := client.Do(req)

	// 3. Evaluate
	if err != nil {
		t.Errorf("Request was not successful: %s", err)
	}
	defer resp.Body.Close()
}
