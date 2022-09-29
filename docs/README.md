 # Documentation: Project-Red

*Par Marianne Corbel et Eva Chibane*
____

## A propos
Ce projet a été réalisé dans le cadre de **Ymmersion** à Ynov Aix par **Marianne Corbel** et **Eva Chibane**, respectant un cahier des charges spécifique.

___

*Ce projet a pour vocation d’être la démo d’un jeu-vidéo court de type RPG tour-par-tour. Les fonctionnalités disponibles dans la démo sont susceptibles d’évoluer avec le temps !*
___

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

`Character`: Utilisé pour le joueur.

`NPC`: Structure interne des NPC.

`ShopKeeper, BlackSmith`:  Structure des méthodes de ces derniers.

`Price`: Structure des items possédés et vendu par les NPC pouvant vendre des objets.

`Stats`: Structure des informations relatives au joueur.

`Enemy`: Utilisé pour les enemies.

`Equipment`: Methodes et définition de l'équipement du joueur (*Armures, Armes*)

`Inventory`: Methodes liées à ce que le joueur peut transporter (*Argent, Items*)

`Item`: Informations et méthodes relatives aux items 
___

## Fonctions

Fonction d'introduction. Permet au joueur de choisir son nom et sa classe.
```go
Player.CharacterCreation()
```

Fonction principale du jeu. Lance le menu à choix multiple.

```go
menu.OpenMenu()
```

Affiche toutes les informations du joueur (*stats, skills, equipement*).
```go
PrintInfo(player)
```

Affiche l'inventaire du joueur.
```go
P1.AccessInventory()
```

Lance le menu de vente du vendeur.
```go
ShopDude.BuyMenu()
```

Lance le menu du forgeron.
```go
SmithGuy.BuyMenu()
```

Lance un combat d'exemple entre le joueur et l'enemie renseigné.
```go
TrainingFight(player, enemy)
```

Lance la fonction du mini-jeu easter egg.
```go
WhoAreThey()
```

Gestion de l'utilisation et de la selection des items.
```go
item.ItemMenu(count, player, inventory, environement)
```

Gestion de la selection des items dans les shops.
```go
shop.SelectShopItem(item, count)
```

Gestion de l'achat d'un item.
```go
shop.BuyItem(item, count)
```

Detection de l'action performé par le joueur
```go
player.RegisterPlayerAction(turn, enemy)
```
Joue l'action enregistré par `RegisterPlayerAction`
```go
player.PlayerTurn(turn, player, enemy)
```

Joue le tour de l'adversaire
```go
enemy.EnemyTurn(turn, player)
```

Donne les resultat du combat
```
GetBattleResults(turn, player, enemy)
```