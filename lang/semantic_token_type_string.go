// Code generated by "stringer -type=SemanticTokenType -output=semantic_token_type_string.go"; DO NOT EDIT.

package lang

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TokenNil-0]
	_ = x[TokenAttrName-1]
	_ = x[TokenBlockType-2]
	_ = x[TokenBlockLabel-3]
	_ = x[TokenBool-4]
	_ = x[TokenString-5]
	_ = x[TokenNumber-6]
	_ = x[TokenObjectKey-7]
	_ = x[TokenMapKey-8]
	_ = x[TokenKeyword-9]
	_ = x[TokenTraversalStep-10]
}

const _SemanticTokenType_name = "TokenNilTokenAttrNameTokenBlockTypeTokenBlockLabelTokenBoolTokenStringTokenNumberTokenObjectKeyTokenMapKeyTokenKeywordTokenTraversalStep"

var _SemanticTokenType_index = [...]uint8{0, 8, 21, 35, 50, 59, 70, 81, 95, 106, 118, 136}

func (i SemanticTokenType) String() string {
	if i >= SemanticTokenType(len(_SemanticTokenType_index)-1) {
		return "SemanticTokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SemanticTokenType_name[_SemanticTokenType_index[i]:_SemanticTokenType_index[i+1]]
}
