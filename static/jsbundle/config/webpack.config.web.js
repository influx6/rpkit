// const webpack = require("webpack");
const path = require("path");
const pre = require("./webpack.preconfig");

// path ways for files.
const rootdir = path.resolve(__dirname);
const srcdir = path.join(rootdir, "../src");
const destdir = path.join(rootdir, "../dest/web");

// export new configuration for web.
module.exports = pre.Build({
    output: {
        path: destdir,
        filename: "index.js",
    },
    module: {
        rules: [{
            test: /\.js?$/,
            include: srcdir,
            exclude: /node_modules/,
            use: [{
                loader: "babel-loader",
                options: {
                    presets: ["env", "stage-2"],
                },
            }],
        }],
    },
});