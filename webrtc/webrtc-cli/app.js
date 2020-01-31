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

ipc.on('showMsg', (_, msg) => {
    dialog.showMessageBoxSync({ message: msg })
})

ipc.on('dir', (event) => {
    event.returnValue = dialog.showOpenDialogSync({ properties: ['openDirectory'] })
})

ipc.on('checkDir', (event, dir) => {
    event.returnValue = fs.existsSync(dir) && fs.lstatSync(dir).isDirectory()
})