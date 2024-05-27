package controllers

import (
	"html/template" // Importe le package html/template pour traiter les modèles HTML
	"net/http"      // Importe le package net/http pour gérer les requêtes et réponses HTTP
)

// Gestionnaire de la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "accueil", nil) // Appelle la fonction renderTemplate pour rendre le modèle "accueil"
}

// Gestionnaire de la page "services"
func ServicesHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "services", nil) // Appelle la fonction renderTemplate pour rendre le modèle "services"
}

// Gestionnaire de la page "projets"
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects", nil) // Appelle la fonction renderTemplate pour rendre le modèle "projects"
}

// Gestionnaire de la page "contact"
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact", nil) // Appelle la fonction renderTemplate pour rendre le modèle "contact"
}

// Fonction utilitaire pour rendre les modèles HTML
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Parse les fichiers de modèles HTML spécifiés
	tpl, err := template.ParseFiles("views/base.html", "views/header.html", "views/footer.html", "views/"+tmpl+".html")
	if err != nil {
		// En cas d'erreur, renvoie une erreur HTTP 500 avec le message d'erreur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Exécute le modèle "base" avec les données fournies
	tpl.ExecuteTemplate(w, "base", data)
}
