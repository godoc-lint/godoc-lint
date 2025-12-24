// godoc with potential doclink to json.Encoder. // want `text "json\.Encoder" should be replaced with "\[json\.Encoder\]" to link to stdlib type`
package with_import

import (
	bytesAlias "bytes"
	"encoding/json"
	ioAlias "io"
)

// (BG: bad godoc)

var _, _, _ = json.Encoder{}, bytesAlias.Buffer{}, ioAlias.EOF

// godoc with potential doclink to json.Encoder. // want `text "json\.Encoder" should be replaced with "\[json\.Encoder\]" to link to stdlib type`
const AlphaBG = 0

// godoc with potential doclink to json.Encoder and json.Encoder. // want `text "json\.Encoder" should be replaced with "\[json\.Encoder\]" to link to stdlib type \(2 instances\)`
const BravoBG = 0

// godoc with potential doclink to json.Encoder and *json.Encoder. // want `text "json\.Encoder" should be replaced with "\[json\.Encoder\]" to link to stdlib type \(2 instances\)`
const CharlieBG = 0

// godoc with doclink to [json.Encoder] and potential doclink to bytesAlias.Buffer. // want `text "bytesAlias\.Buffer" should be replaced with "\[bytesAlias\.Buffer\]" to link to stdlib type`
const DeltaBG = 0

// godoc with doclink to [json.Encoder] and potential doclink to bytesAlias.Buffer and *bytesAlias.Buffer. // want `text "bytesAlias\.Buffer" should be replaced with "\[bytesAlias\.Buffer\]" to link to stdlib type \(2 instances\)`
const EchoBG = 0

// godoc with potential doclink to json.Encoder and bytesAlias.Buffer. // want `text "json\.Encoder" should be replaced with "\[json\.Encoder\]" to link to stdlib type` `text "bytesAlias\.Buffer" should be replaced with "\[bytesAlias\.Buffer\]" to link to stdlib type`
const FoxtrotBG = 0

// godoc with potential doclink to ioAlias.PipeWriter.Close. // want `text "ioAlias\.PipeWriter\.Close" should be replaced with "\[ioAlias\.PipeWriter\.Close\]" to link to stdlib method`
const GolfBG = 0

// godoc with potential doclink to ioAlias.PipeWriter.Close and ioAlias.PipeWriter.Close. // want `text "ioAlias\.PipeWriter\.Close" should be replaced with "\[ioAlias\.PipeWriter\.Close\]" to link to stdlib method \(2 instances\)`
const HotelBG = 0

// godoc with doclink to [ioAlias.PipeWriter.Close] and potential doclink to bytesAlias.Buffer. // want `text "bytesAlias\.Buffer" should be replaced with "\[bytesAlias\.Buffer\]" to link to stdlib type`
const IndiaBG = 0

// godoc with doclink to [ioAlias.PipeWriter.Close] and potential doclink to bytesAlias.Buffer and *bytesAlias.Buffer. // want `text "bytesAlias\.Buffer" should be replaced with "\[bytesAlias\.Buffer\]" to link to stdlib type \(2 instances\)`
const JulietBG = 0

// godoc with potential doclink to ioAlias.PipeWriter.Close and bytesAlias.Buffer. // want `text "ioAlias\.PipeWriter\.Close" should be replaced with "\[ioAlias\.PipeWriter\.Close\]" to link to stdlib method` `text "bytesAlias\.Buffer" should be replaced with "\[bytesAlias\.Buffer\]" to link to stdlib type`
const KiloBG = 0

// godoc with potential doclinks (un-aliased imports names) to encoding/json.Encoder, bytes.Buffer, and io.PipeWriter.Close. // want `text "encoding/json\.Encoder" should be replaced with "\[encoding/json\.Encoder\]" to link to stdlib type` `text "bytes\.Buffer" should be replaced with "\[bytes\.Buffer\]" to link to stdlib type` `text "io\.PipeWriter\.Close" should be replaced with "\[io\.PipeWriter\.Close\]" to link to stdlib method`
const Lima = 0
