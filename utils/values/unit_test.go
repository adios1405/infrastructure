package values

import "testing"

func TestString(t *testing.T) {
	t.Run("BuildStrings", func(t *testing.T) {
		args := []string{"a", "b", "c"}
		result := BuildStrings(args...)
		if result != "abc" {
			t.Errorf("BuildStrings() = %v, want %v", result, "abc")
		}
	})

	t.Run("BuildStringsWithJoin", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			args := []string{"a", "b", "c"}
			result := BuildStringsWithJoin("-", args...)
			if result != "a-b-c" {
				t.Errorf("BuildStringsWIthJoin() = %v, want %v", result, "a-b-c")
			}
		})

		t.Run("NilParts", func(t *testing.T) {
			var args []string = nil
			result := BuildStringsWithJoin(".", args...)
			if result != "" {
				t.Errorf("BuildStringsWIthJoin() = %v, want %v", result, "(empty string)")
			}
		})

		t.Run("EmptyParts", func(t *testing.T) {
			args := []string{}
			result := BuildStringsWithJoin(".", args...)
			if result != "" {
				t.Errorf("BuildStringsWIthJoin() = %v, want %v", result, "(empty string)")
			}
		})
	})

	t.Run("BuildStringsWithJoinIgnoreEmpty", func(t *testing.T) {
		args := []string{"a", "", "b", "c", ""}
		result := BuildStringsWithJoinIgnoreEmpty("-", args...)
		if result != "a-b-c" {
			t.Errorf("BuildStringsWithJoinIgnoreEmpty() = %v, want %v", result, "a-b-c")
		}
	})

	t.Run("BuildStringsWithReplacement", func(t *testing.T) {
		args := []string{"a", "b", "c"}
		replacement := map[string]string{"a": "1", "b": "2", "c": "3"}
		result := BuildStringsWithReplacement(replacement, args...)
		if result != "123" {
			t.Errorf("BuildStringsWithReplacement() = %v, want %v", result, "123")
		}
	})

	t.Run("BuildStringsWithTemplate", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			args := []string{"a", "b", "c"}
			template := "{1}-{2}-{3}"
			result := BuildStringsWithTemplate(template, args...)
			if result != "a-b-c" {
				t.Errorf("BuildStringsWithTemplate() = %v, want %v", result, "a-b-c")
			}
		})

		t.Run("NilArgs", func(t *testing.T) {
			var args []string = nil
			template := "{1}-{2}-{3}"
			result := BuildStringsWithTemplate(template, args...)
			if result != "{1}-{2}-{3}" {
				t.Errorf("BuildStringsWithTemplate() = %v, want %v", result, "{1}-{2}-{3}")
			}
		})

		t.Run("EmptyArgs", func(t *testing.T) {
			args := []string{}
			template := "{1}-{2}-{3}"
			result := BuildStringsWithTemplate(template, args...)
			if result != "{1}-{2}-{3}" {
				t.Errorf("BuildStringsWithTemplate() = %v, want %v", result, "{1}-{2}-{3}")
			}
		})
	})

	t.Run("StringToInt", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			result := StringToInt("123", 0)
			if result != 123 {
				t.Errorf("StringToInt() = %v, want %v", result, 123)
			}
		})

		t.Run("Empty", func(t *testing.T) {
			result := StringToInt("", 123)
			if result != 123 {
				t.Errorf("StringToInt() = %v, want %v", result, 123)
			}
		})

		t.Run("Error", func(t *testing.T) {
			result := StringToInt("fuck you", 123)
			if result != 0 {
				t.Errorf("StringToInt() = %v, want %v", result, 123)
			}
		})
	})

	t.Run("StringToUint", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			result := StringToUint("123", uint(0))
			if result != 123 {
				t.Errorf("StringToUint() = %v, want %v", result, 123)
			}
		})

		t.Run("Empty", func(t *testing.T) {
			result := StringToUint("", uint(123))
			if result != 123 {
				t.Errorf("StringToUint() = %v, want %v", result, 123)
			}
		})

		t.Run("Error", func(t *testing.T) {
			result := StringToUint("fuck you", uint(123))
			if result != 0 {
				t.Errorf("StringToUint() = %v, want %v", result, 123)
			}
		})
	})

	t.Run("StringToFloat64", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			result := StringToFloat64("123.456", 0.0)
			if result != 123.456 {
				t.Errorf("StringToFloat64() = %v, want %v", result, 123.456)
			}
		})

		t.Run("Empty", func(t *testing.T) {
			result := StringToFloat64("", 123.456)
			if result != 123.456 {
				t.Errorf("StringToFloat64() = %v, want %v", result, 123.456)
			}
		})

		t.Run("Error", func(t *testing.T) {
			result := StringToFloat64("fuck you", 123.456)
			if result != 123.456 {
				t.Errorf("StringToFloat64() = %v, want %v", result, 123.456)
			}
		})
	})

	t.Run("StringToBool", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			result := StringToBool("true", false)
			if result != true {
				t.Errorf("StringToBool() = %v, want %v", result, true)
			}
		})

		t.Run("Empty", func(t *testing.T) {
			result := StringToBool("", true)
			if result != true {
				t.Errorf("StringToBool() = %v, want %v", result, true)
			}
		})

		t.Run("Error", func(t *testing.T) {
			result := StringToBool("fuck", true)
			if result != true {
				t.Errorf("StringToBool() = %v, want %v", result, true)
			}
		})
	})

	t.Run("StringToStringPtr", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			result := StringToStringPtr[string]("123")
			if *result != "123" {
				t.Errorf("StringToStringPtr() = %v, want %v", *result, "123")
			}
		})

		t.Run("Empty", func(t *testing.T) {
			result := StringToStringPtr[string]("")
			if result != nil {
				t.Errorf("StringToStringPtr() = %v, want nil", result)
			}
		})
	})

	t.Run("StringToIntPtr", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			result := StringToIntPtr[int]("123")
			if *result != 123 {
				t.Errorf("StringToIntPtr() = %v, want 123", result)
			}
		})

		t.Run("Empty", func(t *testing.T) {
			val := 1
			result := &val
			result = StringToIntPtr[int]("")
			if result != nil {
				t.Errorf("StringToIntPtr() = %v, want nil", result)
			}
		})

		t.Run("Error", func(t *testing.T) {
			val := 1
			result := &val
			result = StringToIntPtr[int]("fuck you")
			if result != nil {
				t.Errorf("StringToIntPtr() = %v, want nil", result)
			}
		})
	})

	t.Run("StringToUintPtr", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			result := StringToUintPtr[uint]("123")
			if *result != 123 {
				t.Errorf("StringToUintPtr() = %v, want 123", result)
			}
		})

		t.Run("Empty", func(t *testing.T) {
			val := uint(1)
			result := &val
			result = StringToUintPtr[uint]("")
			if result != nil {
				t.Errorf("StringToUintPtr() = %v, want nil", result)
			}
		})

		t.Run("Error", func(t *testing.T) {
			val := uint(1)
			result := &val
			result = StringToUintPtr[uint]("fuck you")
			if result != nil {
				t.Errorf("StringToUintPtr() = %v, want nil", result)
			}
		})
	})

	t.Run("StringToFloat64Ptr", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			result := StringToFloat64Ptr[float64]("123.456")
			if *result != 123.456 {
				t.Errorf("StringToFloat64Ptr() = %v, want 123.456", result)
			}
		})

		t.Run("Empty", func(t *testing.T) {
			val := 123.456
			result := &val
			result = StringToFloat64Ptr[float64]("")
			if result != nil {
				t.Errorf("StringToFloat64Ptr() = %v, want nil", result)
			}
		})

		t.Run("Error", func(t *testing.T) {
			val := 123.456
			result := &val
			result = StringToFloat64Ptr[float64]("fuck you")
			if result != nil {
				t.Errorf("StringToFloat64Ptr() = %v, want nil", result)
			}
		})
	})

	t.Run("StringToBoolPtr", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			result := StringToBoolPtr[bool]("true")
			if *result != true {
				t.Errorf("StringToBoolPtr() = %v, want true", result)
			}
		})

		t.Run("Empty", func(t *testing.T) {
			val := true
			result := &val
			result = StringToBoolPtr[bool]("")
			if result != nil {
				t.Errorf("StringToBoolPtr() = %v, want nil", result)
			}
		})

		t.Run("Error", func(t *testing.T) {
			val := true
			result := &val
			result = StringToBoolPtr[bool]("fuck you")
			if result != nil {
				t.Errorf("StringToBoolPtr() = %v, want nil", result)
			}
		})
	})

	t.Run("StringTemplate", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			tmp := NewStringTemplate("dear {name}, i am :me:", map[string]string{
				"name": "name1",
				"me":   "name2",
			})
			if tmp.Parse() != "dear name1, i am name2" {
				t.Errorf("StringTemplate.Parse() = %v, want %v", tmp.Parse(), "dear name1, i am name2")
			}
		})

		t.Run("LongVariableName", func(t *testing.T) {
			tmp := NewStringTemplate("hello, {1234567890123456789012345678901234567890}", map[string]string{
				"1234567890123456789012345678901234567890": "world",
			})
			if tmp.Parse() != "hello, {1234567890123456789012345678901234567890}" {
				t.Errorf("StringTemplate.Parse() = %v, want %v", tmp.Parse(), "hello, {1234567890123456789012345678901234567890}")
			}
		})

		t.Run("EndSignal", func(t *testing.T) {
			tmp := NewStringTemplate("hello, {who}! {", map[string]string{
				"who": "world",
			})
			if tmp.Parse() != "hello, world! {" {
				t.Errorf("StringTemplate.Parse() = %v, want %v", tmp.Parse(), "hello, world! {")
			}
		})

		t.Run("SignalNotFound", func(t *testing.T) {
			tmp := NewStringTemplate("hello, {who!", map[string]string{
				"who": "world",
			})
			if tmp.Parse() != "hello, {who!" {
				t.Errorf("StringTemplate.Parse() = %v, want %v", tmp.Parse(), "hello, {who!")
			}
		})

		t.Run("VariableNotExist", func(t *testing.T) {
			tmp := NewStringTemplate("hello, {who}!", nil)
			if tmp.Parse() != "hello, {who}!" {
				t.Errorf("StringTemplate.Parse() = %v, want %v", tmp.Parse(), "hello, {who}!")
			}
		})
	})
}

