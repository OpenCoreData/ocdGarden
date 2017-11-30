const path = require('path');

module.exports = {
    entry: [  './webcomponents-loader.js', './custom-elements.min.js','./bower_components/slimjs/Slim.js', './tag2.js'],
    output: {
        path: path.resolve(__dirname, '.'),
        filename: 'bundle.js'
    },
    devtool: 'eval-source-maps',    
    module: {
        loaders: [
            {
                test: /\.css$/,
                loaders: ["style-loader", "css-loader", "source-map-loader"]
            }
        ],
        rules: [
            {
              test: /\.js$/,
              use: ["source-map-loader"],
              enforce: "pre"
            }
          ]
    },
};
