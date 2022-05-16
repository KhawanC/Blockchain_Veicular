let toggle = document.querySelector('.toggle');
let nav = document.querySelector('.navigation');
let main = document.querySelector('.main');

toggle.onclick = function() {
    nav.classList.toggle('ativo')
    main.classList.toggle('ativo')
}