$(document).ready(function(){

    //Adiciona listenner nos botões de acesso a cada módulo
    $('.btn-acesse').each(function() {
        $(this).on('click', acessoModulo)
    });
})



//ClickListenner dos botoes btn-acess
function acessoModulo(){

    switch ($(this).data('modulo')) {
        case "configuracao":
            window.location.href = "/configuracao"
            break;
        case "faturamento":
            window.location.href = "/faturamento"
            break;

        case "agendamento":
            window.location.href = "/agendamento"
            break
        case "atendimento":
            window.location.href = "/atendimento"
            break
    
        default: 
            window.location.href = "/atendimento"
            break;
    }

}