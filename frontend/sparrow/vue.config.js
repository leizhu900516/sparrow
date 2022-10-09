const webpack = require('webpack');
const { env } = process;
const CompressionWebpackPlugin = require('compression-webpack-plugin');
const productionGzipExtensions = ['js', 'css'];

module.exports = {
  publicPath: './',
  // 隐藏源码
  productionSourceMap: false,
  devServer: {
        // host: '0.0.0.0', // 允许外部ip访问
        port: 9080, // 端口
        https: false, // 启用https
        proxy: {
            '/api': {
                target: 'http://localhost:12345',
                changeOrigin: true,
                secure: false,
                pathRewrite: {
                    '^/api': '/api'
                }
            }
        }
    },
  chainWebpack: (config) => {
    config.performance.set('hints', false);

    if (env.NODE_ENV === 'production') {
      config.optimization.minimizer('terser').tap((options) => {
        options[0].terserOptions.compress.drop_console = true;
        return options;
      });
    }
  },
  configureWebpack: (config) => {

    // 生产环境重置HtmlWebpackPlugin插件
    if (env.NODE_ENV === 'production') {
      // nginx的配置文件里加上这句gzip_static on; # 开启 gzip 压缩
      config.plugins.push(
        new CompressionWebpackPlugin({
          algorithm: 'gzip',
          test: new RegExp(`\\.(${productionGzipExtensions.join('|')})$`),
          threshold: 10240,
          minRatio: 0.8,
          deleteOriginalAssets: false, // 删除源文件，不建议
        })
      );

      config.plugins.push(
        new webpack.optimize.LimitChunkCountPlugin({
          maxChunks: 5,
          minChunkSize: 100,
        })
      );
      config.plugins.push(
        new webpack.ProvidePlugin({
          'window.Quill': 'quill/dist/quill.js',
          'Quill': 'quill/dist/quill.js'
        }),
      )
    }
  },

};
