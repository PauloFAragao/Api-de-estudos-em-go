$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) 
{
    evento.preventDefault();
    //console.log("Dentro da função usuário!")
    
    if ($('#senha').val() != $('#confirmar-senha').val()){
        alert("As senhas não coincidem!");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data:{
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {// status resposta da pagina 201, 200, 204
        alert("usuário cadastrado com sucesso!");
    }).fail(function(erro){// status resposta da pagina 400, 404, 401, 403, 500
        console.log(erro);
        alert("Erro ao cadastrar usuário!");
    });

}