let info;
var acumulador = 0;

function requestVeiculos() {
    const url = 'http://localhost:5000/listaVeiculos';
    fetch(url)
    .then(response => response.json())  
    .then(json => {
        info = json;

        //Loop para calcular o total de Co2 emitido dos veiculos
        for (var i = 0; i < info.length; i++) {
            let floatCo2 = parseFloat(info[i].Co2Emitido);
            acumulador += floatCo2;
        }

        //Loop para injetar na tabela de Veiculos as suas placas
        for (var i = 0; i < 20; i++) {
            var refTBody = document.getElementById('tableVeiculos');
            var newFileira = refTBody.insertRow();
            var newCell = newFileira.insertCell();
            newCell.innerHTML = info[i].Placa
        }

        document.getElementById('totEmissCo2').innerHTML = (acumulador).toLocaleString('pt-IN');
        document.getElementById('qtdVeiculos').innerHTML = info.length;
    });
};

function requestTrajetos() {
    const url = 'http://localhost:5000/listaTrajetos';
    fetch(url)
    .then(response => response.json())  
    .then(json => {
        info2 = json;

        //Loop para injetar na tabela de Veiculos as suas placas
        for (var i = 0; i < 10; i++) {
            var refTBody = document.getElementById('tableViagens');
            var newFileira = refTBody.insertRow();
            var newCell = newFileira.insertCell();
            var newCell2 = newFileira.insertCell();
            var newCell3 = newFileira.insertCell();
            newCell.innerHTML = info2[i].IdPlaca
            newCell2.innerHTML = info2[i].IdTrajeto
            newCell3.innerHTML = "None"
        }

        document.getElementById('qtdViagens').innerHTML = info2.length
    });
};

requestVeiculos();
requestTrajetos();