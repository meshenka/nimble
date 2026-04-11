const path = require('path');

module.exports = {
  entry: './frontend/src/index.tsx',
  mode: 'production',
  resolveLoader: {
    modules: [path.resolve(__dirname, 'node_modules')]
  },
  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: [
              '@babel/preset-env',
              '@babel/preset-react',
              '@babel/preset-typescript'
            ]
          }
        },
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    modules: [path.resolve(__dirname, 'node_modules')],
    extensions: ['.tsx', '.ts', '.js'],
    fallback: {
      // "path": require.resolve("path-browserify"),
      // "os": require.resolve("os-browserify/browser"),
      "fs": false,
      // Add other Node.js core modules that need polyfills here
    }
  },
  output: {
    filename: '[name].bundle.js',
    path: path.resolve(__dirname, 'public'),
    clean: {
      keep: /^(index\.html|styles\.css)$/,
    },
  },
  optimization: {
    moduleIds: 'deterministic',
    runtimeChunk: 'single',
    splitChunks: {
      cacheGroups: {
        vendor: {
          test: /[\\/]node_modules[\\/]/,
          name: 'vendors',
          chunks: 'all',
        },
      },
    },
  },
};
