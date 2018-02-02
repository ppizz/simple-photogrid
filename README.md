# simple-photogrid
Une gallerie de photo simple pour demonstration  

Cette application montre comment implementer un server WEB avec GO 

Le serveur WEB  accede à un dossier JPEG et fournit les ressources à afficher à la page HTML du navigateur
Cette WEBAPP utilise une base SQLITE pour stocker et gerer un CATALOGUE de photo. Les données sont échangées au format JSON
la requete POST est servie par l'application GO qui renvoie les données JSON au front end javascript (avec l'aide de jquery). La gestion back end de la base est encapsulée dans un PACKAGE GO nommé CATALOG

Cette application est fournie à titre d'exemple et d'initiation sur la mise en oeuvre de standard WEB (JAVASCRIPT, HTML, CSS) et de languages back end (GO, SQL)

Le rendu avec firefox:

![photogrid](https://github.com/ppizz/simple-photogrid/blob/master/EcranDemo.png)

