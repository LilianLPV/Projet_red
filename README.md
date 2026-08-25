# Projet_red

## Description 

Un jeu de rôle jouable dans le terminal, développé en Go.
 
Incarnez un aventurier, choisissez votre classe, explorez, combattez des ennemis.

## Prérequis

### Golang : 1.22 ou supérieur (https://go.dev/doc/install)

## Installation 

Clonez le dépôt :

```bash
git clone https://github.com/LilianLPV/Projet_red.git
cd Projet_red
```

Installez les dépendances :
 
```bash
go mod tidy
```

## Lancement
 
Depuis la racine du projet :
 
```bash
go run main.go
```

## Structure du projet

Le code est organisé **par fonctionnalité** plutôt que par couche technique. Chaque
fonctionnalité a son propre dossier, et à l'intérieur de chaque dossier, les fichiers
sont séparés par responsabilité :

- `model.go` — les structures de données et les définitions statiques (aucune logique)
- `service.go` — la logique métier : calculs, règles du jeu, mutation d'état
- `display.go` — tout ce qui touche au terminal : affichage (ASCII art, menus,
  messages) et lecture des choix du joueur

```
Projet_red/
├── main.go               # Point d'entrée : crée le personnage, lance l'intro et le menu
├── feature/               # Toutes les fonctionnalités du jeu, une par dossier
│   ├── character/          # Le personnage : Player, Stats, Equipment, création, fiche, XP
│   ├── inventory/           # Le sac à dos : Item, ajout/retrait/quantité d'objets
│   ├── combat/               # Le combat au tour par tour : Monster, Attack, déroulé des tours
│   ├── merchant/               # La boutique de Barnabé : achats, attaque spéciale, sac amélioré
│   ├── blacksmith/              # La forge de Thorin : Recipe, craft d'équipement
│   └── story/                    # Le scénario : intro, choix de départ, quête principale
├── menu/                   # Le menu principal : orchestre les fonctionnalités ci-dessus
└── utils/                   # Fonctions transverses : couleurs, affichage lent, entrées clavier
```

`menu/` et `utils/` ne sont pas des fonctionnalités mais du code d'orchestration /
transverse (comme c'était déjà le cas), donc ils ne sont pas subdivisés en
model/service/display.

### Dépendances entre dossiers (pas de cycle)

```
utils
  ↑
feature/inventory
  ↑
feature/character
  ↑        ↖
feature/combat   feature/blacksmith   feature/merchant   feature/story
  ↑                    ↑                    ↑                 ↑
              menu (orchestre tout, rien ne dépend de menu)
```

`inventory` ne connaît rien de `character` : il ne manipule que des listes d'objets.
C'est `character` qui expose des fonctions comme `AddItem`/`RemoveItem` sur le
`Player`, en s'appuyant sur `inventory` en interne. Cela évite les imports circulaires
tout en gardant chaque dossier centré sur une seule responsabilité.

## Comment jouer

Au lancement du jeu, vous créez votre personnage, puis vous accédez au menu principal, où vous pourrez commencer votre partie.

Pour jouer, vous naviguez dans les menus en saisissant le numéro correspondant à l'option souhaitée (le nombre d'options varie selon le menu).