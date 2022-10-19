package token

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		username  string
		role      string
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			args: args{
				username:  "khangpt3",
				role:      "admin",
				secretKey: "dummy",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "test_2",
			args: args{
				username:  "khangpt3",
				role:      "admin",
				secretKey: "dummy",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "test_3",
			args: args{
				username:  "khangpt3",
				role:      "admin",
				secretKey: "dummy",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.username, tt.args.role, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			userClaims, err := Verify(got)
			if err != nil {
				t.Errorf("Error = %v", err.Error())
			}
			if (userClaims.Username != tt.args.username) || (userClaims.Role != tt.args.role) {
				t.Errorf("Generate() username = %v, username %v, Generate() role = %v, role %v", userClaims.Username, tt.args.username, userClaims.Role, tt.args.role)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	type args struct {
		accessToken string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserClaims
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			args: args{
				accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjYyODcwNDQsIlVzZXJuYW1lIjoia2hhbmdwdDMiLCJSb2xlIjoiYWRtaW4ifQ.gsyPEgWaqlJNa1y8yLZFgjk387uS8ikBXUptOwFzhBA",
			},
			want: &UserClaims{},
		},
		{
			name: "test_2",
			args: args{
				accessToken: "",
			},
			want: &UserClaims{},
		},
		{
			name: "test_3",
			args: args{
				accessToken: "",
			},
			want: &UserClaims{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Verify(tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
