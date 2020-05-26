// Copyright 2017 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The main package for the Prometheus server executable.
package prometheus_kafka

import (
	"github.com/caicloud/prometheus-kafka/adapter/config"
	"github.com/caicloud/prometheus-kafka/adapter/kafka"
	"github.com/caicloud/prometheus-kafka/adapter/writer"
	"github.com/prometheus/common/log"
	"net/http"
	_ "net/http/pprof"
)


func main() {
	cfg := config.GetConfig()
	kafkaTopic := cfg.KafkaTopic
	kafkaBrokers := cfg.KafkaBrokers

	http.HandleFunc("/write", writer.Handle)

	log.Infof("Starting server %s...", cfg.ListenAddr)

	go kafka.AsyncProducer(kafkaTopic , kafkaBrokers )

	http.ListenAndServe(cfg.ListenAddr, nil)
}
