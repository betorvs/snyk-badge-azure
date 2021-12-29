package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	Projects = []byte(`{
		"org": {
		  "name": "TestOrg",
		  "id": "b9155176-74f4-2653-23e7-6bba696adaae"
		},
		"projects": [
		  {
			"id": "e48bd952-7a33-0ad8-fec5-e5d644cb9051",
			"name": "TestOrg/repositoryOne",
			"created": "2019-12-02T11:18:13.850Z",
			"origin": "github",
			"type": "sast",
			"readOnly": false,
			"testFrequency": "weekly",
			"isMonitored": true,
			"totalDependencies": null,
			"issueCountsBySeverity": {
			  "low": 0,
			  "high": 0,
			  "medium": 3,
			  "critical": 0
			},
			"lastTestedDate": "2019-11-21T08:37:54.960Z",
			"browseUrl": "https://app.snyk.io/org/TestOrg/project/e48bd952-7a33-0ad8-fec5-e5d644cb9051",
			"owner": null,
			"importingUser": {
			  "id": "6576583d-5e80-0b2e-2962-6624d5274be1",
			  "name": "Roberto Scudeller",
			  "username": "betorvs",
			  "email": "betorvs@TestOrg.com"
			},
			"tags": [],
			"attributes": {
			  "criticality": [],
			  "lifecycle": [],
			  "environment": []
			},
			"branch": "master"
		  },
		  {
			"id": "3d8cba82-98a2-9048-685e-19c1147a3f0b",
			"name": "TestOrg/repositoryTwo",
			"created": "2019-02-13T08:27:05.018Z",
			"origin": "github",
			"type": "sast",
			"readOnly": false,
			"testFrequency": "weekly",
			"isMonitored": true,
			"totalDependencies": null,
			"issueCountsBySeverity": {
			  "low": 0,
			  "high": 0,
			  "medium": 1,
			  "critical": 0
			},
			"lastTestedDate": "2019-12-19T21:00:32.622Z",
			"browseUrl": "https://app.snyk.io/org/TestOrg/project/3d8cba82-98a2-9048-685e-19c1147a3f0b",
			"owner": null,
			"importingUser": {
			  "id": "6576583d-5e80-0b2e-2962-6624d5274be1",
			  "name": "Roberto Scudeller",
			  "username": "betorvs",
			  "email": "betorvs@TestOrg.com"
			},
			"tags": [],
			"attributes": {
			  "criticality": [],
			  "lifecycle": [],
			  "environment": []
			},
			"branch": "master"
		  },
		  {
			"id": "01a88ebb-ee9d-0650-ba1d-c5a93668b36f",
			"name": "TestOrg/repositoryOne:helm/templates/deployment.yaml",
			"created": "2019-02-11T09:36:49.570Z",
			"origin": "github",
			"type": "helmconfig",
			"readOnly": false,
			"testFrequency": "weekly",
			"isMonitored": true,
			"totalDependencies": null,
			"issueCountsBySeverity": {
			  "low": 0,
			  "high": 5,
			  "medium": 0,
			  "critical": 0
			},
			"imageTag": "",
			"imagePlatform": "",
			"imageBaseImage": "",
			"lastTestedDate": "2019-12-18T20:00:14.962Z",
			"browseUrl": "https://app.snyk.io/org/TestOrg/project/01a88ebb-ee9d-0650-ba1d-c5a93668b36f",
			"owner": null,
			"importingUser": {
			  "id": "6576583d-5e80-0b2e-2962-6624d5274be1",
			  "name": "Roberto Scudeller",
			  "username": "betorvs",
			  "email": "betorvs@TestOrg.com"
			},
			"tags": [],
			"attributes": {
			  "criticality": [],
			  "lifecycle": [],
			  "environment": []
			},
			"branch": "master"
		  },
		  {
			"id": "2496d3e5-8a74-ac8d-69c441d3768ee3c6",
			"name": "TestOrg/repositoryThree",
			"created": "2019-02-11T09:36:49.570Z",
			"origin": "github",
			"type": "sast",
			"readOnly": false,
			"testFrequency": "weekly",
			"isMonitored": true,
			"totalDependencies": null,
			"issueCountsBySeverity": {
			  "low": 0,
			  "high": 0,
			  "medium": 0,
			  "critical": 0
			},
			"imageTag": "",
			"imagePlatform": "",
			"imageBaseImage": "",
			"lastTestedDate": "2019-12-18T20:00:14.962Z",
			"browseUrl": "https://app.snyk.io/org/TestOrg/project/2496d3e5-8a74-ac8d-69c441d3768ee3c6",
			"owner": null,
			"importingUser": {
			  "id": "6576583d-5e80-0b2e-2962-6624d5274be1",
			  "name": "Roberto Scudeller",
			  "username": "betorvs",
			  "email": "betorvs@TestOrg.com"
			},
			"tags": [],
			"attributes": {
			  "criticality": [],
			  "lifecycle": [],
			  "environment": []
			},
			"branch": "master"
		  }
		]
	  }`)
	OneProject = []byte(`{
		"id": "3d8cba82-98a2-9048-685e-19c1147a3f0b",
		"name": "TestOrg/repositoryTwo",
		"created": "2019-02-13T08:27:05.018Z",
		"origin": "github",
		"type": "sast",
		"readOnly": false,
		"testFrequency": "weekly",
		"isMonitored": true,
		"totalDependencies": null,
		"issueCountsBySeverity": {
		  "low": 0,
		  "high": 0,
		  "medium": 1,
		  "critical": 0
		},
		"lastTestedDate": "2019-12-19T21:00:32.622Z",
		"browseUrl": "https://app.snyk.io/org/TestOrg/project/3d8cba82-98a2-9048-685e-19c1147a3f0b",
		"owner": null,
		"importingUser": {
		  "id": "6576583d-5e80-0b2e-2962-6624d5274be1",
		  "name": "Roberto Scudeller",
		  "username": "betorvs",
		  "email": "betorvs@TestOrg.com"
		},
		"tags": [],
		"attributes": {
		  "criticality": [],
		  "lifecycle": [],
		  "environment": []
		},
		"branch": "master"
	  }`)
)

