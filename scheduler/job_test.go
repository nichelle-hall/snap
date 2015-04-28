package scheduler

import (
	"testing"
	"time"

	"github.com/intelsdilabs/pulse/core"

	. "github.com/smartystreets/goconvey/convey"
)

type mockCollector struct{}

func (m *mockCollector) CollectMetrics([]core.Metric, time.Time) ([]core.Metric, []error) {
	return nil, nil
}

func TestCollectorJob(t *testing.T) {
	Convey("newCollectorJob()", t, func() {
		Convey("it returns an init-ed collectorJob", func() {
			cj := newCollectorJob([]core.Metric{}, defaultDeadline, &mockCollector{})
			So(cj, ShouldHaveSameTypeAs, &collectorJob{})
		})
	})
	Convey("StartTime()", t, func() {
		Convey("it should return the job starttime", func() {
			cj := newCollectorJob([]core.Metric{}, defaultDeadline, &mockCollector{})
			So(cj.StartTime(), ShouldHaveSameTypeAs, time.Now())
		})
	})
	Convey("Deadline()", t, func() {
		Convey("it should return the job daedline", func() {
			cj := newCollectorJob([]core.Metric{}, defaultDeadline, &mockCollector{})
			So(cj.Deadline(), ShouldResemble, cj.deadline)
		})
	})
	Convey("Type()", t, func() {
		Convey("it should return the job type", func() {
			cj := newCollectorJob([]core.Metric{}, defaultDeadline, &mockCollector{})
			So(cj.Type(), ShouldEqual, collectJobType)
		})
	})
	Convey("ReplChan()", t, func() {
		Convey("it should return the reply channel", func() {
			cj := newCollectorJob([]core.Metric{}, defaultDeadline, &mockCollector{})
			So(cj.ReplChan(), ShouldHaveSameTypeAs, make(chan struct{}))
		})
	})
	// Convey("Metrics()", t, func() {
	// 	Convey("it should return the job metrics", func() {
	// 		cj := newCollectorJob([]core.MetricType{}, defaultDeadline, &mockCollector{})
	// 		So(cj.Metrics(), ShouldResemble, []core.Metric{})
	// 	})
	// })
	Convey("Errors()", t, func() {
		Convey("it should return the errors from the job", func() {
			cj := newCollectorJob([]core.Metric{}, defaultDeadline, &mockCollector{})
			So(cj.Errors(), ShouldResemble, []error{})
		})
	})
	Convey("Run()", t, func() {
		Convey("it should reply on the reply chan", func() {
			cj := newCollectorJob([]core.Metric{}, defaultDeadline, &mockCollector{})
			go cj.Run()
			<-cj.replchan
			So(cj.Errors(), ShouldResemble, []error{})
		})
	})
}
