// Import default react components
import React, { Component } from 'react';
import { StackNavigator } from 'react-navigation';

// Import Halp Components
import HomeScreen from '../Home/HomeScreen';
import Post from '../Posts/Post'
import CanvasTest from '../Canvas/CanvasTest';
import FullscreenImage from '../Posts/FullscreenImage';
import Comments from '../Comments/Comments';
// Generate a stack for navigation
// Generally, this is the component that wraps the child components
// Specifically for this file, App.js will use this as a component because it allows for
// navigating between the Compoents listed

export default StackNavigator(
	{
		Home: {
			screen: HomeScreen,
			navigationOptions: {
				header: null,
				title: "HALP"
			}
		},
		Post: {
			screen: Post,
		},
		Image: {
			screen: FullscreenImage,
      },
      Comment: {
         screen: Comments,
         navigationOptions: {
				header: null,
				title: "Reply"
			}
      },
		Canvas: {
			screen: CanvasTest,
		},
	},
	{
		initialRouteName: 'Home',
		headerMode: 'screen',
	}
)
