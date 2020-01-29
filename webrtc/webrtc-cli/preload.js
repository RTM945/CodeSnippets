const {
    contextBridge,
    ipcRenderer
} = require('electron')

contextBridge.exposeInMainWorld(
    'electron', {
        send: (channel, data) => {
            ipcRenderer.send(channel, data)
        },
        on: (channel, func) => {
            ipcRenderer.on(channel, (event, ...args) => func(...args))
        },
        sendSync: (channel, data, callback) => {
            callback(ipcRenderer.sendSync(channel, data))
        },
    }
)




// const { dialog } = require('electron').remote;
// const fs = require("fs");
// const ipc = require('electron').ipcRenderer;

// window.dir = () => {
//     return dialog.showOpenDialogSync({ properties: ['openDirectory'] })
// }

// window.checkDir = (dir) => {
//     return fs.existsSync(dir) && fs.lstatSync(dir).isDirectory();
// }

// window.showMsg = (msg) => {
//     dialog.showMessageBoxSync({ message: msg });
// }

// window.showError = (msg) => {
//     dialog.showErrorBox("webrtc-cli", msg)
// }

// window.connected = false
// let stompClient = null;
// window.connect = (signal) => {
//     let token = makeid(6);
//     let type = 'answer';
//     let socket = new SockJS(signal);
//     stompClient = Stomp.over(socket);
//     stompClient.connect({
//         "token": token,
//         "type": type,
//     }, function (frame) {
//         window.connected = true
//         stompClient.subscribe('/user/queue/onsdp', function (msg) {
//             stompClient.send("/app/sdp", {}, "answer's sdp");
//         });
//         window.showMsg("connect success, you token: " + token)
//         document.title = document.title + ' ' + token
//     }, function (frame) {
//         window.showMsg("connect failue");
//         stompClient = null
//     });
// }

// function makeid(length) {
//     let result = '';
//     let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
//     let charactersLength = characters.length;
//     for (var i = 0; i < length; i++) {
//         result += characters.charAt(Math.floor(Math.random() * charactersLength));
//     }
//     return result;
// }

// ipc.on('app-close', _ => {
//     if (stompClient != null) {
//         stompClient.disconnect(function () {})
//     }

//     ipc.send('closed');
// });