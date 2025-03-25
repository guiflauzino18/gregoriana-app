$(document).ready(function (){
    $('#sb-hidden').on('click', escondeSidebar)

    //Controle do Form para campos validos e invalidos
    $('.form-control').each(function() {
        if ($(this).val() == ""){
            $(this).addClass('is-invalid')
            $(this).removeClass('is-valid')
        }else {
            $(this).addClass('is-valid')
            $(this).removeClass('is-invalid')
        }

        $(this).on('keyup', confereCampo)
    })
})

function escondeSidebar(){
    $('.sidebar').toggleClass("sb-oculta");
    $('.main').toggleClass('sb-oculta-main');
    $('.sb-oculta-botao').toggleClass('ri-menu-unfold-line');
    $('.sb-oculta-botao').toggleClass('ri-menu-fold-line');
}

function confereCampo(){
    if ($(this).val() == ""){
        $(this).addClass('is-invalid')
        $(this).removeClass('is-valid')
    }else {
        $(this).addClass('is-valid')
        $(this).removeClass('is-invalid')
    }
}