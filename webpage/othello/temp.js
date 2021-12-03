function start(){
	var rankButton = document.getElementById("rank");
	var gameButton = document.getElementById("game");
	var aboutGameButton = document.getElementById("aboutGame");
	var showBoard = document.getElementById("rankboard");
	var iframe = document.getElementById('iframe');
	var loadId = document.getElementById("loading");
	var collapse1 = document.getElementById("collapse1");
	var collapse2 = document.getElementById("collapse2");
	var collapse3 = document.getElementById("collapse3");

	var rankMusic = new Audio("rankMusic.mp3");
	var clickMusic = new Audio("clickMusic.mp3");
	
	rankButton.addEventListener("click",showRank,false);
	gameButton.addEventListener("click",showGame,false);
	aboutGameButton.addEventListener("click",showAboutGame,false);
	
	
	getRank();
	var gameBackGroundMusic = document.getElementById("Test_Audio");
	
	//background music;
	document.addEventListener("click", function () {
	gameBackGroundMusic.muted=false;
    gameBackGroundMusic.play();
	});
	//other music
	collapse1.addEventListener("click",function(){clickMusic.play()},false);
	collapse2.addEventListener("click",function(){clickMusic.play()},false);
	collapse3.addEventListener("click",function(){clickMusic.play()},false);
	

	
	load(loadId,500);
}
arr= [];

function getRank(){
	arr= [ {"rank":1,"name":"Dino","score":64}, 
		   {"rank":2,"name":"Jim","score":54} ,
		   {"rank":3,"name":"Dino","score":64}, 
		   {"rank":4,"name":"Jim","score":54} ,
		   {"rank":5,"name":"Dino","score":64},
		   {"rank":5,"name":"Dino","score":64},
		   {"rank":5,"name":"Dino","score":64},
		   {"rank":5,"name":"Dino","score":64},
		   {"rank":5,"name":"Dino","score":64},
		   {"rank":5,"name":"Dino","score":64},
		   {"rank":5,"name":"Dino","score":64},
		   {"rank":5,"name":"Dino","score":64}
		   ];

}
function showAboutGame(){
	
	var showAboutGameSen = document.getElementById("aboutGame");
	var sen = "";
	sen+="<h1 style='font-size:60px;text-align:center;'>關於作者</h1><br><br>";
	sen+="<h3> 沈彥昭 : 遊戲製作</h3><br>";
	sen+="<h3> 李佳勳 : 前端製作</h3><br><br><br>";
	sen+="<h3> 遊戲設計介紹:........................</h3>";
	
	
}
function showGame(){
	//const gameBackGroundMusic = new Audio("gameBackGroundMusic.mp3");
	//var gameBackGroundMusic = document.getElementById("Test_Audio");
	//gameBackGroundMusic.muted = false;
	
	
	//gameBackGroundMusic.play();
	var showBoard = document.getElementById("rankboard");
	iframe.setAttribute("style","display:block;width:72vw;height:40.5vw; border: none;box-sizing: border-box;border-radius: 7px;");
	showBoard.setAttribute("style","display:none;");
}
function showRank(){
	
	
	var showBoard = document.getElementById("rankboard");
	iframe.setAttribute("style","display:none;");
	showBoard.setAttribute("style","display:block;");
	var temp ="";
	temp+="<table class='tableShow'>";
	temp+="<thead><tr><th>Rank</th><th>Name</th><th>Score</th></thead>";
	temp+="<tbody>";
	for(var i=0;i<arr.length;i++){
		temp+=("<tr class='"+"success"+"'>");
		temp+="<td>"+arr[i].rank+"</td>";
		temp+="<td>"+arr[i].name+"</td>";
        
        
        temp+="<td>"+arr[i].score+"</td>"
      temp+="</tr>"
	
	}
	temp+="</tbody>";
	console.log(temp);
	
	showBoard.innerHTML = temp;
	
}

function fadeIn(el, duration) {

    /*
     * @param el - The element to be faded out.
     * @param duration - Animation duration in milliseconds.
     */

    var step = 10 / duration,
        opacity = 0;
    function next() {
        if (opacity >= 1) { return; }
        el.style.opacity = ( opacity += step );
        setTimeout(next, 10);
    }
    next();
}
arr2 = ["loading","loading.","loading..","loading..."];
var timeForLoading =0;
var prepare = 0;

function load(loadId,duration){
	
	
	function next2(){
		if(prepare++ > 20 ){
			console.log("loading finish");
			var con = document.getElementById("container");
			con.setAttribute("style","display:none");
			
			
			fadeIn(iframe, 1000);
			
			console.log(iframe.style.height);
			showGame();
			
		
			return;
		}
		
		loadId.innerHTML = arr2[(timeForLoading++)%4];
		setTimeout(next2,200);
	
	}
	next2();

}
/*
window.onload = function(){
             setInterval("toggleSound()",100);
        }

    function toggleSound() {
                var music = document.getElementById("Test_Audio");//獲取ID  
                    
                if (music.paused) { //判讀是否播放  
                    music.paused=false;
                    music.play(); //沒有就播放 
                }    
        }
*/

window.addEventListener("load",start,false);


