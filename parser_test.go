package flagcmd

import "testing"

func TestRegisterSubCommand(t *testing.T) {
	type args struct {
		sc *SubCommand
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterSubCommand(tt.args.sc); (err != nil) != tt.wantErr {
				t.Errorf("RegisterSubCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
