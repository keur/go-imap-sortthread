package sortthread

import (
	"errors"
    "strings"
    "fmt"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/responses"
)

type SortResponse struct {
	Ids []uint32
}

type Thread struct {
	Id    uint32
	Child []Thread
}

type ThreadResponse struct {
	Siblings []Thread
}

func (r *SortResponse) Handle(resp imap.Resp) error {
	name, fields, ok := imap.ParseNamedResp(resp)
	if !ok || name != "SORT" {
		return responses.ErrUnhandled
	}

	r.Ids = make([]uint32, len(fields))
	for i, f := range fields {
		if id, err := imap.ParseNumber(f); err != nil {
			return err
		} else {
			r.Ids[i] = id
		}
	}

	return nil
}

func (r *SortResponse) WriteTo(w *imap.Writer) error {
	return errors.New("sortthread: not yet implemented")
}

func (r *ThreadResponse) Handle(resp imap.Resp) error {
    name, fields, ok := imap.ParseNamedResp(resp)
    if !ok || !strings.HasPrefix(name, "THREAD") {
        return responses.ErrUnhandled
    }
    fmt.Println("%s\n", name)
    for _, f := range fields {
        fmt.Println("%+v\n", f)
    }
    return nil
}
