# TP de programmation
 Gauthier Abomnes & Yann Bretheau | EPSI I4 | 2017-2018

## Rappel du sujet
Le but de ce TP était de réaliser une application capable de sauvegarder et de consulter des notes. L'application ne comprend d'interface graphique.

## Descriptif du document
Ce document présente les étapes et les différents points nécessaire afin de faire fonctionner l'application. Il comporte également le descriptif de l'architecture de l'application.

## Prérequis
* Une installation de Go `apt-get install golang-go`
* Une installation de Redis `apt-get install redis-server`

## Installation
Dans un premier temps il faut définir le `$GOPATH` afin de pouvoir utiliser Go. Cela ce fait avec la commande suivante : `export GOPATH=$HOME/go`. Par la suite il faut créer un dossier `src` dans le `$GOPATH` afin de faciliter le développement.

Une fois ces étapes effectuées on peut commencer l'installation de l'application :
* Créer un dossier `tp1` dans le dossier `src`
* Placer le contenu de l'archive dans le dossier `tp1`
* Installer le driver Radix V2 `go get github.com/mediocregopher/radix.v2`
* Installer le wrapper HTTP `go get github.com/julienschmidt/httprouter`

## Utilisation
Pour lancer l'application il suffit de se placer dans le dossier où se trouve le `main.go` et exécuter la commande suivante : `go run main.go`. Il est également possible de le lancer depuis n'importe quel endroit via la commande : `go run $GOPATH/src/tp1/main.go`. L'application écoute désormais sur le port 12345. Il suffit d'ouvrir un autre terminal afin de pouvoir utiliser l'application.

### Opérations de l'application
L'application dispose des trois opérations de base à savoir l'ajout d'une note, la lecture d'une note via son ID et la lecture de toute les notes. De plus l'opération bonus de suppression a également été implémenté.
Ci-dessous la liste des commandes pour les opérations :
* Ajout d'une note : `curl -i -L -d "Titre de la note" http://localhost:12345/notes`
* Visualisation d'une note : `curl -i http://localhost:12345/notes/{id}`
* Visualisation de toute les notes : `curl -i http://localhost:12345/notes`
* Suppression d'une note : `curl -i http://localhost:12345/delete/{id}`

## Architecture
L'application est composée de deux fichiers. Le fichier `main.go` et le fichier `note.go`. Le découpage du projet en deux applications est très interéssant. Dans le fichier `note.go` se trouvent toute l'interaction avec la base Redis. Dans le fichier `main.go` se trouve les fonctions qui gèrent les requetes HTTP.

Fichier `note.go`

Dans ce fichier il y a 3 fonctions principales. Une pour ajouter une note, une autre pour chercher une note et une autre pour supprimer une note. D'autres fonctions utilitaires sont présentes comme la récupération du nombre total de note.

Fichier `main.go`
Dans ce fichier il y a 4 fonctions principales. Une pour ajouter une note, une pour afficher une note par rapport à son ID, une pour afficher la liste des notes et une autre pour supprimer une note.