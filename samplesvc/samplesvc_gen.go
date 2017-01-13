package samplesvc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/doozer-de/rest"
	"github.com/doozer-de/restgen/samplesvc/pb"
	"golang.org/x/net/context"

	"github.com/doozer-de/restgen/pbmap"
)

var (
	baseURI = "/base_uri/"
	_       = ioutil.Discard
)

type QSParameter struct {
	Key   string // Which key
	Field string // goes to which field in a go struct
	Type  string // should have what type
}

func (s *GeneratedService) GetHandlersToRegister() []rest.Register {
	return []rest.Register{
		{Method: "PUT", Path: "/haveall/", Handler: s.CreateHaveAll},
		{Method: "GET", Path: "/haveall/", Handler: s.GetHaveAll},
		{Method: "GET", Path: "/deeppath/", Handler: s.DeepPath},
		{Method: "GET", Path: "/small/:id", Handler: s.GetSmall},
		{Method: "GET", Path: "/small/:id/name", Handler: s.GetSmall1},
		{Method: "PUT", Path: "/small/", Handler: s.PutSmall},
		{Method: "POST", Path: "/small/", Handler: s.PostSmall},
		{Method: "DELETE", Path: "/small/:id", Handler: s.DeleteSmall},
		{Method: "GET", Path: "/getmid/a/:id/b/:id/c/:id/d/:id/e/:id", Handler: s.GetLongPath},
		{Method: "GET", Path: "/deeppathpath/:id", Handler: s.GetPathPath},
		{Method: "GET", Path: "/emptyrequest/", Handler: s.EmptyRequestMessage},
		{Method: "GET", Path: "/pagesortfilter/", Handler: s.PageSortFilter},
	}
}

func (s *GeneratedService) GetBaseURI() string {
	return baseURI
}

func (s *GeneratedService) SetErrorHandler(h rest.ErrorHandler) error {
	if h == nil {
		return fmt.Errorf("ErrorHandler cannot be nil")
	}
	s.errorHandler = h
	return nil
}

type GeneratedService struct {
	GrpcService  pb.BigTestServiceServer
	errorHandler rest.ErrorHandler
}

func NewGeneratedService(s pb.BigTestServiceServer, errorHandler rest.ErrorHandler) (*GeneratedService, error) {
	if s == nil {
		return nil, fmt.Errorf("The given GRPC Service cannot be null")
	}

	return &GeneratedService{
		GrpcService:  s,
		errorHandler: errorHandler,
	}, nil
}

