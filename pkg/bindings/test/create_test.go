package bindings_test

import (
	"time"

	"github.com/hanks177/podman/v4/pkg/bindings/containers"
	"github.com/hanks177/podman/v4/pkg/specgen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Create containers ", func() {
	var (
		bt *bindingTest
		s  *gexec.Session
	)

	BeforeEach(func() {
		bt = newBindingTest()
		bt.RestoreImagesFromCache()
		s = bt.startAPIService()
		time.Sleep(1 * time.Second)
		err := bt.NewConnection()
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		s.Kill()
		bt.cleanup()
	})

	It("create a container running top", func() {
		s := specgen.NewSpecGenerator(alpine.name, false)
		s.Command = []string{"top"}
		s.Terminal = true
		s.Name = "top"
		ctr, err := containers.CreateWithSpec(bt.conn, s, nil)
		Expect(err).To(BeNil())
		data, err := containers.Inspect(bt.conn, ctr.ID, nil)
		Expect(err).To(BeNil())
		Expect(data.Name).To(Equal("top"))
		err = containers.Start(bt.conn, ctr.ID, nil)
		Expect(err).To(BeNil())
		data, err = containers.Inspect(bt.conn, ctr.ID, nil)
		Expect(err).To(BeNil())
		Expect(data.State.Status).To(Equal("running"))
	})

})
