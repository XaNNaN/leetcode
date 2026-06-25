package prefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"fl", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"empty result", args{[]string{"dog", "racecar", "car"}}, ""},
		{"empty input", args{[]string{""}}, ""},
		{"single char", args{[]string{"a"}}, "a"},
		{"dog", args{[]string{"dog", "racecar", "car"}}, ""},
		{"c", args{[]string{"cir", "car"}}, "c"},
		{"ab", args{[]string{"a", "b"}}, ""},
		{"flowers", args{[]string{"flower", "flower"}}, "flower"},
		{"flopers", args{[]string{"flower", "flxper"}}, "fl"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %q, want %q", got, tt.want)
			}
		})
	}
}
