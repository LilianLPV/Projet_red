# Projet_red

## Description 

Un jeu de rôle jouable dans le terminal, développé en Go.
 
Incarnez un aventurier, choisissez votre classe, explorez, combattez des ennemis.

## Prérequis

### Golang : 1.25.5 ou supérieur (https://go.dev/doc/install)

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
 
```
Projet_red/
├── main.go            # Point d'entrée du jeu
├── game/               # Moteur du jeu : menu, combat, création de personnage, scénario
├── feature/             # Systèmes annexes : marchand, forgeron, potions, inventaire
├── models/              # Structures de données : joueur, monstres, objets, recettes, statistiques
└── utils/               # Fonctions utilitaires : affichage, couleurs, lecture des entrées
```

## Comment jouer

Au lancement du jeu, vous créez votre personnage, puis vous accédez au menu principal, où vous pourrez commencer votre partie.

Pour jouer, vous naviguez dans les menus en saisissant le numéro correspondant à l'option souhaitée (le nombre d'options varie selon le menu).