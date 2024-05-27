package main

import (
	"fmt"             // Importe le package fmt pour formater et afficher des messages dans la console
	"net/http"        // Importe le package net/http pour créer un serveur HTTP
	"tp1/controllers" // Importe le package controllers qui contient les gestionnaires de requêtes

	"github.com/gorilla/mux" // Importe le package mux de Gorilla pour créer un routeur HTTP flexible
)

func main() {
	// Initialisation du routeur Gorilla Mux
	router := mux.NewRouter() // Crée une nouvelle instance de routeur Gorilla Mux

	// Définition des routes
	router.HandleFunc("/", controllers.HomeHandler)             // Associe le chemin racine "/" au gestionnaire HomeHandler
	router.HandleFunc("/services", controllers.ServicesHandler) // Associe le chemin "/services" au gestionnaire ServicesHandler
	router.HandleFunc("/projects", controllers.ProjectsHandler) // Associe le chemin "/projects" au gestionnaire ProjectsHandler
	router.HandleFunc("/contact", controllers.ContactHandler)   // Associe le chemin "/contact" au gestionnaire ContactHandler
	// Dossier statique pour les fichiers CSS et images
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// Définit un gestionnaire pour les fichiers statiques (CSS, images, etc.) dans le répertoire "./static/"
	// Utilise StripPrefix pour enlever le préfixe "/static/" des chemins de fichier

	// Démarrage du serveur web
	fmt.Println("http://localhost:8080") // Affiche l'URL du serveur dans la console
	http.ListenAndServe(":8080", router) // Démarre le serveur web sur le port 8080 en utilisant le routeur défini
}
