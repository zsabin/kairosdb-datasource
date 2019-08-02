package kairos_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/zsabin/kairosdb-datasource/pkg/kairos"
	"io/ioutil"
	"testing"
)

func TestKairosDBRequest(t *testing.T) {
	expected := &kairos.Request{
		StartAbsolute: 1357023600000,
		EndAbsolute:   1357024600000,
		Metrics: []*kairos.MetricQuery{
			{
				Name: "abc.123",
				Aggregators: []*kairos.Aggregator{
					{
						Name: "sum",
						Sampling: &kairos.Sampling{
							Value: 10,
							Unit:  "minutes",
						},
					},
				},
				GroupBy: []*kairos.Grouper{
					{
						Name: "tag",
						Tags: []string{
							"data_center",
							"host",
						},
					},
				},
			},
		},
	}

	bytes, readError := ioutil.ReadFile("_testdata/KairosDBRequest.json")
	if readError != nil {
		panic(readError)
	}

	actual := &kairos.Request{}
	parseError := json.Unmarshal(bytes, actual)

	assert.Nil(t, parseError, "Failed to unmarshal request: %v", parseError)
	assert.Equal(t, expected, actual)
}

func TestKairosDBResponse(t *testing.T) {
	expectedResponse := &kairos.Response{
		Queries: []*kairos.QueryResponse{
			{
				Results: []*kairos.QueryResult{
					{
						Name: "abc.123",
						GroupInfo: []*kairos.GroupInfo{
							{
								Name: "type",
							},
							{
								Name: "tag",
								Tags: []string{"host"},
								Group: map[string]string{
									"host": "server1",
								},
							},
						},
						Values: []*kairos.DataPoint{
							{
								1364968800000,
								11019,
							},
							{
								1366351200000,
								2843,
							},
						},
					},
				},
			},
		},
	}

	bytes, readError := ioutil.ReadFile("_testdata/KairosDBResponse.json")
	if readError != nil {
		panic(readError)
	}

	response := &kairos.Response{}
	parseError := json.Unmarshal(bytes, response)

	assert.Nil(t, parseError, "Failed to unmarshal response: %v", parseError)
	assert.Equal(t, expectedResponse, response)
}