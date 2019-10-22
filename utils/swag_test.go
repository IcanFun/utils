package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggo/swag"
)

// Test ParseParamComment
func TestParseParamCommentByPathType(t *testing.T) {
	comment := `@Param  param body form.AccountInfoForm true "param" default({o61okidfd3d9zr7xjg1dzxzkde})`
	operation := swag.NewOperation()
	err := operation.ParseComment(comment, nil)

	assert.NoError(t, err)
	b, _ := json.MarshalIndent(operation, "", "    ")
	expected := `{
    "parameters": [
        {
            "type": "integer",
            "description": "Some ID",
            "name": "some_id",
            "in": "path",
            "required": true
        }
    ]
}`
	assert.Equal(t, expected, string(b))
}
