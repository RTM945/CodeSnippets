$('#dirBtn').on('click', () => {
    window.electron.sendSync('dir', null, (path) => {
        if (path != undefined) {
            $('#fileRoot').val(path)
        }
    })
})

let stompClient = null;
function connect(signal) {
    let token = makeid(6);
    let type = 'answer';
    let socket = new SockJS(signal);
    stompClient = Stomp.over(socket);
    stompClient.connect({
        "token": token,
        "type": type,
    }, function (frame) {
        stompClient.subscribe('/user/queue/onsdp', function (msg) {
            stompClient.send("/app/sdp", {}, "answer's sdp");
        });
        showMsg("connect success, your token: " + token)
        document.title = document.title + ' ' + token
    }, function (frame) {
        showMsg("connect failue");
        stompClient = null
    });
}

window.electron.on('app-close', _ => {
    if(stompClient != null) {
        stompClient.disconnect(() => {
            showMsg("Bye~");
            window.electron.send('closed')
        })
    }
})

$('#connectBtn').on('click', async () => {
    let signalServer = $('#signal').val()
    signalServer = 'http://localhost:8080/signalling' //for test
    if (signalServer == '') {
        showMsg("signal server address can't be null!")
        return
    }
    let check = await checkDir($('#fileRoot').val())
    if (!check) {
        showMsg("wrong file root!")
        return
    }
    if(stompClient != null) {
        window.showMsg("already connected!")
        return
    }
    connect(signalServer)
})

function showMsg(msg) {
    window.electron.send('showMsg', msg)
}

async function checkDir(dir) {
    const result = await new Promise(resolve => {
        window.electron.sendSync('checkDir', dir, (result) => {
            resolve(result)
        })
    })
    return result
}

function makeid(length) {
    let result = '';
    let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let charactersLength = characters.length;
    for (var i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}