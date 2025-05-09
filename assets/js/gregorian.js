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

    //CLique do botão de Ok no Loading (recarrega a página)
    $('#btnLoadingClose').on('click', recarregaPagina)
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


//Modal de Processamento e Sucesso ou Erro.

function showLoading(){

    $('#loading').modal('show')
    $('.loading-erro').addClass('d-none')
    $('.spinner-border').removeClass('d-none')
    $('.loading-sucesso').addClass('d-none')
    $('#btnLoadingClose').addClass('d-none')
    $('#pLoadingMessage').parent().removeClass('d-none');
}

function showLoadingErro(msg){
    $('#pLoadingMessage').text(msg)
    $('.loading-erro').removeClass('d-none');
    $('.spinner-border').addClass('d-none');
    $('.loading-sucesso').addClass('d-none');
    $('#pLoadingMessage').parent().removeClass('d-none');
}

function showLoadingSucesso(msg){
    $('#pLoadingMessage').text(msg)
    $('.loading-erro').addClass('d-none')
    $('.spinner-border').addClass('d-none')
    $('.loading-sucesso').removeClass('d-none')
    $('#btnLoadingClose').removeClass('d-none')
    $('#pLoadingMessage').parent().removeClass('d-none');
}

function hideLoading(){
    $('#loading').modal('dispose');
}

function recarregaPagina(){
    location.reload();
}


//Modal de confirmação 
function showConfirma(m, funcao){
    $('#confirma-message').text(m);
    $('#modal-confirma').modal('show')
    $('#bt-confirma').on('click', funcao);
}

function hideConfirma(){
    $('#modal-confirma').modal('hide');
}


// Controle de histórico de navegação
var navegacao = []


