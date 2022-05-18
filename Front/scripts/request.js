const loading_img = document.getElementById('load')
const erro_img = document.getElementById('erro')
const conectado = document.getElementById('conectado')

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function makeRequest() {
    const url = 'http://localhost:5000/rotaPadrao'
    fetch(url).then(
        response => {
            console.log(response.json())
            if (response.status == 200) {
                console.log(response.status)
                loading_img.style.visibility = "collapse"
                conectado.style.visibility = "visible"

                sleep(2500).then( () => {
                    document.location.href = './pages/resumo.html'
                });
                
            }
            else{
                console.log(response.status)
                loading_img.style.visibility = "collapse"
                erro_img.style.visibility = "visible"
            }
        }
    ).catch(err => {
        console.log("erro")
        loading_img.style.visibility = "collapse"
        erro_img.style.visibility = "visible"
        });
}

sleep(2000).then(() => {
    makeRequest();
});