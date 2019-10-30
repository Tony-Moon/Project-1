function t3load() {
    var a1 = document.getElementById("a1");
    var a2 = document.getElementById("a2");
    var a3 = document.getElementById("a3");
    var b1 = document.getElementById("b1");
    var b2 = document.getElementById("b2");
    var b3 = document.getElementById("b3");
    var c1 = document.getElementById("c1");
    var c2 = document.getElementById("c2");
    var c3 = document.getElementById("c3");
    var resetGame = document.getElementById("r1")
    var turn = 0;

    a1.addEventListenter('click', function() {
        document.getElemendById("a1").disabled = true; 
        alert("clicked");
    });
    a2.addEventListenter('click', function() {
        document.getElemendById("a2").disabled = true; });
    a3.addEventListenter('click', function() {
        document.getElemendById("a3").disabled = true; });

    b1.addEventListenter('click', function() {
        document.getElemendById("b1").disabled = true; });
    b2.addEventListenter('click', function() {
        document.getElemendById("b2").disabled = true; });
    b3.addEventListenter('click', function() {
        document.getElemendById("b3").disabled = true; });

    c1.addEventListenter('click', function() {
        document.getElemendById("c1").disabled = true; });
    c2.addEventListenter('click', function() {
        document.getElemendById("c2").disabled = true; });
    c3.addEventListenter('click', function() {
        document.getElemendById("c3").disabled = true; });

    resetGame.addEventListener('click', function() {
        alert("Game was reset.");
        document.getElemendById("a1").disabled = false;
        document.getElemendById("a2").disabled = false;
        document.getElemendById("a3").disabled = false;
        document.getElemendById("b1").disabled = false;
        document.getElemendById("b2").disabled = false;
        document.getElemendById("b3").disabled = false;
        document.getElemendById("c1").disabled = false;
        document.getElemendById("c2").disabled = false;
        document.getElemendById("c3").disabled = false;
    });
}