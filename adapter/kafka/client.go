package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/glog"
	"github.com/prometheus/common/model"
	"math"
	"strings"
	"time"
)

type Metric struct {
	Value     float64                `json:"value"`
	Timestamp time.Time              `json:"@timestamp"`
	Labels    map[string]interface{} `json:"labels"`
}

var msgChan = make(chan model.Samples)

func buildLabels(m model.Metric) map[string]interface{} {
	fields := make(map[string]interface{}, len(m))
	for l, v := range m {
		fields[string(l)] = string(v)
	}
	return fields
}

func OfferToChan(samples model.Samples) {
	msgChan <- samples
}

func AsyncProducer(kafkaTopic, kafkaBrokers string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewAsyncProducer(strings.Split(kafkaBrokers, ","), config)

	defer p.Close()
	if err != nil {
		return
	}

	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					glog.Errorln(err)
				}
			case <-success:
			}
		}
	}(p)

	for item := range msgChan {
		for _, s := range item {
			v := float64(s.Value)
			if math.IsNaN(v) || math.IsInf(v, 0) {
				continue
			}

			document := Metric{v, s.Timestamp.Time(), buildLabels(s.Metric)}
			metricsContent, err := json.Marshal(document)

			if err != nil {
				fmt.Printf("error while marshaling document, err: %v", err)
				continue
			}

			msg := &sarama.ProducerMessage{
				Topic: kafkaTopic,
				Value: sarama.ByteEncoder(metricsContent),
			}
			p.Input() <- msg
		}
	}
}
