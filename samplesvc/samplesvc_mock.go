package samplesvc

import (
	"fmt"

	"github.com/doozer-de/restgen/samplesvc/pb"

	"golang.org/x/net/context"
)

// BigTestMock represents logic reachable via a REST endpoint.
type BigTestMock struct {
	Small map[uint32]*pb.Small
	pb.UnimplementedBigTestServiceServer
}

// NewBigTestMock instantiates a new BigTestMock
func NewBigTestMock() *BigTestMock {
	return &BigTestMock{
		Small: make(map[uint32]*pb.Small),
	}
}

// CreateHaveAll should create a HaveAll message.
// Excluded from HaveAll are the "repeated"-fields (slices). This is not supported at the moment.
func (b *BigTestMock) CreateHaveAll(ctx context.Context, req *pb.CreateHaveAllRequest) (*pb.CreateHaveAllResponse, error) {
	return &pb.CreateHaveAllResponse{
		HaveAll: &pb.HaveAll{
			Id:                     req.HaveAll.Id,
			DoubleField:            req.HaveAll.DoubleField,
			Float:                  req.HaveAll.Float,
			Int32:                  req.HaveAll.Int32,
			Int64:                  req.HaveAll.Int64,
			UInt32:                 req.HaveAll.UInt32,
			UInt64:                 req.HaveAll.UInt64,
			SInt64:                 req.HaveAll.SInt64,
			SInt32:                 req.HaveAll.SInt32,
			SFixed:                 req.HaveAll.SFixed,
			Fixed64:                req.HaveAll.Fixed64,
			SFixed32:               req.HaveAll.SFixed32,
			SFixed64:               req.HaveAll.SFixed64,
			BoolField:              req.HaveAll.BoolField,
			StringField:            req.HaveAll.StringField,
			BytesField:             req.HaveAll.BytesField,
			Embedded:               req.HaveAll.Embedded,
			RepeatedString:         req.HaveAll.RepeatedString,
			RepeatedBytes:          req.HaveAll.RepeatedBytes,
			RepeatedEmbedded:       req.HaveAll.RepeatedEmbedded,
			MapStringStringField:   req.HaveAll.MapStringStringField,
			MapStringEmbeddedField: req.HaveAll.MapStringEmbeddedField,
		},
	}, nil
}

// GetHaveAll should retrieve the HaveAll message with the id given in the path
// Excluded from HaveAll are the "repeated"-fields (slices). This is not supported at the moment.
func (b *BigTestMock) GetHaveAll(ctx context.Context, req *pb.GetHaveAllRequest) (*pb.GetHaveAllResponse, error) {
	return &pb.GetHaveAllResponse{
		HaveAll: &pb.HaveAll{
			Id:                     req.HaveAll.Id,
			DoubleField:            req.HaveAll.DoubleField,
			Float:                  req.HaveAll.Float,
			Int32:                  req.HaveAll.Int32,
			Int64:                  req.HaveAll.Int64,
			UInt32:                 req.HaveAll.UInt32,
			UInt64:                 req.HaveAll.UInt64,
			SInt64:                 req.HaveAll.SInt64,
			SInt32:                 req.HaveAll.SInt32,
			SFixed:                 req.HaveAll.SFixed,
			Fixed64:                req.HaveAll.Fixed64,
			SFixed32:               req.HaveAll.SFixed32,
			SFixed64:               req.HaveAll.SFixed64,
			BoolField:              req.HaveAll.BoolField,
			StringField:            req.HaveAll.StringField,
			BytesField:             req.HaveAll.BytesField,
			Embedded:               req.HaveAll.Embedded,
			RepeatedString:         req.HaveAll.RepeatedString,
			RepeatedBytes:          req.HaveAll.RepeatedBytes,
			RepeatedEmbedded:       req.HaveAll.RepeatedEmbedded,
			MapStringStringField:   req.HaveAll.MapStringStringField,
			MapStringEmbeddedField: req.HaveAll.MapStringEmbeddedField,
		},
	}, nil
}

// DeepPath retrieves a value deeply nested in structs and returns this value.
func (b *BigTestMock) DeepPath(ctx context.Context, req *pb.DeepHaveAllRequest) (*pb.DeepHaveAllResponse, error) {
	var id uint32
	x := req.HaveAll
	if x != nil {
		y := x.Embedded
		if y != nil {
			y = y.Embedded
			if y != nil {
				y = y.Embedded
				if y != nil {
					y = y.Embedded
					if y != nil {
						y = y.Embedded
						if y != nil {
							y = y.Embedded
							if y != nil {
								y = y.Embedded
								if y != nil {
									id = y.Id
								} else {
									return nil, fmt.Errorf("8th level of request is nil")
								}
							} else {
								return nil, fmt.Errorf("7th level of request is nil")
							}
						} else {
							return nil, fmt.Errorf("6th level of request is nil")
						}
					} else {
						return nil, fmt.Errorf("5th level of request is nil")
					}
				} else {
					return nil, fmt.Errorf("4th level of request is nil")
				}
			} else {
				return nil, fmt.Errorf("3th level of request is nil")
			}
		} else {
			return nil, fmt.Errorf("2th level of request is nil")
		}
	} else {
		return nil, fmt.Errorf("1th level of request is nil")
	}

	return &pb.DeepHaveAllResponse{
		Id: id,
	}, nil
}

