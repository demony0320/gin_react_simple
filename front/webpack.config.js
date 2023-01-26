const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const HtmlWebpackPlugin = require("html-webpack-plugin");
const path = require('path');
const webpack = require('webpack');

const mode = process.env.NODE_ENV || 'development';
const prod = mode === 'production';

module.exports = {
	entry:// './public/index.html'
	{
		'build/bundle': ['./src/main.js'],
		//		main : './src/main.js',
	}
	,
	resolve: {
		alias: {
			svelte: path.dirname(require.resolve('svelte/package.json'))
		},
		extensions: ['.mjs', '.js', '.svelte'],
		mainFields: ['svelte', 'browser', 'module', 'main']
	},
	output: {
		path: path.join(__dirname, '/public'),
		filename: '[name].js',
		chunkFilename: '[name].[id].js',
		publicPath: 'auto',
	},
	module: {
		rules: [
			{
				test: /\.svelte$/,
				use: {
					loader: 'svelte-loader',
					options: {

						emitCss: true,
						hotReload: true,
						hotOptions: {
							// Prevent preserving local component state
							preserveLocalState: false,
			  
							// If this string appears anywhere in your component's code, then local
							// state won't be preserved, even when noPreserveState is false
							noPreserveStateKey: false,
			  
							// Prevent doing a full reload on next HMR update after fatal error
							noReload: false,
			  
							// Try to recover after runtime errors in component init
							optimistic: false,
			  
						  }

						  
					}
				}
			},
			{
				test: /\.css$/,
				use: [
					MiniCssExtractPlugin.loader,
					'css-loader'
				]
			},
			{
				// required to prevent errors from Svelte on Webpack 5+
				test: /node_modules\/svelte\/.*\.mjs$/,
				resolve: {
					fullySpecified: false
				}
			}
		]
	},
	mode,
	plugins: [

		new webpack.HotModuleReplacementPlugin(),
		new MiniCssExtractPlugin({
			filename: '[name].css'
		}),
		//new HtmlWebpackPlugin(),
	],
	devtool: prod ? false : 'source-map',
	devServer: {
		//static: 'public',
		contentBase: 'public',
		host: '0.0.0.0',
		//devMiddleware: { publicPath: '/public' },
  		//static: { directory: path.resolve(__dirname) },
		hot: true,
		overlay: true,
		publicPath: '/',
		historyApiFallback: true


	}
};
