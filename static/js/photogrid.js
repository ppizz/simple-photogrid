
/* ==================================================================
PHOTOGRID.JS 
ppizz 2017 V0.1 photogrid
Gere l'environnement d'affichage des photos
Communique en POST JSON avec le server WEB gallery.go

 ====================================================================*/

$(function() {
	var $getphoto = $("#getphoto");
	var $mosaic = $("#mosaic"); // mosaic de photo: div photo créé dynamiquement
     
	 // les données internes de photogrid.js
	 var _tabPhoto = [];
	 /*
			   Id: index unique de la photo dans la table Photo	
         Name: nom de fichier
         Title: titre - description de la photo
			   Note: note de 1 à 5 
         Orientation: 1=paysage, 8=portrait
     */
	
// ----------- demande les ressources photo
function Get_photo(dir) {
		var reqAjax = {NomDir: dir};
		var url;
		var txtPhoto;
		var starNote;
		var elem;

var jqxhr =
    $.ajax({
        url: '/getphoto',
		type: 'POST',
		data : JSON.stringify(reqAjax),
        contentType: 'application/json; charset=utf-8',
		dataType: 'json',
        async: true,		
    })
    .done (function(data) {  
		_tabPhoto = [];
		_ModifPhoto = false;
		_idShow = 0;
		$mosaic.empty();
		//on met à jour le tableau de photo
		for (var i in data) {
		    _tabPhoto.push({
			   Id: data[i].Id,	
			   Name: data[i].Name,
			   Date: data[i].Date,
			   DirName: data[i].DirName,
			   Album: data[i].Album,
         Title: data[i].Title,
			   Note: data[i].Note,
         Orientation: data[i].Orientation,
			}); 	
		  // on cree dans Mosaic le DIV de la photo avec ses textes
      starNote = GetStarNote(_tabPhoto[i].Note);
		  txtPhoto = data[i].Name + " - " + data[i].Date + "<div class='star'>" + starNote + "</div><br />";
		  txtPhoto = txtPhoto + data[i].Album + " : " + data[i].Title;
		  //console.log(txtPhoto);
		  url = "./static/PHOTO/" + data[i].DirName + "/" + data[i].Name;
		  if (data[i].Orientation !== "8") {
			elem = $("<div class='photo' id=" + i + "><img src='" + url +
			".jpg' alt='' /><div class='divcaption'>" + txtPhoto + "</div></div>");
		  } else {
			elem = $("<div class='photo portrait' id=" + i + "><img src='" + url +
			".jpg' alt='' /><div class='divcaption portrait'>" + txtPhoto + "</div></div>");
		  }
		  $mosaic.append(elem);
		}
	})
		.fail (function(jqxhr, textStatus, errorThrown)  {
			console.log("Error: " + textStatus + " : " + errorThrown) ;
		});	
	}
// ----------- renvoie la note sous forme d 'etoile 
function GetStarNote(note) {
			var starNote = "";
			if (note==0) {
				starNote ="-";
			} else if (note==1){
				starNote = "*";
			  } else if (note==2){
				starNote = "**";
			  } else if (note==3){
				starNote = "***";
				} else if (note==4){
				starNote = "****";
				} else if (note==5){
				starNote = "*****";
				}
	return starNote;		  
  }

//===============================================================================
// gestion des evenements

$getphoto.click(function() {
   console.log('Get Photos()');
   Get_photo();
});

// =================================================================================
// init javascript


});
