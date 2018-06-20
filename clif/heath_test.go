package clif_test

import (
	"fmt"
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "zanini.me/zfsmon/clif"
)

var _ = Describe("NewHealthFromCliOutput", func() {
	healthMap := map[string]Health{
		"ONLINE":   Online,
		"DEGRADED": Degraded,
		"FAULTED":  Faulted,
		"OFFLINE":  Offline,
		"UNAVAIL":  Unavail,
		"REMOVED":  Removed,
	}

	for k, v := range healthMap {
		It(fmt.Sprintf("parses %s correctly", k), func() {
			h, err := NewHealthFromCliOutput(k)

			Expect(*h).To(Equal(v))
			Expect(err).To(BeNil())
		})
	}

	It("returns ErrUnknownValue for any other case", func() {
		h, err := NewHealthFromCliOutput(fmt.Sprintf("any %d", rand.Int()))

		Expect(h).To(BeNil())
		Expect(err).To(Equal(ErrUnknownValue))
	})
})
