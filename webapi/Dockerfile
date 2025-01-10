#Cette ligne spécifie l'image de base à utiliser. Ici, on utilise l'image officielle de Go avec la version 1.21.
FROM golang:1.21

#Définit le répertoire de travail dans le conteneur. À partir de ce #moment, toutes les commandes suivantes s'exécuteront dans ce répertoire.
WORKDIR /app

#Cette ligne copie tous les fichiers du répertoire ./webapi de votre machine locale vers le répertoire de travail (/app) dans l'image Docker. Cela inclut le fichier main.go et le fichier go.mod.
COPY ./webapi .

#Exécute la commande go build pour compiler le code Go dans le fichier exécutable main dans le répertoire actuel. Cela génère le binaire que vous allez exécuter dans le conteneur.
RUN go build -o main .

#Indique que l'application à l'intérieur du conteneur écoutera sur le port 8080.
EXPOSE 8080

#La commande à exécuter lorsque le conteneur démarre. Ici, elle lance l'application Go en exécutant le binaire main compilé précédemment.
CMD ["./main"]

