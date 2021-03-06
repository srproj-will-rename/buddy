package carrot

import (
	"fmt"
	"reflect"
	//"strings"
	"strings"
	"testing"
	"time"
)

func TestNewRequest(t *testing.T) {
	sessions := NewDefaultSessionManager()
	token, _, _ := sessions.NewSession()
	ctx, _ := sessions.Get(token)
	str := fmt.Sprintf("{ "+
		"\"session_token\": \"%v\", "+
		"\"endpoint\": \"test\","+
		"\"payload\": { "+
		"\"offset\": { "+
		"\"x\": 3.1,"+
		"\"y\": 1.3,"+
		"\"z\": 4.0 "+
		"}, \"params\": { "+
		"\"foo\": \"bar\" "+
		"} "+
		"} "+
		"}", token)

	actual := NewRequest(ctx, []byte(str))

	params := make(map[string]interface{})
	params["foo"] = "bar"

	expected := &Request{
		SessionToken: ctx.Token,
		endpoint:     "test",
		Params:       params,
		Offset: &offset{
			X: 3.1,
			Y: 1.3,
			Z: 4.0,
		},
		data:    []byte(str),
		metrics: make([]time.Time, MetricCount),
		err:     nil,
	}

	expected.metrics[RequestCreation] = actual.metrics[RequestCreation]

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual.endpoint != expected.endpoint... \n %v \n %v \n", actual, expected)
	}
}

func TestRequestTokenMismatch(t *testing.T) {
	sessions := NewDefaultSessionManager()
	token, _, _ := sessions.NewSession()
	ctx, _ := sessions.Get(token)
	str := `{ "session_token": "badtoken", "endpoint": "print_foo_param", "payload": { "offset": { "x": 3, "y": 1, "z": 4 }, "params": { "foo": "bar" } } }`

	req := NewRequest(ctx, []byte(str))
	if !strings.Contains(req.err.Error(), "mismatch") {
		t.Errorf("request token validation check failed stop request \n "+
			"Error: %v \n Request Token: %v \n Session Token: %v \n", req.err, "badtoken", req.SessionToken)
	}
}
