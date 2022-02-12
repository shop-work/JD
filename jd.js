
var demo=document.getElementById('demo')
var picture=document.getElementsByClassName('carousel-item')
var li=document.getElementsByClassName('li')
var prev=document.getElementById('prev')
var next=document.getElementById('next')
var len=picture.length
 var   index=0
  var  timer=null 
    function solide(){
        li.onmouseover=function(){
            if(timer){
                clearInterval(timer);
            }
        }
    
        li.onmouseout=function(){
            timer=setInterval(function(){
            index++;
            if(index>=len){
                index=0;
            }
            changImg();
        },3000);
        }
        li.onmouseout();
        for( var j=0;j<len;j++){
            li[j].index=j;
            li[j].onclick=function(){
                index=this.index;
                changImg();
            }
        }
        prev.onclick=function(){
            index--;
            if(index<0){
                index=len-1;
            }
            changImg();
        }
        next.onclick=function(){
            index++;
           if(index>len){
               index=0;
           }
           changImg();
        }
    }
        solide();
     function changImg(){
            for(var i=0;i<len;i++){
                picture[i].style.display='none';
           
            }
            picture[index].style.display='block';
     
        }
