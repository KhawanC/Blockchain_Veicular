let list = document.querySelectorAll('.container .navigation li');

function linkAtivo() {
    list.forEach((item) =>
        item.classList.remove('hovered'));
        this.classList.add('hovered');
}

list.forEach((item) =>
    item.addEventListener('mouseover',linkAtivo)
)