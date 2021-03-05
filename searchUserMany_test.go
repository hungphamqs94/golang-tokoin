package main

import "testing"

func TestSearchUserMany(t *testing.T) {

	readFile()
	termWrongResult := searchUserMany("UserId", "1")
	if termWrongResult != "Term not exists at user" {
		t.Errorf("Output Term not exists at user instead od %v", termWrongResult)
	}

	termValueWrongType := searchUserMany("_id", "dkdjd")
	if termValueWrongType != "Wrong value type" {
		t.Errorf("Output expect Wrong value type instead of %v", termValueWrongType)
	}

	noResultsFound := searchUserMany("_id", "33333")
	if noResultsFound != "No results found" {
		t.Errorf("Output expect No results found instead of %v", noResultsFound)
	}

	resultsFound := searchUserMany("_id", "1")
	if resultsFound != "Have results" {
		t.Errorf("Output expect Have results instead of %v", resultsFound)
	}

}
