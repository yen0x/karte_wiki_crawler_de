package model

type Karte struct {
	CardNameGer string `json:"cardNameGer"`
	CardNameJa  string `json:"cardNameJa"`
	CardNameEn  string `json:"cardNameEn"`

	MonsterAttr string `json:"monsterAttribute"`
	MonsterRank string `json:"monsterRank"`
}
