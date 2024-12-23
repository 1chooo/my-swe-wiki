package main

var Data []Role

func init() {
	Data = []Role{
		Role{
			ID:      1,
			Name:    "阿修羅",
			Summary: "死國魖族最強者，在歷經多場戰役統一死國，直接挑戰天者權威，不服其領導，要求讓三族和平共處，因此被天者以增加資源為名，前往打通連接火宅佛獄、死國及苦境的莫汗走廊，未料在工程半途被楓岫主人所乘駕的隕石撞毀，阿修羅在天者刻意算計下意外身亡，身葬苦境與死國間的異次元。",
			Skills: []RoleSkill{
				RoleSkill{ID: 1, Type: MartialArts, Name: "天之爆"},
				RoleSkill{ID: 2, Type: MartialArts, Name: "魔之狂"},
				RoleSkill{ID: 3, Type: MartialArts, Name: "天之渦"},
				RoleSkill{ID: 4, Type: MartialArts, Name: "闇之爆"},
				RoleSkill{ID: 5, Type: Magic, Name: "山河凝元·天地共引"},
				RoleSkill{ID: 6, Type: Magic, Name: "地之火·九天滅絕"},
			},
		},
		Role{
			ID:      2,
			Name:    "白塵子",
			Summary: "火宅佛獄凱旋侯的副體之一，本名黑枒君，臥底中原武林，向佛獄通風報信，最後被素還真所殺並冒充其身份一探佛獄之秘。",
			Skills: []RoleSkill{
				RoleSkill{ID: 7, Type: MartialArts, Name: "凝宇化空"},
				RoleSkill{ID: 8, Type: MartialArts, Name: "反神源"},
			},
		},
	}
}

// 布袋戲角色
type Role struct {
	ID      uint        `json:"id"`      // Key
	Name    string      `json:"name"`    // 角色名稱
	Summary string      `json:"summary"` // 介紹
	Skills  []RoleSkill `json:"skills"`  // 會使用的招式
}

// RoleSkill 用來記錄招式名稱
type RoleSkill struct {
	ID   uint      `json:"id"`   //key
	Type SkillType `json:"type"` // 招式分類
	Name string    `json:"name"` // 招式名稱
}

type SkillType string

const (
	MartialArts SkillType = "武學"
	Magic       SkillType = "法術"
)
