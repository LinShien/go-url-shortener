// webpack.config.js
var path = require('path')

module.exports = {
  entry: ['./js/app.jsx'], // 在 index 檔案後的 .js 副檔名是可選的

  output: {
    path: path.join(__dirname, 'js'),
    filename: 'bundle.js'
  },

  module: {
      rules: [
          { test: /\.css$/, use: ['css-loader']},
          { test: /\.jsx?$/, use: 'babel-loader'}
        ]
  }
}