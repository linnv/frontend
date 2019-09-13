import React from 'react';
import axios from 'axios';
export class Timer extends React.Component {
	constructor(props) {
		super(props);
		this.state = { info: "" };
	}

	tick() {
		try {
			// axios.get("http://127.0.0.1:8083/upload").then(resp=> {
			axios.get("/upload").then(resp=> {
				console.log(resp);
				const retInstance=resp.data
				this.setState({
					info:retInstance.qps
				});
			})
		} catch (e) {
			console.log(`😱 Axios request failed: ${e}`);
		}
	}

	componentDidMount() {
		this.interval = setInterval(() => this.tick(), 3000);
	}

	componentWillUnmount() {
		clearInterval(this.interval);
	}

	render() {
		return (
			<div>
			QPS: {this.state.info}
			</div>
		);
	}

}

