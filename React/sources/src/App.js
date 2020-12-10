import React from 'react';

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            url: "",
            info: "",
        };
    }

    componentDidMount() {
        fetch('//raspberry:5000/wallpaper/1080x1920', {
            method: "GET",
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
        })
            .then(response => response.json())
            .then(data => {
                this.setState({url: data.url, info: data.info})
            })
    }

    render() {
        return (
            <div className="App">
                <header className="App-header">
                    <img src={this.state.url} className="App-logo" alt="logo"/>
                    <p>
                        {this.state.info}
                    </p>
                </header>
            </div>
        );
    }
}

export default App;
