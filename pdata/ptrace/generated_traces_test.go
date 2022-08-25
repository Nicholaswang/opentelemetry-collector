// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package ptrace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/data"
	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestResourceSpansSlice(t *testing.T) {
	es := NewResourceSpansSlice()
	assert.Equal(t, 0, es.Len())
	es = newResourceSpansSlice(&[]*otlptrace.ResourceSpans{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newResourceSpans(&otlptrace.ResourceSpans{})
	testVal := ResourceSpans(internal.GenerateTestResourceSpans())
	assert.Equal(t, 7, cap(*es.getOrig()))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		internal.FillTestResourceSpans(internal.ResourceSpans(el))
		assert.Equal(t, testVal, el)
	}
}

func TestResourceSpansSlice_CopyTo(t *testing.T) {
	dest := NewResourceSpansSlice()
	// Test CopyTo to empty
	NewResourceSpansSlice().CopyTo(dest)
	assert.Equal(t, NewResourceSpansSlice(), dest)

	// Test CopyTo larger slice
	ResourceSpansSlice(internal.GenerateTestResourceSpansSlice()).CopyTo(dest)
	assert.Equal(t, ResourceSpansSlice(internal.GenerateTestResourceSpansSlice()), dest)

	// Test CopyTo same size slice
	ResourceSpansSlice(internal.GenerateTestResourceSpansSlice()).CopyTo(dest)
	assert.Equal(t, ResourceSpansSlice(internal.GenerateTestResourceSpansSlice()), dest)
}

func TestResourceSpansSlice_EnsureCapacity(t *testing.T) {
	es := ResourceSpansSlice(internal.GenerateTestResourceSpansSlice())
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.ResourceSpans]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.ResourceSpans]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.ResourceSpans]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.getOrig()))
	foundEs = make(map[*otlptrace.ResourceSpans]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)
}

func TestResourceSpansSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := ResourceSpansSlice(internal.GenerateTestResourceSpansSlice())
	dest := NewResourceSpansSlice()
	src := ResourceSpansSlice(internal.GenerateTestResourceSpansSlice())
	src.MoveAndAppendTo(dest)
	assert.Equal(t, ResourceSpansSlice(internal.GenerateTestResourceSpansSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, ResourceSpansSlice(internal.GenerateTestResourceSpansSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	ResourceSpansSlice(internal.GenerateTestResourceSpansSlice()).MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestResourceSpansSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewResourceSpansSlice()
	emptySlice.RemoveIf(func(el ResourceSpans) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := ResourceSpansSlice(internal.GenerateTestResourceSpansSlice())
	pos := 0
	filtered.RemoveIf(func(el ResourceSpans) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestResourceSpans_MoveTo(t *testing.T) {
	ms := ResourceSpans(internal.GenerateTestResourceSpans())
	dest := NewResourceSpans()
	ms.MoveTo(dest)
	assert.Equal(t, NewResourceSpans(), ms)
	assert.Equal(t, ResourceSpans(internal.GenerateTestResourceSpans()), dest)
}

func TestResourceSpans_CopyTo(t *testing.T) {
	ms := NewResourceSpans()
	orig := NewResourceSpans()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = ResourceSpans(internal.GenerateTestResourceSpans())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestResourceSpans_Resource(t *testing.T) {
	ms := NewResourceSpans()
	internal.FillTestResource(internal.Resource(ms.Resource()))
	assert.Equal(t, pcommon.Resource(internal.GenerateTestResource()), ms.Resource())
}

func TestResourceSpans_SchemaUrl(t *testing.T) {
	ms := NewResourceSpans()
	assert.Equal(t, "", ms.SchemaUrl())
	ms.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	assert.Equal(t, "https://opentelemetry.io/schemas/1.5.0", ms.SchemaUrl())
}

func TestResourceSpans_ScopeSpans(t *testing.T) {
	ms := NewResourceSpans()
	assert.Equal(t, NewScopeSpansSlice(), ms.ScopeSpans())
	internal.FillTestScopeSpansSlice(internal.ScopeSpansSlice(ms.ScopeSpans()))
	assert.Equal(t, ScopeSpansSlice(internal.GenerateTestScopeSpansSlice()), ms.ScopeSpans())
}

func TestScopeSpansSlice(t *testing.T) {
	es := NewScopeSpansSlice()
	assert.Equal(t, 0, es.Len())
	es = newScopeSpansSlice(&[]*otlptrace.ScopeSpans{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newScopeSpans(&otlptrace.ScopeSpans{})
	testVal := ScopeSpans(internal.GenerateTestScopeSpans())
	assert.Equal(t, 7, cap(*es.getOrig()))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		internal.FillTestScopeSpans(internal.ScopeSpans(el))
		assert.Equal(t, testVal, el)
	}
}

func TestScopeSpansSlice_CopyTo(t *testing.T) {
	dest := NewScopeSpansSlice()
	// Test CopyTo to empty
	NewScopeSpansSlice().CopyTo(dest)
	assert.Equal(t, NewScopeSpansSlice(), dest)

	// Test CopyTo larger slice
	ScopeSpansSlice(internal.GenerateTestScopeSpansSlice()).CopyTo(dest)
	assert.Equal(t, ScopeSpansSlice(internal.GenerateTestScopeSpansSlice()), dest)

	// Test CopyTo same size slice
	ScopeSpansSlice(internal.GenerateTestScopeSpansSlice()).CopyTo(dest)
	assert.Equal(t, ScopeSpansSlice(internal.GenerateTestScopeSpansSlice()), dest)
}

func TestScopeSpansSlice_EnsureCapacity(t *testing.T) {
	es := ScopeSpansSlice(internal.GenerateTestScopeSpansSlice())
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.ScopeSpans]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.ScopeSpans]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.ScopeSpans]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.getOrig()))
	foundEs = make(map[*otlptrace.ScopeSpans]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)
}

func TestScopeSpansSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := ScopeSpansSlice(internal.GenerateTestScopeSpansSlice())
	dest := NewScopeSpansSlice()
	src := ScopeSpansSlice(internal.GenerateTestScopeSpansSlice())
	src.MoveAndAppendTo(dest)
	assert.Equal(t, ScopeSpansSlice(internal.GenerateTestScopeSpansSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, ScopeSpansSlice(internal.GenerateTestScopeSpansSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	ScopeSpansSlice(internal.GenerateTestScopeSpansSlice()).MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestScopeSpansSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewScopeSpansSlice()
	emptySlice.RemoveIf(func(el ScopeSpans) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := ScopeSpansSlice(internal.GenerateTestScopeSpansSlice())
	pos := 0
	filtered.RemoveIf(func(el ScopeSpans) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestScopeSpans_MoveTo(t *testing.T) {
	ms := ScopeSpans(internal.GenerateTestScopeSpans())
	dest := NewScopeSpans()
	ms.MoveTo(dest)
	assert.Equal(t, NewScopeSpans(), ms)
	assert.Equal(t, ScopeSpans(internal.GenerateTestScopeSpans()), dest)
}

func TestScopeSpans_CopyTo(t *testing.T) {
	ms := NewScopeSpans()
	orig := NewScopeSpans()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = ScopeSpans(internal.GenerateTestScopeSpans())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestScopeSpans_Scope(t *testing.T) {
	ms := NewScopeSpans()
	internal.FillTestInstrumentationScope(internal.InstrumentationScope(ms.Scope()))
	assert.Equal(t, pcommon.InstrumentationScope(internal.GenerateTestInstrumentationScope()), ms.Scope())
}

func TestScopeSpans_SchemaUrl(t *testing.T) {
	ms := NewScopeSpans()
	assert.Equal(t, "", ms.SchemaUrl())
	ms.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	assert.Equal(t, "https://opentelemetry.io/schemas/1.5.0", ms.SchemaUrl())
}

func TestScopeSpans_Spans(t *testing.T) {
	ms := NewScopeSpans()
	assert.Equal(t, NewSpanSlice(), ms.Spans())
	internal.FillTestSpanSlice(internal.SpanSlice(ms.Spans()))
	assert.Equal(t, SpanSlice(internal.GenerateTestSpanSlice()), ms.Spans())
}

func TestSpanSlice(t *testing.T) {
	es := NewSpanSlice()
	assert.Equal(t, 0, es.Len())
	es = newSpanSlice(&[]*otlptrace.Span{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newSpan(&otlptrace.Span{})
	testVal := Span(internal.GenerateTestSpan())
	assert.Equal(t, 7, cap(*es.getOrig()))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		internal.FillTestSpan(internal.Span(el))
		assert.Equal(t, testVal, el)
	}
}

func TestSpanSlice_CopyTo(t *testing.T) {
	dest := NewSpanSlice()
	// Test CopyTo to empty
	NewSpanSlice().CopyTo(dest)
	assert.Equal(t, NewSpanSlice(), dest)

	// Test CopyTo larger slice
	SpanSlice(internal.GenerateTestSpanSlice()).CopyTo(dest)
	assert.Equal(t, SpanSlice(internal.GenerateTestSpanSlice()), dest)

	// Test CopyTo same size slice
	SpanSlice(internal.GenerateTestSpanSlice()).CopyTo(dest)
	assert.Equal(t, SpanSlice(internal.GenerateTestSpanSlice()), dest)
}

func TestSpanSlice_EnsureCapacity(t *testing.T) {
	es := SpanSlice(internal.GenerateTestSpanSlice())
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.Span]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.Span]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.Span]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.getOrig()))
	foundEs = make(map[*otlptrace.Span]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)
}

func TestSpanSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := SpanSlice(internal.GenerateTestSpanSlice())
	dest := NewSpanSlice()
	src := SpanSlice(internal.GenerateTestSpanSlice())
	src.MoveAndAppendTo(dest)
	assert.Equal(t, SpanSlice(internal.GenerateTestSpanSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, SpanSlice(internal.GenerateTestSpanSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	SpanSlice(internal.GenerateTestSpanSlice()).MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestSpanSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewSpanSlice()
	emptySlice.RemoveIf(func(el Span) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := SpanSlice(internal.GenerateTestSpanSlice())
	pos := 0
	filtered.RemoveIf(func(el Span) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestSpan_MoveTo(t *testing.T) {
	ms := Span(internal.GenerateTestSpan())
	dest := NewSpan()
	ms.MoveTo(dest)
	assert.Equal(t, NewSpan(), ms)
	assert.Equal(t, Span(internal.GenerateTestSpan()), dest)
}

func TestSpan_CopyTo(t *testing.T) {
	ms := NewSpan()
	orig := NewSpan()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = Span(internal.GenerateTestSpan())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestSpan_TraceID(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.TraceID(internal.NewTraceID(data.NewTraceID([16]byte{}))), ms.TraceID())
	testValTraceID := pcommon.TraceID(internal.NewTraceID(data.NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})))
	ms.SetTraceID(testValTraceID)
	assert.Equal(t, testValTraceID, ms.TraceID())
}

func TestSpan_SpanID(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.SpanID(internal.NewSpanID(data.NewSpanID([8]byte{}))), ms.SpanID())
	testValSpanID := pcommon.SpanID(internal.NewSpanID(data.NewSpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8})))
	ms.SetSpanID(testValSpanID)
	assert.Equal(t, testValSpanID, ms.SpanID())
}

func TestSpan_TraceState(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, TraceState(""), ms.TraceState())
	testValTraceState := TraceState("congo=congos")
	ms.SetTraceState(testValTraceState)
	assert.Equal(t, testValTraceState, ms.TraceState())
}

func TestSpan_ParentSpanID(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.SpanID(internal.NewSpanID(data.NewSpanID([8]byte{}))), ms.ParentSpanID())
	testValParentSpanID := pcommon.SpanID(internal.NewSpanID(data.NewSpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1})))
	ms.SetParentSpanID(testValParentSpanID)
	assert.Equal(t, testValParentSpanID, ms.ParentSpanID())
}

