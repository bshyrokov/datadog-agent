package dogstatsd

import "github.com/DataDog/datadog-agent/pkg/metrics"

// PacketMetrics contains the samples, events and service checks contained in a
// specific packet.
type PacketMetrics struct {
	Samples       []MetricSample
	Events        []Event
	ServiceChecks []ServiceCheck

	packet *Packet
}

// Release releases the underlying packet by returning it to it's pool.
// The elements of the PacketMetrics are not safe for access beyond this call.
func (pm *PacketMetrics) Release() {
	pm.packet.release()
}

// MetricSample is a metric sample originating from DogStatsD
// Structuraly, this is similar to metrics.MetricSample with []byte slices
// instead of strings. Under the hood those []byte slices are pointing to
// memory allocated in the packet they were received from.
type MetricSample struct {
	packet     *Packet
	Name       []byte
	Value      float64
	SetValue   []byte
	MetricType metrics.MetricType
	Tags       [][]byte
	Hostname   []byte
	SampleRate float64
	Timestamp  float64
}

// Release removes one from the underlying packet reference counting
func (s *MetricSample) Release() {
	if s.packet != nil {
		s.packet.release()
	}
}

// Event is an event originating from DogStatsD
// Structuraly, this is similar to metrics.Event with []byte slices
// instead of strings. Under the hood those []byte slices are pointing to
// memory allocated in the packet they were received from.
type Event struct {
	packet         *Packet
	Title          []byte
	Text           []byte
	Timestamp      int64
	Priority       metrics.EventPriority
	Hostname       []byte
	Tags           [][]byte
	ExtraTags      []string
	AlertType      metrics.EventAlertType
	AggregationKey []byte
	SourceTypeName []byte
}

// Release removes one from the underlying packet reference counting
func (e *Event) Release() {
	if e.packet != nil {
		e.packet.release()
	}
}

// ServiceCheck is a service check originating from DogStatsD
// Structuraly, this is similar to metrics.ServiceCheck with []byte slices
// instead of strings. Under the hood those []byte slices are pointing to
// memory allocated in the packet they were received from.
type ServiceCheck struct {
	packet    *Packet
	Name      []byte
	Hostname  []byte
	Timestamp int64
	Status    metrics.ServiceCheckStatus
	Message   []byte
	Tags      [][]byte
	ExtraTags []string
}

// Release removes one from the underlying packet reference counting
func (sc *ServiceCheck) Release() {
	if sc.packet != nil {
		sc.packet.release()
	}
}