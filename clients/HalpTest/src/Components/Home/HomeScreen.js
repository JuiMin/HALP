import React, { Component } from 'react';

// Import stylesheet and thematic settings
import Styles from '../../Styles/Styles';
import Theme from '../../Styles/Theme';

import { API_URL } from '../../Constants/Constants';
import { Picker } from 'react-native'

// Import react-redux connect 
import { connect } from "react-redux";

// Import the different views based on user state
import LargePost from '../Posts/LargePost';


import { 
   Container,
   Header,
   Body,
   Title,
   Right,
   Button,
	Content,
	CardItem,
	Text,
	Form,
	H3,
   Card,
   Icon
} from 'native-base';

import { addPosts } from '../../Redux/Actions';

const mapStateToProps = state => {
   return {
		posts: state.PostReducer.posts,
		activePost: state.PostReducer.activePost
   };
};

const mapDispatchToProps = dispatch => {
	return {
		addPosts: (posts) => {dispatch(addPosts(posts))}
	}
}

class HomeScreen extends Component {
   constructor(props) {
      super(props)
      this.state = {
         pickerIndex: 0,
			maxPosts: 20
      }
      this.onValueChange = this.onValueChange.bind(this)
		this.increaseMaxPosts = this.increaseMaxPosts.bind(this)
		
		// Gettin posts
		fetch(API_URL + "posts/get/recent?n=" + this.state.maxPosts, {
			method: "GET",
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json',
			},
		}).then(response => {
			if (response.status == 202) {
				return response.json()
			} else {
				return null
			}
		}).then(json => {
			if (json != null) {
				this.props.addPosts(json)
				// this.setState(this.state)
			}
		}).catch(err => {
			console.log(err)
		})
   }

   onValueChange(value) {
      this.setState({
        pickerIndex: value
      });
    }

	increaseMaxPosts() {
		this.setState({
			pickerIndex: this.state.pickerIndex,
			maxPosts: this.state.maxPosts + 20
		})

		// Gettin posts
		fetch(API_URL + "posts/get/recent?n=" + this.state.maxPosts, {
			method: "GET",
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json',
			},
		}).then(response => {
			if (response.status == 202) {
				return response.json()
			} else {
				return null
			}
		}).then(json => {
			if (json != null) {
				this.props.addPosts(json)
			}
		}).catch(err => {
			console.log(err)
		})
	 }

   // Here we should run initialization scripts
   render() {
      // This will be the same any user
      return (
         // <GuestHome {...this.props} />
         <Container>
            <Header style={{
               backgroundColor: Theme.colors.primaryBackgroundColor
            }}>
               <Body>
                  <Title style={{
                     color: Theme.colors.primaryColor,
                     alignSelf: "center",
                     fontWeight: "bold"
                  }}>HALP</Title>
               </Body>
            </Header>

            <Content style={{
					padding: "2%"
				}}>
					<Picker
						iosHeader="Select one"
						mode="dropdown"
						selectedValue={this.state.pickerIndex}
						onValueChange={this.onValueChange.bind(this)}
					>
						<Picker.Item label="Most Recent" value={0} />
						<Picker.Item label="More Filters Coming Soon" value={1} />
					</Picker>
              <Content>
                	{
                    	this.props.posts.map((item, i) => {
								return <LargePost key={i} post={item} {...this.props}/>
							})
               	}
              	</Content>
					<Button rounded style={Styles.button} onPress={this.increaseMaxPosts}>
						<Text>Get More Posts</Text>
					</Button>
            </Content>
         </Container>
      );
   }
}

export default connect(mapStateToProps, mapDispatchToProps)(HomeScreen)