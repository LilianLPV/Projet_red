package combat

import (
	"fmt"

	"Projet-Red/feature/character"
	"Projet-Red/utils"
)

func DisplayCombatIntro() {
	fmt.Println(utils.Yellow + `
                     .::::::  .                                        ..:::::::.                   
                     ^  ......^                                       .^        ^                   
                   .^ .^.^:.::^                                       .^. :: :: :.                  
                   .^~.  ^^.:^^                                       .~^.^:  .~^.                  
                    ^^.    ^. ~                 :::.                  .^.^    .^^                   
                     ^:: .::::                 ~~~~~.                  :::.  :: ^.                  
                   ..^ . ..^^.                 :^. ^^^                ..:^^. .  :::..               
               .. ...::.::::..: .:              ^. : ^::::::::. .:::::..:.::::::^  ... .            
              ^.       ....       ^.             .:......     .:.  ~. ^~^:~:           ^.           
              ^    ..::~^:.^^^:    :.                ...:::::......^..:^.  ^.    .:    ^.           
              ^.    :^:^^::^::^. ..:^                        ......:...:::. ^.   :.    :.           
               ::.::.^.        :^:   ^.                                ^. ^. .::^:::..^             
                :^    .:::...  :~.   ::                                ::  .:       :^              
                 :..      ...::^~...::            ....                 ::   : .   . :.              
                  ^:^.......... ...^^            :^::^:                ::     .^.:: ::              
                  ^. .....:^:..^:..^~^      .....^.  .:                ::       .   ::              
                  ^.     :^.  .::.::::      .....^  ::::::::           ^            ::              
                  ^:::::::^~^^^.:^          .....::. .:               :. .:.:::::::.^:              
                 ^  .....        ::         .....^:. .:             .^      ........ ::             
                :.                .:             :  ::::::::      .:.                 ^             
               .:         .        .:        .::.^.  .:          .:.        .         ^.            
               ^         ^::        .^       ..::^.  .^.        ::        .:^.        .^            
              ^         ^.  ^         :.      .:.:.  .:..      ^.        ::  :.        :.           
 :::::::::::. ^        :.   ...       ....::. :. ...... : ::....       ...    :.       .^ .:::::::. 
 ............ ^       ::....:..^.       ^....:.....:...:.....^:       ::.......^.       ~ ......... 
         :: ^:       ^.    ^:   ::      ^   ^.     ^:  :^    ::     .^  ^.      ^.       ^..~.      
      .:   :.      :..   ...    ::      ^  .:      ^.   ^.   ::     .^   .:      .:.      :.  .:    
    . .:  .:      .:   . :    . ::      ^  ^.      ^.   .:   ::     .^.   ..      .:.     .:  .:.   
 ..^.    .^      ^.   .^     .^ ::      ^ ::       ^.    ^.  ::     .^^.    .^      ^.      ^   ::. 
 ..   .^::     ::    ^      ^.    ^     ^::        ^.     :: ::    :.  :.     .:      ^      ^.  .. 
    .. .^     .:  ....    ...     ^    .:.         ^.      . :.    ^.   ..     . .    .:     ::.    
   .:   ^::::::   :.     .:       ^..... ::.       ^.     ..:  ....^.    :.      ::    .^:::.^::    
 ::     ^.   ^..^       ^.        ^..........:^    ^.  .^......... ^.      ^.      ::   ^.   ^.     
 .      ^.    ^       :.          ....^^......     ^.   .....^:.....        :.      .  .:    ^.     
        .:^::::      .:              .:            ^.        .:              :.        .::::::      
         .^       ..^.              .^             ^:          ^              ^:.                   
	  ` + utils.Reset)
}

func PromptEngageCombat() {
	fmt.Println("\nAppuie sur Entrée pour engager le combat...")
	utils.ReadInput()
}

func DisplayTourHeader(tour int) {
	fmt.Println()
	fmt.Println(utils.Cyan+"=== Tour", tour, "==="+utils.Reset)
}

func DisplayFlee() {
	fmt.Println("Vous prenez la fuite !")
}

// DisplayPlayerStunnedSkip annonce que le joueur passe son tour, étourdi.
func DisplayPlayerStunnedSkip() {
	fmt.Println("Vous êtes étourdi et passez votre tour !")
}

// DisplayPlayerStuns annonce que le joueur étourdit l'ennemi.
func DisplayPlayerStuns(enemyName string) {
	fmt.Println("Vous étourdissez", enemyName, "!")
}

// DisplayEnemyPoisonTick affiche les dégâts de poison subis par l'ennemi.
func DisplayEnemyPoisonTick(name string, degats, hpRestant int) {
	fmt.Println(name, "subit", degats, "dégâts de poison ! PV :", hpRestant)
}

// DisplayEnemyStunnedSkip annonce que l'ennemi passe son tour, étourdi.
func DisplayEnemyStunnedSkip(name string) {
	fmt.Println(name, "est étourdi et passe son tour !")
}

// DisplayPlayerDodges annonce que le joueur esquive l'attaque de l'ennemi.
func DisplayPlayerDodges(name string) {
	fmt.Println("Vous esquivez l'attaque de", name, "!")
}

// DisplayEnemyCritical annonce une attaque critique de l'ennemi.
func DisplayEnemyCritical(name string) {
	fmt.Println(name, "déchaîne une attaque critique !")
}

