$(document).ready(function (){
    $('#sb-hidden').on('click', escondeSidebar)
})

function escondeSidebar(){
    $('.sidebar').toggleClass("sb-oculta");
    $('.main').toggleClass('sb-oculta-main');
    $('.sb-oculta-botao').toggleClass('ri-menu-unfold-line');
    $('.sb-oculta-botao').toggleClass('ri-menu-fold-line');
}