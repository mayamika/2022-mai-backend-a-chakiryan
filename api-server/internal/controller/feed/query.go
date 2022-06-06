package feed

func buildSearchFeedQuery(from []int, search *string) map[string]any {
	var should []map[string]any
	for _, f := range from {
		should = append(should, map[string]any{
			"match": map[string]any{
				"from": f,
			},
		})
	}
	must := []map[string]any{
		{
			"bool": map[string]any{
				"should": should,
			},
		},
	}
	if search != nil {
		must = append(must, map[string]any{
			"match": map[string]any{
				"text": map[string]any{
					"query":                *search,
					"fuzziness":            "AUTO",
					"fuzzy_transpositions": true,
					"operator":             "or",
					"minimum_should_match": 1,
					"analyzer":             "standard",
					"zero_terms_query":     "none",
					"lenient":              false,
					"cutoff_frequency":     0.01,
					"prefix_length":        0,
					"max_expansions":       50,
					"boost":                1,
				},
			},
		})
	}
	return map[string]any{
		"query": map[string]any{
			"bool": map[string]any{
				"must": must,
			},
		},
	}
}
