package handlers

type HeartContainer struct {
	userID           int
	currentFragments int
	maxFragments     int
	healingFactor    int //# of fragments you regenerate each cycle
}

func NewHeartContainer(id int, hearts int, hf int) *HeartContainer {
	frags := hearts * 4

	hc := HeartContainer{
		userID:           id,
		currentFragments: frags,
		maxFragments:     frags,
		healingFactor:    hf,
	}

	return &hc
}

func (hc HeartContainer) CanAfford(cost int) bool {
	return hc.currentFragments >= cost
}

func (hc *HeartContainer) Hurt(pain int) bool {
	if hc.CanAfford(pain) {
		hc.currentFragments = hc.currentFragments - pain

		return true
	}

	return false
}

func (hc HeartContainer) GetHealingFactor() int {
	return hc.healingFactor
}

func (hc HeartContainer) GetHeartID() int {
	return hc.userID
}

func (hc HeartContainer) CurrentLives() float64 {
	return float64(hc.currentFragments) / 4.0
}

func (hc HeartContainer) MaxLives() float64 {
	return float64(hc.maxFragments) / 4.0
}

func (hc HeartContainer) CurrentFragments() int {
	return hc.currentFragments
}

func (hc HeartContainer) MaxFragments() int {
	return hc.maxFragments
}

func (hc HeartContainer) IsFullyHealed() bool {
	return hc.currentFragments == hc.maxFragments
}

//returns true if hc becomes fully healed
func (hc *HeartContainer) Heal() bool {

	hc.currentFragments = hc.currentFragments + hc.healingFactor

	if hc.currentFragments >= hc.maxFragments {
		hc.currentFragments = hc.maxFragments
		return true
	}

	return false
}
