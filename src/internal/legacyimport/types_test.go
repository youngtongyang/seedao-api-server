package legacyimport

import (
	"testing"

	"github.com/seedao/seedao-api-server/internal/config"
)

func TestResolveSBTNames(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		sbt      apiSBT
		want     string
		wantDesc string
	}{
		{
			name:     "metadata name",
			sbt:      apiSBT{Metadata: apiSBTMetadata{Name: "A", Headline: "B", Description: "C"}},
			want:     "A",
			wantDesc: "C",
		},
		{
			name:     "fallback headline",
			sbt:      apiSBT{Metadata: apiSBTMetadata{Headline: "B", Description: "C"}},
			want:     "B",
			wantDesc: "C",
		},
		{
			name:     "fallback description",
			sbt:      apiSBT{Metadata: apiSBTMetadata{Description: "C"}},
			want:     "C",
			wantDesc: "C",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, desc := resolveSBTNames(tc.sbt)
			if got != tc.want || desc != tc.wantDesc {
				t.Fatalf("got (%q,%q) want (%q,%q)", got, desc, tc.want, tc.wantDesc)
			}
		})
	}
}

func TestNormalizeImageURL(t *testing.T) {
	t.Parallel()
	im := &Importer{cfg: config.V1ImportConfig{IPFSGateway: "https://dweb.link/ipfs/"}}
	cases := map[string]string{
		"ipfs://bafyxxx/1.png":      "https://dweb.link/ipfs/bafyxxx/1.png",
		"https://cdn.example/a.png": "https://cdn.example/a.png",
		"QmUxRHHVeKT73qvBEtjwBa7PYaKz2D4PV91DDTC2d8r13G": "https://dweb.link/ipfs/QmUxRHHVeKT73qvBEtjwBa7PYaKz2D4PV91DDTC2d8r13G",
		"https://ipfs.io/ipfs/QmYbJWpwwzCM6U98pjpz32MSppFRg3tvzNfuJWZkE9L29G": "https://dweb.link/ipfs/QmYbJWpwwzCM6U98pjpz32MSppFRg3tvzNfuJWZkE9L29G",
		"": "",
	}
	for in, want := range cases {
		if got := im.normalizeImageURL(in); got != want {
			t.Fatalf("normalizeImageURL(%q)=%q want %q", in, got, want)
		}
	}
}
