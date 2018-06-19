package clif_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "zanini.me/zfsmon/clif"
)

var _ = Describe("ParseListOutput", func() {
	Specify("returns an empty list when no output is given", func() {
		out := bytes.NewBufferString("")
		Expect(ParseListOutput(out)).To(Equal([]string{}))

		out = bytes.NewBufferString("\n")
		Expect(ParseListOutput(out)).To(Equal([]string{}))
	})

	Specify("returns a list with one element when only one element is returned", func() {
		out := bytes.NewBufferString("my element")
		expected := []string{"my element"}

		Expect(ParseListOutput(out)).To(Equal(expected))

		out = bytes.NewBufferString("my element\n")
		Expect(ParseListOutput(out)).To(Equal(expected))

		out = bytes.NewBufferString("\nmy element\n")
		Expect(ParseListOutput(out)).To(Equal(expected))
	})

	Specify("returns a list with two element when two elements are returned", func() {
		out := bytes.NewBufferString("first\nsecond")
		expected := []string{"first", "second"}

		Expect(ParseListOutput(out)).To(Equal(expected))

		out = bytes.NewBufferString("\nfirst\nsecond")
		Expect(ParseListOutput(out)).To(Equal(expected))

		out = bytes.NewBufferString("\nfirst\nsecond\n")
		Expect(ParseListOutput(out)).To(Equal(expected))
	})
})
