package registry

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

var supportedTypes = []string{"string", "int", "float", "bool", "bytes"}

// QueryStringParams is a collection of query string parameter definitions
// The Format of a Query string has the following format:
// "key1=<field1:type1>&key2=<field2:type1>"
// The <> around field:type can be omitted
// Semantic: The value for key1 will be mapped to the field field1 in a go struct and is expected to have type type1
// Supported types are int, string, float, bool, bytes
// Bytes are expected to be in Base64-URL (RFC 4648 Section 5)
type QueryStringParams []QSParameter

// QSParameter represents the information of the mapping between parameters transported in the
// query string and how they map to a go struct
type QSParameter struct {
	Key      string // Which key
	Field    string // goes to which field in a go struct
	Type     string // should have what type
	Metadata *FieldMetadatas
}

func qsParameterFromKeyValue(key string, value []string) (*QSParameter, error) {
	qsp := &QSParameter{}
	qsp.Key = key

	// one key should be defined only once
	if len(value) != 1 {
		return nil, fmt.Errorf("1 value expected, got %d (key: %s)", len(value), key)
	}

	s := value[0]
	s = strings.Trim(s, "<>")
	split := strings.Split(s, ":")

	if len(split) > 2 {
		return nil, fmt.Errorf("Invalid format of the parameter defintion, at most one : allowed: %s", s)
	}

	qsp.Field = split[0]

	if len(split) == 2 {
		if !isIn(split[1], supportedTypes...) {
			return nil, fmt.Errorf("%s is not in the supported types %v", split[1], supportedTypes)
		}
		qsp.Type = split[1]
	}

	return qsp, nil
}

// GetParamForKey searches for the param with the given key in the QueryString. If non is found nil is returned.
func (q *QueryStringParams) GetParamForKey(key string) *QSParameter {
	for _, p := range *q {
		if p.Key == key {
			return &p
		}
	}
	return nil
}

// NewQueryStringParams creates at new QueryString from the given definition string
func NewQueryStringParams(definition string) (*QueryStringParams, error) {
	values, err := url.ParseQuery(definition)
	if err != nil {
		return nil, fmt.Errorf("Error Parsing QueryString: %s", err)
	}

	sortedValues := sortParsedQueryString(values)

	var qs QueryStringParams

	for _, v := range sortedValues {
		p, err := qsParameterFromKeyValue(v.Key, v.Values)

		if err != nil {
			return nil, err
		}

		qs = append(qs, *p)
	}

	return &qs, nil
}

type SortedParsedQueryString []QueryStringKV

type QueryStringKV struct {
	Key    string
	Values []string
}

func sortParsedQueryString(vals url.Values) SortedParsedQueryString {
	s := make(SortedParsedQueryString, 0, len(vals))

	for k, v := range vals {
		s = append(s, QueryStringKV{
			Key:    k,
			Values: v,
		})
	}

	sort.Sort(ByKey(s))

	return s
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByKey SortedParsedQueryString

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }
