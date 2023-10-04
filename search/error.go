package search

import (
	"fmt"

	"github.com/imroc/req/v3"
)

type SearchError struct {
	ErrorBody struct {
		Status  int
		Message string
	} `json:"error"`
}

func (err *SearchError) Error() string {
	return fmt.Sprintf("api error: %v (%v)", err.ErrorBody.Message, err.ErrorBody.Status)
}

func init() {
	client.
		OnBeforeRequest(func(c *req.Client, r *req.Request) error {
			if r.RetryAttempt == 0 {
				r.EnableDump()
			}
			return nil
		}).
		SetCommonErrorResult(&SearchError{}).
		OnAfterResponse(func(c *req.Client, resp *req.Response) error {
			if err, ok := resp.ErrorResult().(*SearchError); ok {
				return err
			}

			if !resp.IsSuccessState() {
				return fmt.Errorf("Spotify Bad Response: %v", resp.Dump())
			}

			return nil
		})
}
