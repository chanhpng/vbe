const { app } = require('electron');
const path = require('path');

const osShortName = function() {
    switch (process.platform) {
        case "win32":
            return "win"
        case "darwin":
            return "mac"
        case "linux":
            return "linux"
        default:
            return null
    }
}();

module.exports = {
    resourcesPath: function () {
        if (!app.isPackaged) {
            return path.join(__dirname, "..", "resources", osShortName);
        }
        return process.resourcesPath;
    },
    defaultServerBinary: function () {
        if (!app.isPackaged) {
            return {
                "mac": path.join(__dirname, "..", "..", "dist", "vbackup_darwin_amd64", "vbackup"),
                "win": path.join(__dirname, "..", "..", "dist", "vbackup_windows_amd64_v1", "vbackup.exe"),
                "linux": path.join(__dirname, "..", "..", "dist", "vbackup_linux_amd64", "vbackup"),
            }[osShortName]
        }

        return {
            "mac": path.join(process.resourcesPath, "server", "vbackup"),
            "win": path.join(process.resourcesPath, "server", "vbackup.exe"),
            "linux": path.join(process.resourcesPath, "server", "vbackup"),
        }[osShortName]
    },
    selectByOS: function (x) {
        return x[osShortName]
    },
}