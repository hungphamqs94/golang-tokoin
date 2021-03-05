package main

import "testing"

func TestSearchOrganizationMany(t *testing.T){

	readFile();

	termWrongResult := searchOrganizationMany("OrganizationId", "1")
	if termWrongResult != "Term not exists at organization"{
		t.Errorf("Output Term not exists at organization instead od %v", termWrongResult)
	}

	termValueWrongType := searchOrganizationMany("_id", "dksdjsd")
	if termValueWrongType != "Wrong value type" {
		t.Errorf("Output expect Wrong value type instead of %v", termValueWrongType)
	}

	noResultsFound := searchOrganizationMany("_id", "25251")
	if noResultsFound != "No results found"{
		t.Errorf("Output expect No results found instead of %v", noResultsFound)
	}

	resultsFound := searchOrganizationMany("_id", "101")
	if resultsFound != "Have results"{
		t.Errorf("Output expect Have results instead of %v", resultsFound)
	}
}