package main

import "testing"


func TestSearchTicketMany(t *testing.T){

	readFile()

	termWrongResult := searchTicketMany("TicketId", "1")
	if termWrongResult != "Term not exists at ticket"{
		t.Errorf("Output Term not exists at ticket instead od %v", termWrongResult)
	}

	noResultsFound := searchTicketMany("_id", "25251")
	if noResultsFound != "No results found"{
		t.Errorf("Output expect No results found instead of %v", noResultsFound)
	}

	resultsFound := searchTicketMany("_id", "436bf9b0-1147-4c0a-8439-6f79833bff5b")
	if resultsFound != "Have results"{
		t.Errorf("Output expect Have results instead of %v", resultsFound)
	}

}