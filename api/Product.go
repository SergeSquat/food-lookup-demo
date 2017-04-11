package api

type Product struct {
	Carbohydrate_g float64 `json:"carbohydrate_g"`
	Protein_g      float64 `json:"protein_g"`
	Fa_sat_g       float64 `json:"fa_sat_g"`
	Fa_mono_g      float64 `json:"fa_mono_g"`
	Fa_poly_g      float64 `json:"fa_poly_g"`
	Kcal           float64 `json:"kcal"`
	Description    string  `json:"description"`
}

