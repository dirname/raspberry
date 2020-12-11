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
            <div style={{backgroundColor: 'black'}}>
                <p style={{color: 'white', margin: 0}}>Battery: 10%</p>
                <div style={{
                    backgroundImage: `url(${this.state.image.url})`,
                    backgroundSize: 'cover',
                    height: '100%',
                    width: '100%',
                    position: 'absolute',
                    backgroundPosition: 'center',
                    backgroundRepeat: 'no-repeat',
                    paddingTop: '30px',
                    display: 'flex',
                    flexDirection: 'column',
                    justifyContent: 'space-around',
                    padding: '0 10vw',
                }}>
                    <p style={{color: "white"}}>{this.state.info}</p>
                </div>
            </div>

            // <div className="App" style={{backgroundImage: `url(${this.state.url})`}}>
            //     <header className="App-header">
            //         {/*<img src={this.state.url} className="App-logo" alt="logo"/>*/}
            //         <p>
            //             {this.state.info}
            //         </p>
            //     </header>
            // </div>
        );
    };
}

export default App;
