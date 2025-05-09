$(document).ready(function(){

    //Adiciona listenner nos botões de acesso a cada módulo
    $('.btn-acesse').each(function() {
        $(this).on('click', acessoModulo)
    });

    //Adiciona o home na navegacao
    
})



//ClickListenner dos botoes btn-acess
function acessoModulo(){

    switch ($(this).data('modulo')) {
        case "configuracao":
            navegacao.push("configuraçao")
            window.location.href = "/configuracao"
            break;
        case "faturamento":
            navegacao.push("faturamento")
            window.location.href = "/faturamento"
            break;

        case "agendamento":
            navegacao.push("agendamento")
            window.location.href = "/agendamento"
            break
        case "atendimento":
            navegacao.push("atendimento")
            window.location.href = "/atendimento"
            break
    
        default: 
            navegacao.push("atendimento")
            window.location.href = "/atendimento"
            break;
    }

}