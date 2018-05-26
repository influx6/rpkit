// const webpack = require("webpack");
const path = require("path");
const lodash = require("lodash");

// path ways for files.
const rootdir = path.resolve(__dirname);
const srcdir = path.join(rootdir, "../src");

function maker(moreConfig) {
    if (typeof moreConfig !== "object") {
        throw new Error("only javascript object allowed");
    }

    return lodash.merge({
        output: {},
        module: {},
        plugins: [],
        mode: process.env.mode.trim(),
        devtool: "sourcemap",
        entry: path.join(srcdir, "app.js"),
    }, moreConfig);
}

module.exports = {
    Build: maker,
};