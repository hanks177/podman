package bindings_test

import (
	"sync"
	"time"

	"github.com/hanks177/podman/v4/pkg/bindings/containers"
	"github.com/hanks177/podman/v4/pkg/bindings/pods"
	"github.com/hanks177/podman/v4/pkg/bindings/system"
	"github.com/hanks177/podman/v4/pkg/bindings/volumes"
	"github.com/hanks177/podman/v4/pkg/domain/entities"
	"github.com/hanks177/podman/v4/pkg/domain/entities/reports"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Podman system", func() {
	var (
		bt     *bindingTest
		s      *gexec.Session
		newpod string
	)

	BeforeEach(func() {
		bt = newBindingTest()
		bt.RestoreImagesFromCache()
		newpod = "newpod"
		bt.Podcreate(&newpod)
		s = bt.startAPIService()
		time.Sleep(1 * time.Second)
		err := bt.NewConnection()
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		s.Kill()
		bt.cleanup()
	})

	It("podman events", func() {
		var name = "top"
		_, err := bt.RunTopContainer(&name, nil)
		Expect(err).To(BeNil())

		filters := make(map[string][]string)
		filters["container"] = []string{name}

		binChan := make(chan entities.Event)
		done := sync.Mutex{}
		done.Lock()
		eventCounter := 0
		go func() {
			defer done.Unlock()
			for range binChan {
				eventCounter++
			}
		}()
		options := new(system.EventsOptions).WithFilters(filters).WithStream(false)
		err = system.Events(bt.conn, binChan, nil, options)
		Expect(err).To(BeNil())
		done.Lock()
		Expect(eventCounter).To(BeNumerically(">", 0))
	})

	It("podman system prune - pod,container stopped", func() {
		// Start and stop a pod to enter in exited state.
		_, err := pods.Start(bt.conn, newpod, nil)
		Expect(err).To(BeNil())
		_, err = pods.Stop(bt.conn, newpod, nil)
		Expect(err).To(BeNil())
		// Start and stop a container to enter in exited state.
		var name = "top"
		_, err = bt.RunTopContainer(&name, nil)
		Expect(err).To(BeNil())
		err = containers.Stop(bt.conn, name, nil)
		Expect(err).To(BeNil())

		options := new(system.PruneOptions).WithAll(true)
		systemPruneResponse, err := system.Prune(bt.conn, options)
		Expect(err).To(BeNil())
		Expect(len(systemPruneResponse.PodPruneReport)).To(Equal(1))
		Expect(len(systemPruneResponse.ContainerPruneReports)).To(Equal(1))
		Expect(len(systemPruneResponse.ImagePruneReports)).
			To(BeNumerically(">", 0))
		Expect(len(systemPruneResponse.VolumePruneReports)).To(Equal(0))
	})

	It("podman system prune running alpine container", func() {
		// Start and stop a pod to enter in exited state.
		_, err := pods.Start(bt.conn, newpod, nil)
		Expect(err).To(BeNil())
		_, err = pods.Stop(bt.conn, newpod, nil)
		Expect(err).To(BeNil())

		// Start and stop a container to enter in exited state.
		var name = "top"
		_, err = bt.RunTopContainer(&name, nil)
		Expect(err).To(BeNil())
		err = containers.Stop(bt.conn, name, nil)
		Expect(err).To(BeNil())

		// Start container and leave in running
		var name2 = "top2"
		_, err = bt.RunTopContainer(&name2, nil)
		Expect(err).To(BeNil())

		// Adding an unused volume
		_, err = volumes.Create(bt.conn, entities.VolumeCreateOptions{}, nil)
		Expect(err).To(BeNil())
		options := new(system.PruneOptions).WithAll(true)
		systemPruneResponse, err := system.Prune(bt.conn, options)
		Expect(err).To(BeNil())
		Expect(len(systemPruneResponse.PodPruneReport)).To(Equal(1))
		Expect(len(systemPruneResponse.ContainerPruneReports)).To(Equal(1))
		Expect(len(systemPruneResponse.ImagePruneReports)).
			To(BeNumerically(">", 0))
		// Alpine image should not be pruned as used by running container
		Expect(reports.PruneReportsIds(systemPruneResponse.ImagePruneReports)).
			ToNot(ContainElement("docker.io/library/alpine:latest"))
		// Though unused volume is available it should not be pruned as flag set to false.
		Expect(len(systemPruneResponse.VolumePruneReports)).To(Equal(0))
	})

	It("podman system prune running alpine container volume prune", func() {
		// Start a pod and leave it running
		_, err := pods.Start(bt.conn, newpod, nil)
		Expect(err).To(BeNil())

		// Start and stop a container to enter in exited state.
		var name = "top"
		_, err = bt.RunTopContainer(&name, nil)
		Expect(err).To(BeNil())
		err = containers.Stop(bt.conn, name, nil)
		Expect(err).To(BeNil())

		// Start second container and leave in running
		var name2 = "top2"
		_, err = bt.RunTopContainer(&name2, nil)
		Expect(err).To(BeNil())

		// Adding an unused volume should work
		_, err = volumes.Create(bt.conn, entities.VolumeCreateOptions{}, nil)
		Expect(err).To(BeNil())

		options := new(system.PruneOptions).WithAll(true).WithVolumes(true)
		systemPruneResponse, err := system.Prune(bt.conn, options)
		Expect(err).To(BeNil())
		Expect(len(systemPruneResponse.PodPruneReport)).To(Equal(0))
		Expect(len(systemPruneResponse.ContainerPruneReports)).To(Equal(1))
		Expect(len(systemPruneResponse.ImagePruneReports)).
			To(BeNumerically(">", 0))
		// Alpine image should not be pruned as used by running container
		Expect(reports.PruneReportsIds(systemPruneResponse.ImagePruneReports)).
			ToNot(ContainElement("docker.io/library/alpine:latest"))
		// Volume should be pruned now as flag set true
		Expect(len(systemPruneResponse.VolumePruneReports)).To(Equal(1))
	})

	It("podman system prune running alpine container volume prune --filter", func() {
		// Start a pod and leave it running
		_, err := pods.Start(bt.conn, newpod, nil)
		Expect(err).To(BeNil())

		// Start and stop a container to enter in exited state.
		var name = "top"
		_, err = bt.RunTopContainer(&name, nil)
		Expect(err).To(BeNil())
		err = containers.Stop(bt.conn, name, nil)
		Expect(err).To(BeNil())

		// Start second container and leave in running
		var name2 = "top2"
		_, err = bt.RunTopContainer(&name2, nil)
		Expect(err).To(BeNil())

		// Adding an unused volume should work
		_, err = volumes.Create(bt.conn, entities.VolumeCreateOptions{}, nil)
		Expect(err).To(BeNil())

		// Adding an unused volume with label should work
		_, err = volumes.Create(bt.conn, entities.VolumeCreateOptions{Label: map[string]string{
			"label1": "value1",
		}}, nil)
		Expect(err).To(BeNil())

		f := make(map[string][]string)
		f["label"] = []string{"label1=idontmatch"}

		options := new(system.PruneOptions).WithAll(true).WithVolumes(true).WithFilters(f)
		systemPruneResponse, err := system.Prune(bt.conn, options)
		Expect(err).To(BeNil())
		Expect(len(systemPruneResponse.PodPruneReport)).To(Equal(0))
		Expect(len(systemPruneResponse.ContainerPruneReports)).To(Equal(0))
		Expect(len(systemPruneResponse.ImagePruneReports)).To(Equal(0))
		// Alpine image should not be pruned as used by running container
		Expect(reports.PruneReportsIds(systemPruneResponse.ImagePruneReports)).
			ToNot(ContainElement("docker.io/library/alpine:latest"))
		// Volume shouldn't be pruned because the PruneOptions filters doesn't match
		Expect(len(systemPruneResponse.VolumePruneReports)).To(Equal(0))

		// Fix filter and re prune
		f["label"] = []string{"label1=value1"}
		options = new(system.PruneOptions).WithAll(true).WithVolumes(true).WithFilters(f)
		systemPruneResponse, err = system.Prune(bt.conn, options)
		Expect(err).To(BeNil())

		// Volume should be pruned because the PruneOptions filters now match
		Expect(len(systemPruneResponse.VolumePruneReports)).To(Equal(1))
	})
})
