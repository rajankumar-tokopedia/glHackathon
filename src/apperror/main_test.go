package apperror

import (
	"reflect"
	"testing"
)

func TestBadError(t *testing.T) {
	type args struct {
		code    string
		message string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "BadError",
			args: args{
				"BD_ERR",
				"Something went wrong",
			},
			want: &Error{
				"BD_ERR",
				"Something went wrong",
				400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BadError(tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BadError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBadGatewayError(t *testing.T) {
	type args struct {
		code    string
		message string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "BadGatewayError",
			args: args{
				"BD_ERR",
				"Something went wrong",
			},
			want: &Error{
				"BD_ERR",
				"Something went wrong",
				502,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BadGatewayError(tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BadGatewayError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Code(t *testing.T) {
	type fields struct {
		code     string
		message  string
		httpCode int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "GetCode",
			fields: fields{
				httpCode: 403,
				code:     "BD_ERR",
			},
			want: "BD_ERR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{
				code:     tt.fields.code,
				message:  tt.fields.message,
				httpCode: tt.fields.httpCode,
			}
			if got := err.Code(); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		code     string
		message  string
		httpCode int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Error",
			fields: fields{
				message: "something went wrong",
			},
			want: "something went wrong",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_HttpCode(t *testing.T) {
	type fields struct {
		code     string
		message  string
		httpCode int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "GET HTTP CODE",
			fields: fields{
				httpCode: 403,
			},
			want: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{
				code:     tt.fields.code,
				message:  tt.fields.message,
				httpCode: tt.fields.httpCode,
			}
			if got := err.HttpCode(); got != tt.want {
				t.Errorf("HttpCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForbiddenError(t *testing.T) {
	type args struct {
		code    string
		message string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "ForbiddenError",
			args: args{
				"BD_ERR",
				"Something went wrong",
			},
			want: &Error{
				"BD_ERR",
				"Something went wrong",
				403,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ForbiddenError(tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForbiddenError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternalServerError(t *testing.T) {
	type args struct {
		code    string
		message string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "InternalServerError",
			args: args{
				"BD_ERR",
				"Something went wrong",
			},
			want: &Error{
				"BD_ERR",
				"Something went wrong",
				500,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InternalServerError(tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InternalServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotFoundError(t *testing.T) {
	type args struct {
		code    string
		message string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "NotFoundError",
			args: args{
				"BD_ERR",
				"Something went wrong",
			},
			want: &Error{
				"BD_ERR",
				"Something went wrong",
				404,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotFoundError(tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotImplementedError(t *testing.T) {
	type args struct {
		code    string
		message string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "NotFoundError",
			args: args{
				"BD_ERR",
				"Something went wrong",
			},
			want: &Error{
				"BD_ERR",
				"Something went wrong",
				501,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotImplementedError(tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotImplementedError() = %v, want %v", got, tt.want)
			}
		})
	}
}
