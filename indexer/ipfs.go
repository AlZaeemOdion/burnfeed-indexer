package indexer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/log"
)

var (
	TypeAggregatedActions = "action"
	ipfsUriPrefix         = "ipfs:"
)

// InvalidActionEventError represents an error which means the coresponding action event is invalid,
// and the indexer will ignore it.
type InvalidActionEventError struct{ err error }

func (e InvalidActionEventError) Error() string { return e.err.Error() }

func NewInvalidActionEventError(err error) *InvalidActionEventError {
	return &InvalidActionEventError{err}
}

// GetActionsByUri gets the aggregated actions list from IPFS.
func (i *ActionsIndexer) GetActionsByUri(uri string, burn uint64) ([]Action, error) {
	uri = strings.TrimPrefix(uri, ipfsUriPrefix)

	log.Debug("Get actions by uri", "uri", uri, "burn", burn)

	// TODO: check the file size
	// Check the aggregated actions list's size at first.
	// stats, err := i.ipfsClient.List(uri)
	// if err != nil {
	// 	return nil, err
	// }

	// log.Info("File stat", "uri", uri, "stat", stats)
	// if len(stats) == 0 {
	// 	return nil, NewInvalidActionEventError(
	// 		fmt.Errorf("empty file stat, uri: %s", uri),
	// 	)
	// }

	// if stats[0].Size > i.sizeLimit {
	// 	return nil, NewInvalidActionEventError(
	// 		fmt.Errorf(
	// 			"file size (%d) exceeds the limit (%d)",
	// 			stats[0].Size,
	// 			i.sizeLimit,
	// 		),
	// 	)
	// }

	// Parse the JSON array.
	rawData, err := i.ipfsClient.Cat(uri)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(rawData)
	if err != nil {
		return nil, err
	}

	var rawActions *AggregatedActions
	if err := json.Unmarshal(b, &rawActions); err != nil {
		return nil, NewInvalidActionEventError(err)
	}

	if rawActions.Type != TypeAggregatedActions {
		return nil, NewInvalidActionEventError(fmt.Errorf("invalid type: %s", rawActions.Type))
	}

	return rawActions.ToSubActions(burn), nil
}

// NewClient creates an http.Client that automatically perform basic auth on each request.
func NewClient(projectId, projectSecret string) *http.Client {
	return &http.Client{
		Transport: &authTransport{
			RoundTripper:  http.DefaultTransport,
			ProjectId:     projectId,
			ProjectSecret: projectSecret,
		},
	}
}

// authTransport decorates each request with a basic auth header.
type authTransport struct {
	http.RoundTripper
	ProjectId     string
	ProjectSecret string
}

// RoundTrip implements http.RoundTripper interface.
func (t *authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.ProjectId, t.ProjectSecret)
	return t.RoundTripper.RoundTrip(r)
}