// DoNotCreate is a testcase for a method without an method_map option. Therefore no REST frontend should be generated
func (b *BigTestMock) DoNotCreate(ctx context.Context, req *pb.DoNotCreateRequest) (*pb.DoNotCreateResponse, error) {
	panic("function 'DoNotCreate' should not be reachable over a REST endpoint. Check, why there is an REST endpoint for this function")
}

// GetSmall returns a Small message identified by a given ID.
func (b *BigTestMock) GetSmall(ctx context.Context, req *pb.GetSmallRequest) (*pb.GetSmallResponse, error) {
	s, ok := b.Small[req.Id]
	if !ok {
		return nil, fmt.Errorf("error during getting the Small-Object with ID %d", req.Id)
	}

	return &pb.GetSmallResponse{
		Small: s,
	}, nil
}

// GetSmall1 is just a test for route registration, that's why no logic.
func (b *BigTestMock) GetSmall1(ctx context.Context, req *pb.GetSmallRequest) (*pb.GetSmallResponse, error) {

	return nil, nil
}

// PutSmall adds a Small message to a storage.
func (b *BigTestMock) PutSmall(ctx context.Context, req *pb.PutSmallRequest) (*pb.PutSmallResponse, error) {
	if _, ok := b.Small[req.Small.Id]; !ok {
		return nil, fmt.Errorf("object with ID %d not found", req.Small.Id)
	}

	// Update object
	b.Small[req.Small.Id] = req.Small

	return &pb.PutSmallResponse{}, nil
}

// PostSmall is a flaw: id of the updated resource should be in the URI. An id from the URI should be mapped into posted objects.
func (b *BigTestMock) PostSmall(ctx context.Context, req *pb.PostSmallRequest) (*pb.PostSmallResponse, error) {
	nextID := len(b.Small) + 1

	// add item with next ID
	b.Small[uint32(nextID)] = req.Small

	return &pb.PostSmallResponse{}, nil
}

// DeleteSmall deletes a Small message from a storage.
func (b *BigTestMock) DeleteSmall(ctx context.Context, req *pb.DeleteSmallRequest) (*pb.DeleteSmallResponse, error) {
	if _, ok := b.Small[req.Id]; !ok {
		return nil, fmt.Errorf("'small'-Object with ID %d doesn't exist", req.Id)
	}

	delete(b.Small, req.Id)

	return &pb.DeleteSmallResponse{}, nil
}

// GetLongPath returns one value of more than one path parameter.
func (b *BigTestMock) GetLongPath(ctx context.Context, req *pb.GetLongPathRequest) (*pb.GetLongPathResponse, error) {
	return &pb.GetLongPathResponse{
		// XXX(ph) why a "wrong" value is returned?
		Idf: req.Ida,
	}, nil
}

// GetPathPath returns the given value of a deeply nested struct.
func (b *BigTestMock) GetPathPath(ctx context.Context, req *pb.DeepPathPathRequest) (*pb.DeepPathPathResponse, error) {
	var id uint32

	x := req.HaveAll
	if x != nil {
		y := x.Embedded
		if y != nil {
			y = y.Embedded
			if y != nil {
				id = y.Id
			} else {
				return nil, fmt.Errorf("3rd level is nil")
			}
		} else {
			return nil, fmt.Errorf("2nd level is nil")
		}
	} else {
		return nil, fmt.Errorf("1st level is nil")
	}

	return &pb.DeepPathPathResponse{
		Id: id,
	}, nil
}

// EmptyRequestMessage returns an empty protobuf message.
func (b *BigTestMock) EmptyRequestMessage(ctx context.Context, req *pb.EmptyRequestMessageRequest) (*pb.EmptyRequestMessageResponse, error) {
	return &pb.EmptyRequestMessageResponse{}, nil
}

// PageSortFilter returns the given values of the Page, Sort and Filter messages.
func (b *BigTestMock) PageSortFilter(ctx context.Context, req *pb.PageSortFilterRequest) (*pb.PageSortFilterResponse, error) {
	return &pb.PageSortFilterResponse{
		PageLimit:   req.Page.Limit,
		PageOffset:  req.Page.Offset,
		SortField:   req.Sort.Field,
		SortDesc:    req.Sort.Desc,
		FilterQuery: req.Filter.Query,
	}, nil
}
