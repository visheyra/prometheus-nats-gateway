package nats

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/prometheus/prometheus/prompb"
	"go.uber.org/zap"
)

//ForwardTimeSerie ...
func ForwardTimeSerie(t prompb.TimeSeries, forwardAddr, user, pass, topic string) {

	l, err := zap.NewProduction()
	if err != nil {
		return
	}
	defer l.Sync()
	logger := l.Sugar()

	//
	// Connect to nats server
	//
	cnx, err := nats.Connect(forwardAddr, nats.UserInfo(user, pass))
	if err != nil {
		logger.Errorw("Can't establish connection to nats endpoint",
			"endpoint", forwardAddr,
			"error", err.Error(),
		)
		return
	}
	defer cnx.Close()

	//
	// Format labels as map[string]string
	//

	lbs := make(map[string]string)
	for _, l := range t.GetLabels() {
		lbs[l.GetName()] = l.GetValue()
	}

	logger.Debugw("Sending time serie",
		"length", len(t.GetSamples()),
		"endpoint", forwardAddr,
		"labels", lbs,
	)

	for _, v := range t.GetSamples() {

		//
		// Marshal point to Json
		//

		value := strconv.FormatFloat(v.GetValue(), 'E', -1, 64)
		b, err := json.Marshal(point{
			Value:     value,
			Labels:    lbs,
			TimeStamp: time.Now().Unix(),
		})

		if err != nil {
			logger.Errorw("Can't marshal point to json",
				"labels", lbs,
				"value", value,
			)
			continue
		}

		//
		// Publish point to nats
		//

		if cnx.Publish(topic, b) != nil {
			logger.Errorw("Can't send Json to endpoint",
				"endpoint", forwardAddr,
				"topic", topic,
				"data", string(b))
		} else {
			logger.Debugw("Successfully sent metric",
				"endpoint", forwardAddr,
				"topic", topic,
				"size", len(b))
		}
	}
}

type point struct {
	Value     string            `json:"value"`
	Labels    map[string]string `json:"labels"`
	TimeStamp int64             `json:"unix"`
}
