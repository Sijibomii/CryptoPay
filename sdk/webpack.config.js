const path = require('path');
const webpack = require("webpack");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

const libraryName = "cryptopay-sdk";
module.exports = {
  entry: path.join(__dirname, "src"),
  output: {
    path: path.join(__dirname, "lib"),
    // library: libraryName,
    // filename: libraryName + ".js",
    // libraryTarget: "umd",
    // umdNamedDefine: true,
    // publicPath: "/"
  },
  watch: true,
  optimization: {
    splitChunks: {
        cacheGroups: {
            styles: {
                name: "styles",
                test: /\.css$/,
                chunks: "all",
                enforce: true
            }
        }
    },
  },
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
  },
  plugins: [
    new webpack.HotModuleReplacementPlugin(),

    new MiniCssExtractPlugin({
        filename: libraryName + ".css",
        chunkFilename: libraryName + ".[id].css"
    }),

  ]
};