$('#dirBtn').on('click', () => {
    window.electron.sendSync('dir', null, (path) => {
        if (path != undefined) {
            $('#fileRoot').val(path)
        }
    })
})

let stompClient = null
function connect(signal) {
    let token = makeid(6)
    let type = 'answer'
    let socket = new SockJS(signal)
    stompClient = Stomp.over(socket)
    stompClient.connect({
        "token": token,
        "type": type,
    }, _ => {
        $('#fileRoot').prop('readonly', true)
        $('#dirBtn').off()
        let pc = createPeer(stompClient)
        stompClient.subscribe('/user/queue/onsdp', function (msg) {
            console.log(msg.body);
            let desc = JSON.parse(msg.body)
            if (desc.type == 'offer') {
                pc.setRemoteDescription(new RTCSessionDescription(desc))
                    .then(_ => pc.createAnswer())
                    .then(answer => pc.setLocalDescription(answer))
                    .then(_ => stompClient.send("/app/sdp", {}, JSON.stringify(pc.localDescription)));
            }
        })

        stompClient.subscribe('/user/queue/oncandidate', function (msg) {
            if (msg.body) {
                console.log(msg.body);
                pc.addIceCandidate(new RTCIceCandidate(JSON.parse(msg.body)));
            }
        })

        showMsg("connect success, your token: " + token)
        document.title = document.title + ' ' + token
    }, _ => {
        showMsg("connect failue")
        stompClient = null
    })
}

const servers = { iceServers: [{ "urls": ["stun:stun.l.google.com:19302"] }] }
function createPeer(stompClient) {
    let pc = new RTCPeerConnection(servers)
    console.log(pc)
    pc.onicecandidate = (event => event.candidate ? stompClient.send("/app/candidate", {}, JSON.stringify(event.candidate)) : console.log("Sent All Ice"))
    pc.ondatachannel = ({ channel }) => {
        console.log("answer data channel created!")
        channel.onmessage = ({ data }) => handler(data)
        window.electron.sendSync('listFiles', null, (data) => {
            channel.send(JSON.stringify({ handler: 'listFiles', data: data }))
        })
    }
    return pc
}

function handler(data) {

}

window.electron.on('app-close', _ => {
    if (stompClient != null) {
        stompClient.disconnect(_ => {
            showMsg("Bye~")
        })
    }
    window.electron.send('closed')
})

$('#connectBtn').on('click', async _ => {
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
    if (stompClient != null) {
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
    let result = ''
    let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
    let charactersLength = characters.length
    for (var i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength))
    }
    return result
}