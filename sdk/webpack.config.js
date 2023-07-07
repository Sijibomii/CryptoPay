const path = require('path');

module.exports = {
  entry: './src/index.ts',
  mode: 'production',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'dist'),
  },
  watch: true,
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
      {
        test: /\.js$/,
        use: [
          {
            loader: "babel-loader",
            options: {
              presets: [
                [
                  "env",
                  {
                    modules: "umd",
                    targets: {
                      browsers: ["last 2 Chrome versions"]
                    }
                  }
                ]
              ],
              plugins: ["add-module-exports"]
            }
          }
        ],
        exclude: /node_modules/
      },
      {
        test: /\.(jpg|png)$/,
        use: [
            {
                loader: "url-loader"
            }
        ]
      },
      {
        test: /\.css$/,
        use: [
          { loader: "style-loader" },
          { loader: "css-loader" },
          {
            loader: "postcss-loader",
            options: {
              sourceMap: true,
              plugins: [require("autoprefixer")({ grid: true })]
            }
          }
        ]
      }
    ],
  },
  resolve: {
    extensions: ['.ts', '.js'],
  },
  devtool: "source-map",
  devServer: {
    static: {
        directory: path.join(__dirname, 'public'),
    },
    host: "localhost",
    port: "5000",
    compress: true,
    // hot: true,
    // liveReload: true,
    // watchFiles: true
  }
};