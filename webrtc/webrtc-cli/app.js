const { app, BrowserWindow, Menu, dialog } = require('electron')
const path = require('path')
const url = require('url')
const ipc = require('electron').ipcMain
const fs = require("fs")

let win
let status = 0
app.once('ready', () => {
    win = new BrowserWindow({
        width: 400,
        height: 250,
        backgroundColor: '#D6D8DC',
        // resizable: false,
        // minimizable: false,
        // maximizable: false,
        show: false,
        webPreferences: {
            nodeIntegration: false,
            contextIsolation: true,
            enableRemoteModule: false,
            preload: path.join(__dirname, 'preload.js')
        }
    })
    Menu.setApplicationMenu(null)

    win.loadURL(url.format({
        pathname: path.join(__dirname, 'index.html'),
        protocol: 'file:',
        slashes: true
    }))
    win.webContents.openDevTools() // fot test
    win.once('ready-to-show', () => {
        win.show()
    })
    win.on('close', function (e) {
        if (status == 0) {
            if (win) {
                e.preventDefault()
                win.webContents.send('app-close')
            }
        }
    })
})

ipc.on('closed', _ => {
    status = 1
    win = null
    if (process.platform !== 'darwin') {
        app.quit()
    }
})

ipc.handle('showMsg', (_, msg) => {
    showMsg(msg)
})

function showMsg(msg) {
    dialog.showMessageBoxSync({ message: msg })
}

ipc.handle('dir', (event) => {
    let dir = dialog.showOpenDialogSync({ properties: ['openDirectory'] })
    if (dir) {
        root = dir[0]
    }
    return dir
})

ipc.handle('checkDir', (event, dir) => {
    return fs.existsSync(dir) && fs.lstatSync(dir).isDirectory()
})

let root
ipc.handle('listFiles', (event, dir) => {
    if (!dir) {
        dir = root
    }
    let list = []
    fs.readdirSync(dir).forEach(file => {
        let filePath = path.join(dir, file)
        let stat = fs.lstatSync(filePath)
        list.push({ name: file, dir: stat.isDirectory() })
    });
    return list
})

// ipc.handle('connect', async (event, signal, dir) => {
//     if (dir == null) {
//         showMsg("file root can't be empty!")
//         return undefined
//     }
//     if (!existsSync(dir)) {
//         showMsg("file root is not exist!")
//         return undefined
//     }
//     if (!lstatSync(dir).isDirectory()) {
//         showMsg("file root must be a folder!")
//         return undefined
//     }
//     root = dir
//     let token
//     try {
//         token = await connect(signal)
//     } catch (e) {
//         token = undefined
//     }
//     return token
// })


// let stompClient
// async function connect(signal) {
//     stompClient = over(new SockJS(signal))
//     let token = makeid(6)
//     return new Promise((resolve, reject) => {
//         stompClient.connect({
//             "token": token,
//             "type": 'answer',
//         }, _ => {
//             stompClient.subscribe('/user/queue/onsdp', onsdp)
//             stompClient.subscribe('/user/queue/oncandidate', oncandidate)
//             showMsg("connect success, your token: " + token)
//             resolve(token)
//         }, (error) => {
//             showMsg(error.headers.message)
//             console.log('Additional details: ' + frame.body)
//             reject(error)
//         })
//     })
// }

// function onsdp(msg) {
//     console.log(msg.body)
//     let dto = JSON.parse(msg.body)
//     let remote = dto.remote
//     let desc = dto.value
//     if (desc.type == 'offer') {
//         let pc = createPeer(remote)
//         pc.setRemoteDescription(new RTCSessionDescription(desc))
//             .then(_ => pc.createAnswer())
//             .then(answer => pc.setLocalDescription(answer))
//             .then(_ => stompClient.send("/app/sdp", {}, JSON.stringify(pc.localDescription)))
//     }
// }

// function oncandidate(msg) {
//     console.log(msg.body)
//     let dto = JSON.parse(msg.body)
//     let remote = dto.remote
//     let candidate = dto.value
//     let pc = peerMap.get(remote)
//     if (pc) {
//         pc.addIceCandidate(new RTCIceCandidate(JSON.parse(candidate)))
//     }
// }

// const peerMap = new Map()

// const servers = { iceServers: [{ "urls": ["stun:stun.l.google.com:19302"] }] }
// function createPeer(remote) {
//     let pc = new RTCPeerConnection(servers)
//     console.log(pc)
//     pc.onicecandidate = (event => event.candidate ? stompClient.send("/app/candidate", {}, JSON.stringify(event.candidate)) : console.log("Sent All Ice"))
//     pc.ondatachannel = ({ channel }) => {
//         console.log("answer data channel created!")
//         channel.onmessage = ({ data }) => handler(data)
//         channel.send(JSON.stringify({ handler: 'listFiles', data: listFiles(root) }))
//     }
//     peerMap.set(remote, pc)
//     return pc
// }

// function handle(data) {
//     console.log(data)
//     let protocol = JSON.parse(data)
//     let handle = protocol.handle
//     let value = protocol.value
//     switch (protocol.handle) {
//         case "listFilesReq":

//             break;

//         default:
//             break;
//     }
// }

// function makeid(length) {
//     let result = ''
//     let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
//     let charactersLength = characters.length
//     for (var i = 0; i < length; i++) {
//         result += characters.charAt(Math.floor(Math.random() * charactersLength))
//     }
//     return result
// }

// function checkDir(dir) {
//     return existsSync(dir) && lstatSync(dir).isDirectory()
// }

// function listFiles(dir) {
//     if (!dir) {
//         dir = root
//     }
//     let list = []
//     console.log(dir)
//     readdirSync(dir).forEach(file => {
//         let filePath = join(dir, file)
//         console.log(filePath)
//         let stat = lstatSync(filePath)
//         list.push({ name: file, dir: stat.isDirectory() })
//     });
//     return list
// }