func (s *GeneratedService) CreateHaveAll(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.CreateHaveAllRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}

	resp, err := s.GrpcService.CreateHaveAll(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) GetHaveAll(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.GetHaveAllRequest{}

	values := r.URL.Query()

	if v, ok := values["key3"]; ok {
		if x, ok := rest.ToFloat32(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.Float = x
		}
	}

	if v, ok := values["key4"]; ok {
		if x, ok := rest.ToInt32(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.Int32 = x
		}
	}

	if v, ok := values["key5"]; ok {
		if x, ok := rest.ToInt64(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.Int64 = x
		}
	}

	if v, ok := values["key8"]; ok {
		if x, ok := rest.ToInt64(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.SInt64 = x
		}
	}

	if v, ok := values["key15"]; ok {
		if x, ok := rest.ToString(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.StringField = x
		}
	}

	if v, ok := values["key2"]; ok {
		if x, ok := rest.ToFloat64(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.DoubleField = x
		}
	}

	if v, ok := values["key13"]; ok {
		if x, ok := rest.ToInt64(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.SFixed64 = x
		}
	}

	if v, ok := values["key14"]; ok {
		if x, ok := rest.ToBool(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.BoolField = x
		}
	}

	if v, ok := values["key1"]; ok {
		if x, ok := rest.ToUint32(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.Id = x
		}
	}

	if v, ok := values["key6"]; ok {
		if x, ok := rest.ToUint32(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.UInt32 = x
		}
	}

	if v, ok := values["key16"]; ok {
		if x, ok := rest.ToBytes(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.BytesField = x
		}
	}

	if v, ok := values["key7"]; ok {
		if x, ok := rest.ToUint64(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.UInt64 = x
		}
	}

	if v, ok := values["key9"]; ok {
		if x, ok := rest.ToInt32(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.SInt32 = x
		}
	}

	if v, ok := values["key10"]; ok {
		if x, ok := rest.ToUint32(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.SFixed = x
		}
	}

	if v, ok := values["key11"]; ok {
		if x, ok := rest.ToUint64(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.Fixed64 = x
		}
	}

	if v, ok := values["key12"]; ok {
		if x, ok := rest.ToInt32(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
			}
			req.HaveAll.SFixed32 = x
		}
	}

	resp, err := s.GrpcService.GetHaveAll(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) DeepPath(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.DeepHaveAllRequest{}

	values := r.URL.Query()

	if v, ok := values["key"]; ok {
		if x, ok := rest.ToUint32(v[0]); ok {
			if req.HaveAll == nil {
				req.HaveAll = &pb.HaveAll{}
				if req.HaveAll.Embedded == nil {
					req.HaveAll.Embedded = &pb.Embedded{}
					if req.HaveAll.Embedded.Embedded == nil {
						req.HaveAll.Embedded.Embedded = &pb.Embedded{}
						if req.HaveAll.Embedded.Embedded.Embedded == nil {
							req.HaveAll.Embedded.Embedded.Embedded = &pb.Embedded{}
							if req.HaveAll.Embedded.Embedded.Embedded.Embedded == nil {
								req.HaveAll.Embedded.Embedded.Embedded.Embedded = &pb.Embedded{}
								if req.HaveAll.Embedded.Embedded.Embedded.Embedded.Embedded == nil {
									req.HaveAll.Embedded.Embedded.Embedded.Embedded.Embedded = &pb.Embedded{}
									if req.HaveAll.Embedded.Embedded.Embedded.Embedded.Embedded.Embedded == nil {
										req.HaveAll.Embedded.Embedded.Embedded.Embedded.Embedded.Embedded = &pb.Embedded{}
										if req.HaveAll.Embedded.Embedded.Embedded.Embedded.Embedded.Embedded.Embedded == nil {
											req.HaveAll.Embedded.Embedded.Embedded.Embedded.Embedded.Embedded.Embedded = &pb.Embedded{}
										}
									}
								}
							}
						}
					}
				}
			}
			req.HaveAll.Embedded.Embedded.Embedded.Embedded.Embedded.Embedded.Embedded.Id = x
		}
	}

	resp, err := s.GrpcService.DeepPath(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) GetSmall(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.GetSmallRequest{}

	params := rest.GetParams(ctx)

	if x, ok := rest.ToUint32(params[0].Value); ok {

		req.Id = x
	}

	resp, err := s.GrpcService.GetSmall(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) GetSmall1(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.GetSmallRequest{}

	params := rest.GetParams(ctx)

	if x, ok := rest.ToUint32(params[0].Value); ok {

		req.Idtest = x
	}

	resp, err := s.GrpcService.GetSmall1(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) PutSmall(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.PutSmallRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}

	resp, err := s.GrpcService.PutSmall(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) PostSmall(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.PostSmallRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}

	resp, err := s.GrpcService.PostSmall(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) DeleteSmall(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.DeleteSmallRequest{}

	params := rest.GetParams(ctx)

	if x, ok := rest.ToUint32(params[0].Value); ok {

		req.Id = x
	}

	resp, err := s.GrpcService.DeleteSmall(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) GetLongPath(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.GetLongPathRequest{}

	params := rest.GetParams(ctx)

	if x, ok := rest.ToUint32(params[0].Value); ok {

		req.Ida = x
	}

	if x, ok := rest.ToString(params[1].Value); ok {

		req.Idb = x
	}

	if x, ok := rest.ToUint64(params[2].Value); ok {

		req.Idc = x
	}

	if x, ok := rest.ToInt32(params[3].Value); ok {

		req.Idd = x
	}

	if x, ok := rest.ToString(params[4].Value); ok {

		req.Ide = x
	}

	resp, err := s.GrpcService.GetLongPath(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) GetPathPath(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.DeepPathPathRequest{}

	params := rest.GetParams(ctx)

	if x, ok := rest.ToUint32(params[0].Value); ok {
		if req.HaveAll == nil {
			req.HaveAll = &pb.HaveAll{}
			if req.HaveAll.Embedded == nil {
				req.HaveAll.Embedded = &pb.Embedded{}
				if req.HaveAll.Embedded.Embedded == nil {
					req.HaveAll.Embedded.Embedded = &pb.Embedded{}
				}
			}
		}
		req.HaveAll.Embedded.Embedded.Id = x
	}

	resp, err := s.GrpcService.GetPathPath(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) EmptyRequestMessage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.EmptyRequestMessageRequest{}

	resp, err := s.GrpcService.EmptyRequestMessage(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}

func (s *GeneratedService) PageSortFilter(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := pb.PageSortFilterRequest{}

	values := r.URL.Query()

	if v, ok := values["pageOffset"]; ok {
		if x, ok := rest.ToUint64(v[0]); ok {
			if req.Page == nil {
				req.Page = &pbmap.Page{}
			}
			req.Page.Offset = x
		}
	}

	if v, ok := values["pageLimit"]; ok {
		if x, ok := rest.ToUint64(v[0]); ok {
			if req.Page == nil {
				req.Page = &pbmap.Page{}
			}
			req.Page.Limit = x
		}
	}

	if v, ok := values["sort"]; ok {
		if x, ok := rest.ToString(v[0]); ok {
			if req.Sort == nil {
				req.Sort = &pbmap.Sort{}
			}
			req.Sort.Field = x
		}
	}

	if v, ok := values["desc"]; ok {
		if x, ok := rest.ToBool(v[0]); ok {
			if req.Sort == nil {
				req.Sort = &pbmap.Sort{}
			}
			req.Sort.Desc = x
		}
	}

	if v, ok := values["q"]; ok {
		if x, ok := rest.ToString(v[0]); ok {
			if req.Filter == nil {
				req.Filter = &pbmap.Filter{}
			}
			req.Filter.Query = x
		}
	}

	resp, err := s.GrpcService.PageSortFilter(ctx, &req)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		s.errorHandler(ctx, w, r, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rest.SetStatus(w, resp)
	w.Write(b)
}
