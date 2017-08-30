package registry

// PathParams is a slice of PathParams
type PathParams []PathParam

// PathParam collects information about the different URL path parameters.
type PathParam struct {
	N              int
	FieldRaw       string
	FieldSanitized string
	Metadata       *FieldMetadatas
}

// NewPathParam returns a correct PathParam value
func NewPathParam(field string) (*PathParam, error) {
	return &PathParam{
		FieldSanitized: Sanitize(field),
		FieldRaw:       field,
	}, nil
}
