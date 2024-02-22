// Code generated by "stringer -type TokType -trimprefix Tok"; DO NOT EDIT.

package lisp1_5

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[tokenError-0]
	_ = x[tokenEOF-1]
	_ = x[tokenAtom-2]
	_ = x[tokenConst-3]
	_ = x[tokenNumber-4]
	_ = x[tokenLpar-5]
	_ = x[tokenRpar-6]
	_ = x[tokenDot-7]
	_ = x[tokenChar-8]
	_ = x[tokenQuote-9]
	_ = x[tokenNewline-10]
}

const _TokType_name = "tokenErrortokenEOFtokenAtomtokenConsttokenNumbertokenLpartokenRpartokenDottokenChartokenQuotetokenNewline"

var _TokType_index = [...]uint8{0, 10, 18, 27, 37, 48, 57, 66, 74, 83, 93, 105}

func (i TokType) String() string {
	if i < 0 || i >= TokType(len(_TokType_index)-1) {
		return "TokType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokType_name[_TokType_index[i]:_TokType_index[i+1]]
}