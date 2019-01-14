addEventListener("message", function (event){
	var xhr = new XMLHttpRequest();
	xhr.open("GET","http://127.0.0.1:9999");
	xhr.onload = function(){
		postMessage(xhr.responseText);
	};
	
	xhr.send();
},false);