// DisplayEnemyHit affiche les dégâts infligés au joueur par l'ennemi.
func DisplayEnemyHit(name string, degats, hpRestant int) {
	fmt.Println(name, "vous inflige", degats, "dégâts. PV:", hpRestant)
}

// DisplayEnemyStuns annonce que l'ennemi étourdit le joueur.
func DisplayEnemyStuns(name string) {
	fmt.Println(name, "vous étourdit !")
}

// DisplayEnemyDodges annonce que l'ennemi esquive l'attaque du joueur.
func DisplayEnemyDodges(name string) {
	fmt.Println(name, "esquive votre attaque !")
}

// DisplayPlayerBaseAttackHit affiche les dégâts de l'attaque de base du joueur.
func DisplayPlayerBaseAttackHit(degats int) {
	fmt.Println("Vous infligez", degats, "dégâts !")
}

// DisplaySkillNotBought prévient que le joueur n'a pas encore acheté d'attaque spéciale.
func DisplaySkillNotBought() {
	fmt.Println("Tu n'as pas d'attaque spéciale ! Achète-la chez le marchand.")
}

// DisplayNotEnoughMana prévient que le joueur n'a pas assez de mana pour son attaque spéciale.
func DisplayNotEnoughMana(actuel, requis int) {
	fmt.Println("Pas assez de mana ! (", actuel, "/", requis, ")")
}

// DisplaySkillHit affiche les dégâts infligés par l'attaque spéciale du joueur.
func DisplaySkillHit(nomSkill string, degats, coutMana int) {
	fmt.Println("Vous lancez", nomSkill, "et infligez", degats, "dégâts ! (-", coutMana, "mana)")
}

// DisplayCombatStatus affiche l'état courant du combat (PV, mana des deux combattants).
func DisplayCombatStatus(p *character.Player, e *Monster) {
	fmt.Println()
	fmt.Println("=== COMBAT ===")
	fmt.Println("Vous :", p.HP, "/", p.HPMax, "PV  |  Mana:", p.Mana, "/", p.ManaMax)
	fmt.Println(e.Name, ":", e.HP, "/", e.HPMax, "PV")
	fmt.Println(" ")
	fmt.Println("C'est à vous de combattre !")
	fmt.Println(" ")
}

// DisplayPlayerMenu affiche les actions possibles pour le tour du joueur.
func DisplayPlayerMenu(p *character.Player) {
	fmt.Println("\033[96m1.\033[0m Attaquer", p.BaseAttack)
	if len(p.Skill) > 0 {
		fmt.Println("\033[96m2.\033[0m Attaque spéciale :", p.Skill[0])
	} else {
		fmt.Println("\033[96m2.\033[0m Attaque spéciale : (non achetée)")
	}
	fmt.Println("\033[96m3.\033[0m Inventaire")
	fmt.Println("\033[96m4.\033[0m Fuir le combat")
	fmt.Print("Choisissez une option: ")
}

// ReadPlayerChoice lit le choix du joueur pour son tour de combat.
func ReadPlayerChoice() string {
	return utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4"})
}

// DisplayInventoryMenu affiche le sous-menu des potions utilisables en combat.
func DisplayInventoryMenu(p *character.Player) {
	fmt.Println("--- Inventaire ---")
	fmt.Println("1) Potion de santé (x", character.QuantityOf(p, "Potion de santé"), ")")
	fmt.Println("2) Potion de mana (x", character.QuantityOf(p, "Potion de mana"), ")")
	fmt.Println("3) Potion de poison (x", character.QuantityOf(p, "Potion de poison"), ")")
	fmt.Println("4) Retour")
}

// ReadInventoryChoice lit le choix du joueur dans le sous-menu de l'inventaire.
func ReadInventoryChoice() string {
	return utils.ReadChoice("Ton choix : ", []string{"1", "2", "3", "4"})
}

// DisplayNoHealthPotion prévient que le joueur n'a pas de potion de santé.
func DisplayNoHealthPotion() {
	fmt.Println("Tu n'as pas de potion de santé !")
}

// DisplayNoManaPotion prévient que le joueur n'a pas de potion de mana.
func DisplayNoManaPotion() {
	fmt.Println("Tu n'as pas de potion de mana !")
}

// DisplayNoPoisonPotion prévient que le joueur n'a pas de potion de poison.
func DisplayNoPoisonPotion() {
	fmt.Println("Vous n'avez pas de potion de poison dans votre inventaire")
}

// DisplayAlreadyPoisoned prévient que l'ennemi est déjà empoisonné.
func DisplayAlreadyPoisoned(name string) {
	fmt.Println(name, "est déjà empoisonné ! Impossible de cumuler.")
}

// DisplayPoisonThrown annonce que le joueur jette une potion de poison sur l'ennemi.
func DisplayPoisonThrown(name string) {
	fmt.Println("Vous jetez la potion sur", name, "!")
}

// DisplayVictory annonce la victoire du joueur et le butin récupéré.
func DisplayVictory(enemyName string, or int) {
	fmt.Println("Vous avez vaincu", enemyName, "et gagné", or, "PO !")
}

// DisplayDefeatRevive annonce que le joueur se relève une dernière fois après sa défaite.
func DisplayDefeatRevive(hp int) {
	fmt.Println("Vous êtes tombé... mais vous vous relevez une dernière fois avec", hp, "PV !")
}

// DisplayDefeatFinal annonce la mort définitive du joueur.
func DisplayDefeatFinal() {
	fmt.Println("Vous succombez définitivement...")
}
