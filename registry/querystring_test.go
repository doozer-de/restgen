package registry

import "testing"

func TestParse(t *testing.T) {
	s := "key1=<value1:int>&key2=<value2:string>&key3=<value3:float>"
	qs, err := NewQueryStringParams(s)

	if err != nil {
		t.Fatalf("Error parsing QueryString definition: %s", err)
	}

	if len(*qs) != 3 {
		t.Fatalf("Expected 3 parameters in qs, got %d", len(*qs))
	}

	key1 := "key1"
	p1 := qs.GetParamForKey(key1)

	if p1 == nil {
		t.Fatalf("Parameters should have been found, wasn't")
	}

	if p1.Key != key1 {
		t.Fatalf("Expected to get params with key %s, got one with key %s", key1, p1.Key)
	}

	if p1.Type != "int" {
		t.Fatalf("Expected to get type int, got %s", p1.Type)
	}

	if p1.Field != "value1" {
		t.Fatalf("Expcted to get field %s, got %s", "value1", p1.Field)
	}
}

func TestMalformedDefinitions(t *testing.T) {
	testcases := []string{
		"key1=value1:int16",                 // int16 is not supported for the interface definition
		"key1=value1:int&key1=value2:float", // Double defintion should be rejected
		"key1&key1",
	}

	for _, tc := range testcases {
		_, err := NewQueryStringParams(tc)

		if err == nil {
			t.Errorf("QueryString definition should have been rejected (gotten an error), but was accepted: %s", tc)
		}
	}
}

func TestAcceptedDefintions(t *testing.T) {
	testcases := []string{
		"key1=field1:int",
		"key1&key2&key3&key4&key5&key6",
		"key1=field1:int&key2=field2:float&key3=field3:bool&key4=field4:string&key5=field5:bytes",
		"key1=field1",
		"key1=:int",
	}

	for _, tc := range testcases {
		_, err := NewQueryStringParams(tc)

		if err != nil {
			t.Errorf("QueryString definition %s should have been accepted, but go an error: %s", tc, err)
		}
	}
}
