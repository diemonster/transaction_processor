package circular_buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	size := 100
	q := New(size)

	for i := 0; i < size; i++ {
		assert.NoError(t, q.Add(i))
	}
	assert.Error(t, q.Add(0))
	assert.Equal(t, errFull, q.Add(0))

	for i := 0; i < size; i++ {
		v, err := q.Delete()
		assert.Equal(t, i, v.(int))
		assert.NoError(t, err)
	}

	_, err := q.Delete()
	assert.Error(t, err)
	assert.Equal(t, errNoTask, err)
}

func BenchmarkCircularBufferAddDelete(b *testing.B) {
	q := New(b.N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = q.Add(i)
		_, _ = q.Delete()
	}
}

func BenchmarkCircularBufferAdd(b *testing.B) {
	q := New(b.N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = q.Add(i)
	}
}

func BenchmarkCircularBufferDelete(b *testing.B) {
	q := New(b.N)

	for i := 0; i < b.N; i++ {
		_ = q.Add(i)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = q.Delete()
	}
}
