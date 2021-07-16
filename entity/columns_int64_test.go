// Code generated by go generate; DO NOT EDIT
// This file is generated by go genrated at 2021-07-15 15:55:19.25274366 +0800 CST m=+0.002663389

//Package entity defines entities used in sdk
package entity 

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/milvus-io/milvus-sdk-go/v2/internal/proto/schema"
)

func TestColumnInt64(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	columnName := fmt.Sprintf("column_Int64_%d", rand.Int())
	columnLen := 8 + rand.Intn(10)

	v := make([]int64, columnLen)
	column := NewColumnInt64(columnName, v)

	t.Run("test meta", func(t *testing.T) {
		ft := FieldTypeInt64
		assert.Equal(t, "Int64", ft.Name())
		assert.Equal(t, "int64", ft.String())
		pbName, pbType := ft.PbFieldType()
		assert.Equal(t, "Long", pbName)
		assert.Equal(t, "int64", pbType)
	})

	t.Run("test column attribute", func(t *testing.T) {
		assert.Equal(t, columnName, column.Name())
		assert.Equal(t, FieldTypeInt64, column.Type())
		assert.Equal(t, columnLen, column.Len())
	})

	t.Run("test column field data", func(t *testing.T) {
		fd := column.FieldData()
		assert.NotNil(t, fd)
		assert.Equal(t, fd.GetFieldName(), columnName)
	})

	t.Run("test column value by idx", func(t *testing.T) {
		_, err := column.ValueByIdx(-1)
		assert.NotNil(t, err)
		_, err = column.ValueByIdx(columnLen)
		assert.NotNil(t, err)
		for i := 0; i < columnLen; i++ {
			v, err := column.ValueByIdx(i)
			assert.Nil(t, err)
			assert.Equal(t, column.values[i], v)
		}
	})
}

func TestFieldDataInt64Column(t *testing.T) {
	len := rand.Intn(10) + 8
	name := fmt.Sprintf("fd_Int64_%d", rand.Int())
	fd := &schema.FieldData{
		Type: schema.DataType_Int64,
		FieldName: name,
	}

	t.Run("normal usage", func(t *testing.T) {
		fd.Field = &schema.FieldData_Scalars{
			Scalars: &schema.ScalarField{
				Data: &schema.ScalarField_LongData{
					LongData: &schema.LongArray{
						Data: make([]int64, len),
					},
				},
			},
		}
		column, err:= FieldDataColumn(fd, 0, len)
		assert.Nil(t, err)
		assert.NotNil(t, column)
 
		assert.Equal(t, name, column.Name())
		assert.Equal(t, len, column.Len())
		assert.Equal(t, FieldTypeInt64, column.Type())
	})

	
	t.Run("nil data", func(t *testing.T) {
		fd.Field = nil
		_, err := FieldDataColumn(fd, 0, len)
		assert.NotNil(t, err)
	})
}