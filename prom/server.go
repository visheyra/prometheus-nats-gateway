package prom

import (
	"errors"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"

	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/prometheus/prompb"

	"github.com/visheyra/prometheus-nats-gateway/nats"
)

type handler struct {
	forwardEndpoint string
	user            string
	pass            string
	topic           string
}

func (h handler) writeHandler(w http.ResponseWriter, r *http.Request) {
	compressed, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reqBuf, err := snappy.Decode(nil, compressed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req prompb.WriteRequest
	if err := proto.Unmarshal(reqBuf, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ts := req.GetTimeseries()
	for _, s := range ts {
		nats.ForwardTimeSerie(*s, h.forwardEndpoint, h.user, h.pass, h.topic)
	}

	w.WriteHeader(http.StatusAccepted)
}

//StartServer ...
func StartServer(listenEndpoint, forwardEndpoint, user, pass, topic string) error {

	l, err := zap.NewProduction()
	if err != nil {
		return errors.New("Can't start server")
	}
	defer l.Sync()
	logger := l.Sugar()

	h := handler{
		forwardEndpoint: forwardEndpoint,
		user:            user,
		pass:            pass,
		topic:           topic,
	}
	http.HandleFunc("/prom", h.writeHandler)
	if err := http.ListenAndServe(listenEndpoint, nil); err != nil {
		logger.Fatalw("Can't start server",
			"listen adress", listenEndpoint,
			"error", err.Error(),
		)
		return err
	}
	return nil
}
