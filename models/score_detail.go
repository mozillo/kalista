package models

// belongs_to Game
// belongs_to Tenant
// belongs_to Scoreable (polymorphic) Gorm 不支持

// ScoreDetail 久远来看应该是一个多态关联, ScoreableType ScoreableID
type ScoreDetail struct {
	ID            uint   `gorm:"primary_key"`
	ScoreableType string `gorm:"not null"`
	ScoreableID   uint   `gorm:"not null"`
	Date          string `gorm:"type:varchar(10);unique_index:idx_for_score_details"`
	Game          Game
	GameID        uint `gorm:"unique_index:idx_for_score_details;index:idx_game_details"`
	Tenant        Tenant
	TenantID      uint `gorm:"unique_index:idx_for_score_details;index:idx_tenant_details"`
}

func NewScoreDetail(date string, game Game, tenant Tenant) {
	scoreDetail := ScoreDetail{
		Date:   date,
		Game:   game,
		Tenant: tenant,
	}
	db.Create(scoreDetail)
}

func (ScoreDetail) TableName() string {
	return "score_details"
}

//User Story
