/**
 * Malabar Mobile Client
 * @flow
 */

import React, { Component } from 'react';
import {
	AppRegistry,
	StyleSheet,
	Text,
	ListView,
	View
} from 'react-native';
import 'grpc';

class mobile extends Component {
	// Initialize the hardcoded data
	constructor(props) {
		super(props);
		const ds = new ListView.DataSource({rowHasChanged: (r1, r2) => r1 !== r2});
		this.state = {
			dataSource: ds.cloneWithRows([
				'John', 'Joel', 'James', 'Jimmy', 'Jackson', 'Jillian', 'Julie', 'Devin'
			])
		};
	}
	render() {
		let data = [
			{id: 1, author: "Pete Hunt", text: "This is one comment"},
			{id: 2, author: "Jordan Walke", text: "This is *another* comment"}
		];
		let test = "hej";

		return (
			<View style={styles.container}>
				<Text style={styles.welcome}>
				Welcome to Malabar Mobile Client!
				</Text>
				<Text style={styles.instructions}>
				This data just in from the other side:
				</Text>
				<ListView
					dataSource={this.state.dataSource}
					renderRow={(rowData) => <Text style={styles.list_item}>{rowData}</Text>}
				/>
				</View>
		);
	}
}

const styles = StyleSheet.create({
	container: {
		flex: 1,
		justifyContent: 'center',
		alignItems: 'center',
		backgroundColor: '#F5FCFF',
	},
	welcome: {
		fontSize: 20,
		textAlign: 'center',
		margin: 10,
	},
	instructions: {
		textAlign: 'center',
		color: '#333333',
		marginBottom: 5,
	},
	list_item: {
		color: '#ff00ff'
	}
});

AppRegistry.registerComponent('mobile', () => mobile);
