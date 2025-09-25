package db

type SpireGenome struct {
	Completeness  float64 `json:"completeness"`
	Contamination float64 `json:"contamination"`
	GeneCount     int     `json:"gene_count"`
	GenomeSize    int     `json:"genome_size"`
	GuncCssScore  float64 `json:"gunc_css"`
	GuncRrsScore  float64 `json:"gunc_rrs"`
	N50           int     `json:"n50"`
	NumContings   int     `json:"num_contings"`
	SampleId      string  `json:"sample_id"`
	SpireCluster  string  `json:"spire_cluster"`
}

type SpireSample struct {
	Lat          float64  `json:"lat"`
	Long         float64  `json:"long"`
	Mags         bool     `json:"mags"`
	Microntology []string `json:"microntology"`
	MagsList     []SpireGenome
}

type SpireStudy struct {
	Samples []SpireSample
}
