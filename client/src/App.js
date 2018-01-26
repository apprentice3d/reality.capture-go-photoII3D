import React, {Component} from 'react';
import './App.css';
import DropToUpload from 'react-drop-to-upload'

class App extends Component {

    handleDrop = (files) => {
        let data = new FormData();

        files.forEach((file, index) => {
            data.append('file['+index + ']', file);
        });

        fetch('http://localhost/upload', {
            method: 'POST',
            body: data
        }).then((data) => {
            console.log(data.text())
        }).catch((problem) => {
            console.log(problem)
        })
    };

    render() {
        return (
            <div className="App">
                <header className="App-header">
                    <h1 className="App-title">This is a stub</h1>
                </header>
                <p className="App-intro">
                    To get started, edit <code>src/App.js</code> and save to reload.
                </p>
                <DropToUpload
                    onDrop={this.handleDrop}
                >
                    drop files here!
                </DropToUpload>
            </div>
        );
    }
}

export default App;
