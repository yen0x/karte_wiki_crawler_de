package model

type Karte struct {
	Code string `json:"cardCode"`

	Categories []string `json:"cardCategories"`

	NameGer string `json:"cardNameGer"`
	NameJa  string `json:"cardNameJa"`
	NameEn  string `json:"cardNameEn"`

	Description string `json:"cardDescription"`
	EffectType  string `json:"effectType"`

	MonsterAttr    string `json:"monsterAttribute"`
	MonsterRank    string `json:"monsterRank"`
	MonsterAttack  string `json:"monsterAttack"`
	MonsterDefense string `json:"monsterDefense"`

	PictureUrl string `json:"pictureUrl"`
	WikiUrl    string `json:"wikiUrl"`
}
