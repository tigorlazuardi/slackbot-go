package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true for actual json",
			args: args{
				b: []byte(`{"foo":"bar"}`),
			},
			want: true,
		},
		{
			name: "returns false for not json",
			args: args{
				b: []byte(`asdf`),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsJSON(tt.args.b))
		})
	}
}

func TestIsJSONString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true for actual json",
			args: args{
				s: `{"foo":"bar","baz":1234}`,
			},
			want: true,
		},
		{
			name: "test json string",
			args: args{
				s: "[{\"foo\":\"bar\"}]",
			},
			want: true,
		},
		{
			name: "returns false for not json",
			args: args{
				s: "asdf",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				require.Equal(t, tt.want, IsJSONString(tt.args.s))
			})
		})
	}
}

// func TestSafeRawJSON_MarshalJSON(t *testing.T) {
// 	type T struct {
// 		Value SafeRawJSON
// 	}
// 	tests := []struct {
// 		name  string
// 		value T
// 		want  []byte
// 	}{
// 		{
// 			name:  "test number",
// 			value: T{Value: SafeRawJSON(`1234`)},
// 			want:  []byte(`{"Value":1234}`),
// 		},
// 		{
// 			name:  "test float",
// 			value: T{Value: SafeRawJSON(`1234.678`)},
// 			want:  []byte(`{"Value":1234.678}`),
// 		},
// 		{
// 			name:  "test negative float",
// 			value: T{Value: SafeRawJSON(`-1234.678`)},
// 			want:  []byte(`{"Value":-1234.678}`),
// 		},
// 		{
// 			name:  "test single digit 0",
// 			value: T{Value: SafeRawJSON(`0`)},
// 			want:  []byte(`{"Value":0}`),
// 		},
// 		{
// 			name:  "test real number with decimal digit that starts from 0.",
// 			value: T{Value: SafeRawJSON(`0.123`)},
// 			want:  []byte(`{"Value":0.123}`),
// 		},
// 		{
// 			name:  "test single digit 1",
// 			value: T{Value: SafeRawJSON(`1`)},
// 			want:  []byte(`{"Value":1}`),
// 		},
// 		{
// 			name:  "test broken json",
// 			value: T{Value: SafeRawJSON(`{"Value":"aaa}`)},
// 			want:  []byte(`{"Value":1}`),
// 		},
// 		{
// 			name:  "test string",
// 			value: T{Value: SafeRawJSON(`abra kadabra`)},
// 			want:  []byte(`{"Value":"abra kadabra"}`),
// 		},
// 		{
// 			name:  "test not valid json number",
// 			value: T{Value: SafeRawJSON(`012345`)},
// 			want:  []byte(`{"Value":"012345"}`),
// 		},
// 		{
// 			name:  "test not valid json number float test",
// 			value: T{Value: SafeRawJSON(`01.2345`)},
// 			want:  []byte(`{"Value":"01.2345"}`),
// 		},
// 		{
// 			name:  "test not valid json number 2",
// 			value: T{Value: SafeRawJSON(`123 45`)},
// 			want:  []byte(`{"Value":"123 45"}`),
// 		},
// 		{
// 			name:  "test json object",
// 			value: T{Value: SafeRawJSON(`{"foo":"bar"}`)},
// 			want:  []byte(`{"Value":{"foo":"bar"}}`),
// 		},
// 		{
// 			name:  "test underlying json object not affected",
// 			value: T{Value: SafeRawJSON(`{"foo":"1234"}`)},
// 			want:  []byte(`{"Value":{"foo":"1234"}}`),
// 		},
// 		{
// 			name:  "test json array",
// 			value: T{Value: SafeRawJSON(`["1","2"]`)},
// 			want:  []byte(`{"Value":["1","2"]}`),
// 		},
// 		{
// 			name:  "test null",
// 			value: T{Value: nil},
// 			want:  []byte(`{"Value":null}`),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := json.Marshal(tt.value)
// 			require.Nil(t, err)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }
