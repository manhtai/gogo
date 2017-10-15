package findunique_test

import (
	. "math"
	"math/rand"

	. "github.com/manhtai/gogo/findunique"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FindUniq", func() {
	It("should work for some basic cases", func() {
		Expect(FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0})).To(Equal(float32(2)))
		Expect(FindUniq([]float32{0, 0, 0.55, 0, 0})).To(Equal(float32(0.55)))
	})

	It("should work with some easy cases", func() {
		// Basic tests (shuffled)
		Expect(FindUniq([]float32{4, 4, 4, 3, 4, 4, 4, 4})).To(Equal(float32(3)))
		Expect(FindUniq([]float32{5, 5, 5, 5, 4, 5, 5, 5})).To(Equal(float32(4)))
		Expect(FindUniq([]float32{6, 6, 6, 6, 6, 5, 6, 6})).To(Equal(float32(5)))
		Expect(FindUniq([]float32{7, 7, 7, 7, 7, 7, 6, 7})).To(Equal(float32(6)))
		// The last item
		Expect(FindUniq([]float32{8, 8, 8, 8, 8, 8, 8, 7})).To(Equal(float32(7)))
		Expect(FindUniq([]float32{3, 3, 2, 3, 3, 3, 3, 3})).To(Equal(float32(2)))
		Expect(FindUniq([]float32{2, 1, 2, 2, 2, 2, 2, 2})).To(Equal(float32(1)))
		// The first item
		Expect(FindUniq([]float32{0, 1, 1, 1, 1, 1, 1, 1})).To(Equal(float32(0)))
	})

	It("should work with random cases", func() {
		ans := float32(rand.Intn(100))
		mass := float32(rand.Intn(100))
		Expect(FindUniq(generateTestArr(ans, mass, 30))).To(Equal(ans))
	})

	It("should work with some edge cases", func() {
		// Very big number
		Expect(FindUniq(generateTestArr(MaxFloat32, MaxFloat32/2, 1000))).To(Equal(float32(MaxFloat32)))
		// Negative number
		Expect(FindUniq(generateTestArr(float32(-30), 0.0, 10))).To(Equal(float32(-30)))
		// Very big array
		ans := float32(rand.Intn(100))
		mass := float32(rand.Intn(100))
		Expect(FindUniq(generateTestArr(ans, mass, MaxInt32/32))).To(Equal(ans))
	})

})

func generateTestArr(answer, mass float32, length int) []float32 {
	arr := make([]float32, length)
	answerIndex := rand.Intn(length - 1)
	for i := range arr {
		arr[i] = mass
		if i == answerIndex {
			arr[i] = answer
		}
	}

	return arr
}