func TestNumber(t *testing.T) {
	t.Run("IntToString", func(t *testing.T) {
		result := IntToString(123)
		if result != "123" {
			t.Errorf("IntToString() = %v, want 123", result)
		}
	})

	t.Run("Int64ToString", func(t *testing.T) {
		result := Int64ToString(int64(123))
		if result != "123" {
			t.Errorf("Int64ToString() = %v, want 123", result)
		}
	})

	t.Run("UintToString", func(t *testing.T) {
		result := UintToString(uint(123))
		if result != "123" {
			t.Errorf("UintToString() = %v, want 123", result)
		}
	})

	t.Run("Uint64ToString", func(t *testing.T) {
		result := Uint64ToString(uint64(123))
		if result != "123" {
			t.Errorf("Uint64ToString() = %v, want 123", result)
		}
	})

	t.Run("Float32ToString", func(t *testing.T) {
		result := Float32ToString(float32(123.456))
		if result != "123.456" {
			t.Errorf("Float32ToString() = %v, want 123.456", result)
		}
	})

	t.Run("Float64ToString", func(t *testing.T) {
		result := Float64ToString(123.456)
		if result != "123.456" {
			t.Errorf("Float64ToString() = %v, want 123.456", result)
		}
	})
}

func TestMultiply(t *testing.T) {
	t.Run("NumberOrStringValueGetString", func(t *testing.T) {
		testCases := []struct {
			Case  string
			Value any
			Want  string
		}{
			{
				Case:  "StringType",
				Value: "string",
				Want:  "string",
			},
			{
				Case:  "IntType",
				Value: 123,
				Want:  "123",
			},
			{
				Case:  "Int8Type",
				Value: int8(123),
				Want:  "123",
			},
			{
				Case:  "Int16Type",
				Value: int16(123),
				Want:  "123",
			},
			{
				Case:  "Int32Type",
				Value: int32(123),
				Want:  "123",
			},
			{
				Case:  "Int64Type",
				Value: int64(123),
				Want:  "123",
			},
			{
				Case:  "UintType",
				Value: uint(233),
				Want:  "233",
			},
			{
				Case:  "Uint8Type",
				Value: uint8(233),
				Want:  "233",
			},
			{
				Case:  "Uint16Type",
				Value: uint16(233),
				Want:  "233",
			},
			{
				Case:  "Uint32Type",
				Value: uint32(233),
				Want:  "233",
			},
			{
				Case:  "Uint64Type",
				Value: uint64(233),
				Want:  "233",
			},
			{
				Case:  "Float32Type",
				Value: float32(123.456),
				Want:  "123.456",
			},
			{
				Case:  "Float64Type",
				Value: 123.456,
				Want:  "123.456",
			},
			{
				Case:  "BoolType",
				Value: true,
				Want:  "true",
			},
			{
				Case:  "OtherType",
				Value: nil,
				Want:  "",
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.Case, func(t *testing.T) {
				if NumberOrStringValueGetString(testCase.Value) != testCase.Want {
					t.Errorf("NumberOrStringValueGetString() = %v, want %v", NumberOrStringValueGetString(testCase.Value), testCase.Want)
				}
			})
		}
	})

	t.Run("NumberOrStringValueGetInt", func(t *testing.T) {
		testCases := []struct {
			Case  string
			Value any
			Want  int
		}{
			{
				Case:  "StringType",
				Value: "string",
				Want:  0,
			},
			{
				Case:  "IntType",
				Value: 123,
				Want:  123,
			},
			{
				Case:  "Int8Type",
				Value: int8(123),
				Want:  123,
			},
			{
				Case:  "Int16Type",
				Value: int16(123),
				Want:  123,
			},
			{
				Case:  "Int32Type",
				Value: int32(123),
				Want:  123,
			},
			{
				Case:  "Int64Type",
				Value: int64(123),
				Want:  123,
			},
			{
				Case:  "UintType",
				Value: uint(233),
				Want:  233,
			},
			{
				Case:  "Uint8Type",
				Value: uint8(233),
				Want:  233,
			},
			{
				Case:  "Uint16Type",
				Value: uint16(233),
				Want:  233,
			},
			{
				Case:  "Uint32Type",
				Value: uint32(233),
				Want:  233,
			},
			{
				Case:  "Uint64Type",
				Value: uint64(233),
				Want:  233,
			},
			{
				Case:  "Float32Type",
				Value: float32(123.456),
				Want:  123,
			},
			{
				Case:  "Float64Type",
				Value: 123.456,
				Want:  123,
			},
			{
				Case:  "BoolType",
				Value: true,
				Want:  1,
			},
			{
				Case:  "BoolType2",
				Value: false,
				Want:  0,
			},
			{
				Case:  "OtherType",
				Value: nil,
				Want:  0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.Case, func(t *testing.T) {
				if NumberOrStringValueGetInt(testCase.Value) != testCase.Want {
					t.Errorf("NumberOrStringValueGetInt() = %v, want %v", NumberOrStringValueGetInt(testCase.Value), testCase.Value)
				}
			})
		}
	})

	t.Run("NumberOrStringValueGetFloat", func(t *testing.T) {
		testCases := []struct {
			Case  string
			Value any
			Want  float64
		}{
			{
				Case:  "StringType",
				Value: "string",
				Want:  0,
			},
			{
				Case:  "IntType",
				Value: 123,
				Want:  123,
			},
			{
				Case:  "Int8Type",
				Value: int8(123),
				Want:  123,
			},
			{
				Case:  "Int16Type",
				Value: int16(123),
				Want:  123,
			},
			{
				Case:  "Int32Type",
				Value: int32(123),
				Want:  123,
			},
			{
				Case:  "Int64Type",
				Value: int64(123),
				Want:  123,
			},
			{
				Case:  "UintType",
				Value: uint(233),
				Want:  233,
			},
			{
				Case:  "Uint8Type",
				Value: uint8(233),
				Want:  233,
			},
			{
				Case:  "Uint16Type",
				Value: uint16(233),
				Want:  233,
			},
			{
				Case:  "Uint32Type",
				Value: uint32(233),
				Want:  233,
			},
			{
				Case:  "Uint64Type",
				Value: uint64(233),
				Want:  233,
			},
			{
				Case:  "Float32Type",
				Value: float32(123),
				Want:  123,
			},
			{
				Case:  "Float64Type",
				Value: 123.456,
				Want:  123.456,
			},
			{
				Case:  "BoolType",
				Value: true,
				Want:  1,
			},
			{
				Case:  "BoolType2",
				Value: false,
				Want:  0,
			},
			{
				Case:  "OtherType",
				Value: nil,
				Want:  0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.Case, func(t *testing.T) {
				if NumberOrStringValueGetFloat(testCase.Value) != testCase.Want {
					t.Errorf("NumberOrStringValueGetFloat() = %v, want %v", NumberOrStringValueGetFloat(testCase.Value), testCase.Want)
				}
			})
		}
	})
}

func TestDefault(t *testing.T) {
	t.Run("NilValue", func(t *testing.T) {
		Nil[int]()
	})

	t.Run("NilFunction", func(t *testing.T) {
		NilFunction[int]()()
	})
}

func TestArray(t *testing.T) {
	t.Run("Reverse", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			arr := []int{1, 2, 3, 4, 5}
			rev := []int{5, 4, 3, 2, 1}
			for i, v := range ReverseArray(arr) {
				if rev[i] != v {
					t.Errorf("ReverseArray[%d] = %v, want %v", i, v, rev[i])
				}
			}
		})

		t.Run("Empty", func(t *testing.T) {
			arr := []int{}
			rev := ReverseArray(arr)
			if rev == nil || len(rev) != 0 {
				t.Errorf("ReverseArray = %v, want %v", rev, arr)
			}
		})
	})
}