func TestSpan_Name(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, "", ms.Name())
	ms.SetName("test_name")
	assert.Equal(t, "test_name", ms.Name())
}

func TestSpan_Kind(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, SpanKind(otlptrace.Span_SpanKind(0)), ms.Kind())
	testValKind := SpanKind(otlptrace.Span_SpanKind(3))
	ms.SetKind(testValKind)
	assert.Equal(t, testValKind, ms.Kind())
}

func TestSpan_StartTimestamp(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.Timestamp(0), ms.StartTimestamp())
	testValStartTimestamp := pcommon.Timestamp(1234567890)
	ms.SetStartTimestamp(testValStartTimestamp)
	assert.Equal(t, testValStartTimestamp, ms.StartTimestamp())
}

func TestSpan_EndTimestamp(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.Timestamp(0), ms.EndTimestamp())
	testValEndTimestamp := pcommon.Timestamp(1234567890)
	ms.SetEndTimestamp(testValEndTimestamp)
	assert.Equal(t, testValEndTimestamp, ms.EndTimestamp())
}

func TestSpan_Attributes(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, pcommon.NewMap(), ms.Attributes())
	internal.FillTestMap(internal.Map(ms.Attributes()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.Attributes())
}

func TestSpan_DroppedAttributesCount(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, uint32(0), ms.DroppedAttributesCount())
	ms.SetDroppedAttributesCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedAttributesCount())
}

