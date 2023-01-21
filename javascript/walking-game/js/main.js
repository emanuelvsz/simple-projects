function restartGame(){
    var positionX = Math.floor(Math.random() * (275));
    var positionY = 0;

    desenhar.clearRect(0, 0, 500, 500)
    desenhar.fillRect(positionX, positionY, 15, 7)
}