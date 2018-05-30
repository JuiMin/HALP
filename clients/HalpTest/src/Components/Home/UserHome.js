// UserHome describes the home screen seen by a known user.
// This just means a user that has logged in.

// Import React Components
import React, { Component } from 'react';
import { Button, View, Text } from 'react-native';
import { StackNavigator } from 'react-navigation';
import Icon from 'react-native-vector-icons/MaterialIcons'

// Import redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

const mapStateToProps = (state) => {
	return {
            AuthToken: state.AuthReducer.authToken,
		loggedIn: state.AuthReducer.loggedIn
	}
}

// Export the default class
class UserHome extends Component {
   render() {
      return(
         <View style={Styles.home}>
            <Text>Dashboard for logged in User</Text>
            <Text>{this.props.AuthToken}</Text>
         </View>
      )
   }
}

export default connect(mapStateToProps)(UserHome)