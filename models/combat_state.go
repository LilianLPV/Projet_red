package models

// Structure des effets pendant les combats + le nbr de tour 

type CombatState struct {
	JoueurEtourdi    bool
	EnnemiEtourdi    bool
	EnnemiEmpoisonne int
	Tour             int
}
