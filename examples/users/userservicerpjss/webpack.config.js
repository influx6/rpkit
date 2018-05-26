// path ways for files.
const webConfig = require("./config/webpack.config.web");
const nodeConfig = require("./config/webpack.config.node");

let webpackConfig = null;

// based on provided environment variable,
// build library for either.
switch (process.env.pack.trim()) {
case "web":
    webpackConfig = webConfig;
    break;
case "node":
    webpackConfig = nodeConfig;
    break;
default:
    throw new Error("unknown build/pack type, please specify either 'web' or 'node' eg env PACK=web:");
}

module.exports = webpackConfig;