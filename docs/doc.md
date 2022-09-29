 # Documentation: Project-Red

*Par Marianne Corbel et Eva Chibane*
____

## A propos
Ce projet a été réalisé dans le cadre de **Ymmersion** à Ynov Aix par **Marianne Corbel** et **Eva Chibane**, respectant un cahier des charges spécifique.

___

*Ce projet a pour vocation d’être la démo d’un jeu-vidéo court de type RPG tour-par-tour. Les fonctionnalités disponibles dans la démo sont susceptibles d’évoluer avec le temps !*
___

## Packages & Dépendances
___Packages Internes:___

* ProjectRed/menu: Contient la majorité des fonctions créées pour le programme. Divisé en plusieurs fichiers dont les noms sont supposés être explicites et facilement identifiables.

* ProjectRed/utils: Contient d’autres fonctions personnalisées qui ne sont pas nécessairement liées au projet en lui-même (fonctions de formattage, aléatoire, Prints “customisés”...
Dépendances:
Box CLI Maker: Utilisé pour créer les entêtes rectangulaires colorées.
Aurora et ansi: Utilisation des couleurs dans le terminal.
___

## Programme
Project-Red comporte actuellement quatre grandes fonctionnalités disponibles pour le personnage d’un joueur depuis le menu principal: l’accès et l’utilisation de l’inventaire, le magasin (système d’achat), le forgeron (système de craft) et l’entraînement (combat tour par tour).
___

## Structures:

`Character`: Utilisé pour le joueur

`NPC`: Structure interne des NPC.

`ShopKeeper, BlackSmith:`  Structure des méthodes de ces derniers