func TestVulnerabilitiesFound(t *testing.T) {
	var data map[string]interface{}
	_ = json.Unmarshal([]byte(Projects), &data)

	projects := data["projects"].([]interface{})
	name := "TestOrg/repositoryTwo"
	totalIssues, valid, crit := vulnerabilitiesFound(projects, name, "")
	assert.Equal(t, 1, totalIssues)
	assert.True(t, valid)
	assert.False(t, crit)

	// sum all repositories that match name and id
	name2 := "TestOrg/repositoryOne"
	totalIssues2, valid2, crit := vulnerabilitiesFound(projects, name2, "01a88ebb-ee9d-0650-ba1d-c5a93668b36f")
	assert.Equal(t, 8, totalIssues2)
	assert.True(t, valid2)
	assert.True(t, crit)
}

func TestCountVulnerabilities(t *testing.T) {
	var data map[string]interface{}
	_ = json.Unmarshal([]byte(OneProject), &data)
	total, critical := countVulnerabilities(data)
	assert.Equal(t, 1, total)
	assert.False(t, critical)

}

func TestVersionHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/version", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(versionHandler)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHandler(t *testing.T) {
	var res *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test that r is what you expect it to be
		fmt.Println(r.Header)
		res = r
		_, _ = w.Write([]byte(Projects))
		// fmt.Fprintln(w, []byte(Projects))
	}))
	defer ts.Close()
	APIURL = ts.URL
	Client = &http.Client{}
	tss := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test that r is what you expect it to be
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "body")
	}))
	defer tss.Close()
	GreenURL = tss.URL
	// to avoid it: http://127.0.0.1:65354-5-red?logo=snyk
	// set RedURL with a / in the end
	FoundURL = fmt.Sprintf("%s/", tss.URL)

	// test empty parameters
	req1, err1 := http.NewRequest("GET", "/api/badges", nil)
	assert.NoError(t, err1)
	rr1 := httptest.NewRecorder()
	handler1 := http.HandlerFunc(Handler)

	handler1.ServeHTTP(rr1, req1)
	assert.Equal(t, http.StatusOK, rr1.Code)
	assert.Empty(t, res)

	// repository:  TestOrg/repositoryOne:helm/templates/deployment.yaml
	req2, err2 := http.NewRequest("GET", "/api/badges?org=TestOrg&name=repositoryOne", nil)
	assert.NoError(t, err2)
	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(Handler)

	handler2.ServeHTTP(rr2, req2)
	assert.Equal(t, http.StatusOK, rr2.Code)
	assert.Equal(t, res.Header["Content-Type"], []string{"application/json"})

	// repository:  TestOrg/repositoryOne:helm/templates/deployment.yaml
	req3, err3 := http.NewRequest("GET", "/api/badges?org=TestOrg&name=repositoryOne&id=01a88ebb-ee9d-0650-ba1d-c5a93668b36f", nil)
	assert.NoError(t, err3)
	rr3 := httptest.NewRecorder()
	handler3 := http.HandlerFunc(Handler)

	handler3.ServeHTTP(rr3, req3)
	assert.Equal(t, http.StatusOK, rr3.Code)

	// ??
	fmt.Println("/api/badges?org=TestOrg&name=repositoryThree")
	req4, err4 := http.NewRequest("GET", "/api/badges?org=TestOrg&name=repositoryThree", nil)
	assert.NoError(t, err4)
	rr4 := httptest.NewRecorder()
	handler4 := http.HandlerFunc(Handler)

	handler4.ServeHTTP(rr4, req4)
	assert.Equal(t, http.StatusOK, rr4.Code)

	//repository:  TestOrg/repositoryOne , issues:  0 0 3 0
	// repository:  TestOrg/repositoryOne:helm/templates/deployment.yaml
	req5, err5 := http.NewRequest("GET", "/api/badges?org=TestOrg&name=repositoryOne&id=e48bd952-7a33-0ad8-fec5-e5d644cb9051&id=01a88ebb-ee9d-0650-ba1d-c5a93668b36f", nil)
	assert.NoError(t, err5)
	rr5 := httptest.NewRecorder()
	handler5 := http.HandlerFunc(Handler)

	handler5.ServeHTTP(rr5, req5)
	assert.Equal(t, http.StatusOK, rr5.Code)

	req6, err6 := http.NewRequest("GET", "/api/badges?org=TestOrg&name=repositoryOne&id=e48bd952-7a33-0ad8-fec5-e5d644cb9051,01a88ebb-ee9d-0650-ba1d-c5a93668b36f", nil)
	assert.NoError(t, err6)
	rr6 := httptest.NewRecorder()
	handler6 := http.HandlerFunc(Handler)

	handler6.ServeHTTP(rr6, req6)
	assert.Equal(t, http.StatusOK, rr6.Code)

}

