let info;
var acumulador = 0;

function makeRequest() {
    const url = 'http://localhost:8989/data';
    fetch(url)
    .then(response => response.json())  
    .then(json => {
        info = json;

        //Loop para calcular o total de Co2 emitido dos veiculos
        for (var i = 0; i < info.Data[0].length; i++) {
            let floatCo2 = parseFloat(info.Data[0][i].Co2Emitido);
            acumulador += floatCo2;
        }

        //Loop para injetar na tabela de Veiculos as suas placas
        for (var i = 0; i < 10; i++) {
            var refTBody = document.getElementById('tableVeiculos');
            var newFileira = refTBody.insertRow();
            var newCell = newFileira.insertCell();
            newCell.innerHTML = info.Data[0][i].Placa

        }

        document.getElementById('totEmissCo2').innerHTML = acumulador.toFixed(2);
        document.getElementById('qtdVeiculos').innerHTML = info.Data[0].length;
    });
};

makeRequest();