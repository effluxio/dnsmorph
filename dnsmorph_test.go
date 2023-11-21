package main

import (
	"testing"
)

/*
Helper functions tests
*/
func TestCountChar(t *testing.T) {
	count := countChar("test")
	if len(count) != 3 {
		t.Error("expected map keys lenght of 4, got", len(count))
	}
	if count['t'] != 2 {
		t.Error("expected count['t'] to be 2, got", count['t'])
	}
	if count['e'] != 1 {
		t.Error("expected count['t'] to be 1, got", count['t'])
	}
}

func TestProcessInput(t *testing.T) {
	sanitizedInput, tld := ProcessInput("subdomain.test.co.uk")
	if sanitizedInput != "test" && tld != "co.uk" {
		t.Error("expected 'test' and 'co.uk', got", sanitizedInput, tld)
	}
}

func TestDomainValidation(t *testing.T) {
	if !validateDomainName("yahoo.co.uk") {
		t.Error("expected 'yahoo.co.uk' to be a valid domain")
	}
	if validateDomainName("test") != false {
		t.Error("expected 'test' to be an invalid domain")
	}
}

type testcase struct {
	testString        string
	function          func(string) []string
	expectedResultLen int
	firstResult       string
}

var tests = []testcase{
	{"test", TranspositionAttack, 3, "etst"},
	{"test", AdditionAttack, 26, "testa"},
	{"test", VowelswapAttack, 5, "tast"},
	{"test", SubdomainAttack, 3, "t.est"},
	{"test", ReplacementAttack, 31, "6est"},
	{"test", RepetitionAttack, 4, "ttest"},
	{"test", OmissionAttack, 4, "est"},
	{"test", HyphenationAttack, 3, "t-est"},
	{"test", BitsquattingAttack, 31, "test"},
	{"test", HomographAttack, 27, "τest"},
	{"test.test", DoppelgangerAttack, 1, "testtest"},
}

func TestAttackResults(t *testing.T) {
	for _, test := range tests {
		results := test.function(test.testString)
		if len(results) != test.expectedResultLen {
			t.Errorf("expected array of lenght %d, got %d", test.expectedResultLen, len(results))
		}
		if results[0] != test.firstResult {
			t.Errorf("expected first element of array to be '%s', got %s", test.firstResult, results[0])
		}
	}
}

func TestWhoisLookup(t *testing.T) {

	result := whoisLookup("google.com")
	if result[0] != "1997-09-15T04:00:00Z" {
		t.Errorf("expected 1997-09-15T04:00:00Z, got %s", result[0])
	}

	if result[1] != "2019-09-09T15:39:04Z" {
		t.Errorf("expected 2019-09-09T15:39:04Z, got %s", result[1])
	}
}