func TestSpan_Events(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, NewSpanEventSlice(), ms.Events())
	internal.FillTestSpanEventSlice(internal.SpanEventSlice(ms.Events()))
	assert.Equal(t, SpanEventSlice(internal.GenerateTestSpanEventSlice()), ms.Events())
}

func TestSpan_DroppedEventsCount(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, uint32(0), ms.DroppedEventsCount())
	ms.SetDroppedEventsCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedEventsCount())
}

func TestSpan_Links(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, NewSpanLinkSlice(), ms.Links())
	internal.FillTestSpanLinkSlice(internal.SpanLinkSlice(ms.Links()))
	assert.Equal(t, SpanLinkSlice(internal.GenerateTestSpanLinkSlice()), ms.Links())
}

func TestSpan_DroppedLinksCount(t *testing.T) {
	ms := NewSpan()
	assert.Equal(t, uint32(0), ms.DroppedLinksCount())
	ms.SetDroppedLinksCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedLinksCount())
}

func TestSpan_Status(t *testing.T) {
	ms := NewSpan()
	internal.FillTestSpanStatus(internal.SpanStatus(ms.Status()))
	assert.Equal(t, SpanStatus(internal.GenerateTestSpanStatus()), ms.Status())
}

func TestSpanEventSlice(t *testing.T) {
	es := NewSpanEventSlice()
	assert.Equal(t, 0, es.Len())
	es = newSpanEventSlice(&[]*otlptrace.Span_Event{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newSpanEvent(&otlptrace.Span_Event{})
	testVal := SpanEvent(internal.GenerateTestSpanEvent())
	assert.Equal(t, 7, cap(*es.getOrig()))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		internal.FillTestSpanEvent(internal.SpanEvent(el))
		assert.Equal(t, testVal, el)
	}
}

func TestSpanEventSlice_CopyTo(t *testing.T) {
	dest := NewSpanEventSlice()
	// Test CopyTo to empty
	NewSpanEventSlice().CopyTo(dest)
	assert.Equal(t, NewSpanEventSlice(), dest)

	// Test CopyTo larger slice
	SpanEventSlice(internal.GenerateTestSpanEventSlice()).CopyTo(dest)
	assert.Equal(t, SpanEventSlice(internal.GenerateTestSpanEventSlice()), dest)

	// Test CopyTo same size slice
	SpanEventSlice(internal.GenerateTestSpanEventSlice()).CopyTo(dest)
	assert.Equal(t, SpanEventSlice(internal.GenerateTestSpanEventSlice()), dest)
}

func TestSpanEventSlice_EnsureCapacity(t *testing.T) {
	es := SpanEventSlice(internal.GenerateTestSpanEventSlice())
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.Span_Event]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.Span_Event]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.Span_Event]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.getOrig()))
	foundEs = make(map[*otlptrace.Span_Event]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)
}

func TestSpanEventSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := SpanEventSlice(internal.GenerateTestSpanEventSlice())
	dest := NewSpanEventSlice()
	src := SpanEventSlice(internal.GenerateTestSpanEventSlice())
	src.MoveAndAppendTo(dest)
	assert.Equal(t, SpanEventSlice(internal.GenerateTestSpanEventSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, SpanEventSlice(internal.GenerateTestSpanEventSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	SpanEventSlice(internal.GenerateTestSpanEventSlice()).MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestSpanEventSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewSpanEventSlice()
	emptySlice.RemoveIf(func(el SpanEvent) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := SpanEventSlice(internal.GenerateTestSpanEventSlice())
	pos := 0
	filtered.RemoveIf(func(el SpanEvent) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestSpanEvent_MoveTo(t *testing.T) {
	ms := SpanEvent(internal.GenerateTestSpanEvent())
	dest := NewSpanEvent()
	ms.MoveTo(dest)
	assert.Equal(t, NewSpanEvent(), ms)
	assert.Equal(t, SpanEvent(internal.GenerateTestSpanEvent()), dest)
}

func TestSpanEvent_CopyTo(t *testing.T) {
	ms := NewSpanEvent()
	orig := NewSpanEvent()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = SpanEvent(internal.GenerateTestSpanEvent())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestSpanEvent_Timestamp(t *testing.T) {
	ms := NewSpanEvent()
	assert.Equal(t, pcommon.Timestamp(0), ms.Timestamp())
	testValTimestamp := pcommon.Timestamp(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.Equal(t, testValTimestamp, ms.Timestamp())
}

func TestSpanEvent_Name(t *testing.T) {
	ms := NewSpanEvent()
	assert.Equal(t, "", ms.Name())
	ms.SetName("test_name")
	assert.Equal(t, "test_name", ms.Name())
}

func TestSpanEvent_Attributes(t *testing.T) {
	ms := NewSpanEvent()
	assert.Equal(t, pcommon.NewMap(), ms.Attributes())
	internal.FillTestMap(internal.Map(ms.Attributes()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.Attributes())
}

func TestSpanEvent_DroppedAttributesCount(t *testing.T) {
	ms := NewSpanEvent()
	assert.Equal(t, uint32(0), ms.DroppedAttributesCount())
	ms.SetDroppedAttributesCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedAttributesCount())
}

func TestSpanLinkSlice(t *testing.T) {
	es := NewSpanLinkSlice()
	assert.Equal(t, 0, es.Len())
	es = newSpanLinkSlice(&[]*otlptrace.Span_Link{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newSpanLink(&otlptrace.Span_Link{})
	testVal := SpanLink(internal.GenerateTestSpanLink())
	assert.Equal(t, 7, cap(*es.getOrig()))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		internal.FillTestSpanLink(internal.SpanLink(el))
		assert.Equal(t, testVal, el)
	}
}

func TestSpanLinkSlice_CopyTo(t *testing.T) {
	dest := NewSpanLinkSlice()
	// Test CopyTo to empty
	NewSpanLinkSlice().CopyTo(dest)
	assert.Equal(t, NewSpanLinkSlice(), dest)

	// Test CopyTo larger slice
	SpanLinkSlice(internal.GenerateTestSpanLinkSlice()).CopyTo(dest)
	assert.Equal(t, SpanLinkSlice(internal.GenerateTestSpanLinkSlice()), dest)

	// Test CopyTo same size slice
	SpanLinkSlice(internal.GenerateTestSpanLinkSlice()).CopyTo(dest)
	assert.Equal(t, SpanLinkSlice(internal.GenerateTestSpanLinkSlice()), dest)
}

func TestSpanLinkSlice_EnsureCapacity(t *testing.T) {
	es := SpanLinkSlice(internal.GenerateTestSpanLinkSlice())
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.Span_Link]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.Span_Link]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.Span_Link]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.getOrig()))
	foundEs = make(map[*otlptrace.Span_Link]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)
}

func TestSpanLinkSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := SpanLinkSlice(internal.GenerateTestSpanLinkSlice())
	dest := NewSpanLinkSlice()
	src := SpanLinkSlice(internal.GenerateTestSpanLinkSlice())
	src.MoveAndAppendTo(dest)
	assert.Equal(t, SpanLinkSlice(internal.GenerateTestSpanLinkSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, SpanLinkSlice(internal.GenerateTestSpanLinkSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	SpanLinkSlice(internal.GenerateTestSpanLinkSlice()).MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestSpanLinkSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewSpanLinkSlice()
	emptySlice.RemoveIf(func(el SpanLink) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := SpanLinkSlice(internal.GenerateTestSpanLinkSlice())
	pos := 0
	filtered.RemoveIf(func(el SpanLink) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestSpanLink_MoveTo(t *testing.T) {
	ms := SpanLink(internal.GenerateTestSpanLink())
	dest := NewSpanLink()
	ms.MoveTo(dest)
	assert.Equal(t, NewSpanLink(), ms)
	assert.Equal(t, SpanLink(internal.GenerateTestSpanLink()), dest)
}

func TestSpanLink_CopyTo(t *testing.T) {
	ms := NewSpanLink()
	orig := NewSpanLink()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = SpanLink(internal.GenerateTestSpanLink())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestSpanLink_TraceID(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, pcommon.TraceID(internal.NewTraceID(data.NewTraceID([16]byte{}))), ms.TraceID())
	testValTraceID := pcommon.TraceID(internal.NewTraceID(data.NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})))
	ms.SetTraceID(testValTraceID)
	assert.Equal(t, testValTraceID, ms.TraceID())
}

func TestSpanLink_SpanID(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, pcommon.SpanID(internal.NewSpanID(data.NewSpanID([8]byte{}))), ms.SpanID())
	testValSpanID := pcommon.SpanID(internal.NewSpanID(data.NewSpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8})))
	ms.SetSpanID(testValSpanID)
	assert.Equal(t, testValSpanID, ms.SpanID())
}

func TestSpanLink_TraceState(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, TraceState(""), ms.TraceState())
	testValTraceState := TraceState("congo=congos")
	ms.SetTraceState(testValTraceState)
	assert.Equal(t, testValTraceState, ms.TraceState())
}

func TestSpanLink_Attributes(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, pcommon.NewMap(), ms.Attributes())
	internal.FillTestMap(internal.Map(ms.Attributes()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.Attributes())
}

func TestSpanLink_DroppedAttributesCount(t *testing.T) {
	ms := NewSpanLink()
	assert.Equal(t, uint32(0), ms.DroppedAttributesCount())
	ms.SetDroppedAttributesCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedAttributesCount())
}

func TestSpanStatus_MoveTo(t *testing.T) {
	ms := SpanStatus(internal.GenerateTestSpanStatus())
	dest := NewSpanStatus()
	ms.MoveTo(dest)
	assert.Equal(t, NewSpanStatus(), ms)
	assert.Equal(t, SpanStatus(internal.GenerateTestSpanStatus()), dest)
}

func TestSpanStatus_CopyTo(t *testing.T) {
	ms := NewSpanStatus()
	orig := NewSpanStatus()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = SpanStatus(internal.GenerateTestSpanStatus())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestSpanStatus_Code(t *testing.T) {
	ms := NewSpanStatus()
	assert.Equal(t, StatusCode(0), ms.Code())
	testValCode := StatusCode(1)
	ms.SetCode(testValCode)
	assert.Equal(t, testValCode, ms.Code())
}

func TestSpanStatus_Message(t *testing.T) {
	ms := NewSpanStatus()
	assert.Equal(t, "", ms.Message())
	ms.SetMessage("cancelled")
	assert.Equal(t, "cancelled", ms.Message())
}
