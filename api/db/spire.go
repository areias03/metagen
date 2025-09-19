package db

type SpireGenome struct {
	id             string
	completeness   float64 `json:"completeness"`
	contamination  float64 `json:"contamination"`
	gene_count     int     `json:"gene_count"`
	genome_size    int     `json:"genome_size"`
	gunc_css_score float64 `json:"gunc_css"`
	gunc_rrs_score float64 `json:"gunc_rrs"`
	n50            int     `json:"n50"`
	num_contings   int     `json:"num_contings"`
	sample_id      string  `json:"sample_id"`
	spire_cluster  string  `json:"spire_cluster"`
}

type SpireSample struct {
	id           string
	lat          float64  `json:"lat"`
	long         float64  `json:"long"`
	mags         bool     `json:"mags"`
	microntology []string `json:"microntology"`
	mags_list    []SpireGenome
}

type SpireStudy struct {
	id      string
	samples []SpireSample
}

// TODO: function that initializes structs from config
