package csv

import (
	"reflect"
	"testing"

	"github.com/MustafaNafizDurukan/HashWalker/pkg/constants"
)

func TestCompare(t *testing.T) {
	type args struct {
		baseline   constants.MapPathHash
		fileHashes constants.MapPathHash
	}
	tests := []struct {
		name string
		args args
		want constants.MapEntryPathHash
	}{
		{
			name: "Deleted",
			args: args{
				baseline: constants.MapPathHash{
					"path1": "hashvalue",
					"path2": "hashvalue",
				},
				fileHashes: constants.MapPathHash{},
			},
			want: constants.MapEntryPathHash{
				constants.Deleted: constants.MapPathHash{
					"path1": "hashvalue",
					"path2": "hashvalue",
				},
			},
		},
		{
			name: "Added",
			args: args{
				baseline: constants.MapPathHash{},
				fileHashes: constants.MapPathHash{
					"path1": "hashvalue",
					"path2": "hashvalue",
				},
			},
			want: constants.MapEntryPathHash{
				constants.Added: constants.MapPathHash{
					"path1": "hashvalue",
					"path2": "hashvalue",
				},
			},
		},
		{
			name: "Modified",
			args: args{
				baseline: constants.MapPathHash{
					"path1": "hashvalue",
					"path2": "hashvalue",
				},
				fileHashes: constants.MapPathHash{
					"path1": "hashvalue_new",
					"path2": "hashvalue_new",
				},
			},
			want: constants.MapEntryPathHash{
				constants.ModifiedNew: constants.MapPathHash{
					"path1": "hashvalue_new",
					"path2": "hashvalue_new",
				},
				constants.ModifiedOld: constants.MapPathHash{
					"path1": "hashvalue",
					"path2": "hashvalue",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compare(tt.args.baseline, tt.args.fileHashes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