func TestHandlerErrors(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something failed", http.StatusInternalServerError)
	}))
	defer ts.Close()
	APIURL = ts.URL
	Client = &http.Client{}
	tss := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test that r is what you expect it to be
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "body")
	}))
	defer tss.Close()
	GreenURL = tss.URL
	// to avoid it: http://127.0.0.1:65354-5-red?logo=snyk
	// set RedURL with a / in the end
	FoundURL = fmt.Sprintf("%s/", tss.URL)

	// simple error repository not found
	req1, err1 := http.NewRequest("GET", "/api/badges?org=TestOrg&name=repositoryFour", nil)
	assert.NoError(t, err1)
	rr1 := httptest.NewRecorder()
	handler1 := http.HandlerFunc(Handler)

	handler1.ServeHTTP(rr1, req1)
	assert.Equal(t, http.StatusOK, rr1.Code)

	tss2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something failed", http.StatusInternalServerError)
	}))
	defer tss2.Close()
	UnknownURL = tss2.URL

	// force fail with required name or id as parameters
	req2, err2 := http.NewRequest("GET", "/api/badges?org=TestOrg", nil)
	assert.NoError(t, err2)
	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(Handler)

	handler2.ServeHTTP(rr2, req2)
	assert.Equal(t, http.StatusOK, rr2.Code)

	// force url.ParseQuery to fail
	req3, err3 := http.NewRequest("GET", "/api/badges?org=TestOrg%%", nil)
	assert.NoError(t, err3)
	rr3 := httptest.NewRecorder()
	handler3 := http.HandlerFunc(Handler)

	handler3.ServeHTTP(rr3, req3)
	assert.Equal(t, http.StatusOK, rr3.Code)

	tss3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test that r is what you expect it to be
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "body")
	}))
	defer tss3.Close()
	APIURL = tss3.URL
	// forcing json.Unmarshal error
	req4, err4 := http.NewRequest("GET", "/api/badges?org=TestOrg&name=repositoryTwo", nil)
	assert.NoError(t, err4)
	rr4 := httptest.NewRecorder()
	handler4 := http.HandlerFunc(Handler)

	handler4.ServeHTTP(rr4, req4)
	assert.Equal(t, http.StatusOK, rr4.Code)

}
