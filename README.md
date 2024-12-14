Jeu du Pendu (Hangman)

Ce projet est une implémentation du célèbre jeu du Pendu en Go avec une interface web. Les joueurs doivent deviner un mot en proposant des lettres une par une, tout en évitant de dépasser le nombre maximum de tentatives. ( Qui est de 10)

Fonctionnalités

- Interface utilisateur simple et responsive.
- Plusieurs niveaux de difficulté : facile, moyen et difficile.
- Gestion des tentatives et des erreurs.
- Affichage dynamique des lettres trouvées et des cœurs restants.
- Fin de partie avec un message personnalisé en cas de victoire ou de défaite.

Structure du projet : 
Hangmann Web : Game : go.mod - main.go - words.txt - words2.txt - words3.txt
               Template : gameover.html - index.html - hangman.html
                          Css : style.css
                          Image : 0.png - 1.png - 2.png - 3.png - 4.png - 5.png - 6.png - 7.png - 8.png - 9.png - 10.png

- `main.go` : Code source principal du serveur Go.
- `template/` : Contient les fichiers HTML et CSS.
- `index.html` : Page d'accueil.
- `hangman.html` : Page principale du jeu.
- `gameover.html` : Page de fin (victoire/défaite).
- `css/` : Styles pour le projet.
- `words.txt` : Fichier contenant les mots pour le niveau facile.
- `words2.txt` : Fichier pour le niveau moyen.
- `words3.txt` : Fichier pour le niveau difficile.




                          
Prérequis

- Go installé sur votre ordinateur
- Un navigateur web pour accéder au jeu.

Exécution

Je run mon main.go dans le terminal sur VS Code et je vais sur mon navigateur sur localhost:8080.
On peut aussi l'ouvrir avec LiveServer ( mais il y a trop de bug alors je le fait manuellement ) 

J'ai utilisé les langages Golang ; HTML ; CSS. 

Précision : Je me suis inspiré de la DA (= Direction Artistique du jeu Undertale )
