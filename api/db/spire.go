package db

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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

type SpireSampleGenomes struct {
	Mags map[string]SpireGenome
}

type FloatOrString float64

func (f *FloatOrString) UnmarshalJSON(data []byte) error {
	// treat empty string as null
	if string(data) == `""` {
		*f = FloatOrString(0)
		return nil
	}

	// let json do the float parsing normally
	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		*f = FloatOrString(num)
		return nil
	}

	// try parsing a stringified number
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if s == "" {
			*f = 0
			return nil
		}
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*f = FloatOrString(v)
		return nil
	}

	return fmt.Errorf("invalid lat/lon: %s", data)
}

type SpireSampleMetadata struct {
	Latitute     FloatOrString  `json:"lat"`
	Longitude    FloatOrString  `json:"lon"`
	Mags         bool     `json:"mags"`
	Microntology []string `json:"microntology"`
}

type SpireStudy struct {
	Samples map[string]SpireSampleMetadata
}
