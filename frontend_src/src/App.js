import React, {Component} from 'react';

import {Body, Button, GlobalNav, H3, Modal} from 'hig-react';
import Content from "./components/Content";

import 'hig-react/lib/hig-react.css';
import 'react-dropzone-component/styles/filepicker.css';
import "dropzone/dist/min/dropzone.min.css";
import './App.css';





class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            hasCredentials: true
        }
    }


    render() {

        const {hasCredentials} = this.state;
        let showContent;

        if(hasCredentials) {
            showContent = <Content/>
        } else {
            showContent = <Modal
                title="Forge Secrets Not Properly Set"
                open={!hasCredentials}
                style="alternate"
            >
                <H3>Could not find credentials</H3>
                <Body>Please refer to README.md for proper setting the Forge secrets and restart the server</Body>
            </Modal>
        }


        return (
            <GlobalNav

                topNav={
                    {
                        logo: 'https://viewer-rocks.autodesk.io/images/forge-logo.png',
                        logoLink: 'https://developer.autodesk.com/',
                        projectAccountSwitcher: {
                            projects: [
                                {
                                    label: 'Photo-II-3D',
                                    image: 'https://developer.static.autodesk.com/homepage_after_au/img/icon-reality-cap.svg',
                                    id: 'p1'
                                }
                            ]
                        }

                    }
                }
            >

                {showContent}





            </GlobalNav>


        );
    }
}

export default App;


/*

 <div className="App">
                    <header className="App-header">
                        <h1 className="App-title">This is a stub</h1>
                    </header>
                    <DropToUpload
                        onDrop={this.handleDrop}
                    >
                        <Button title="DROP ZONE"/>
                    </DropToUpload>
                </div>


                ===============
    handleDrop = (files) => {
        let data = new FormData();

        files.forEach((file, index) => {
            data.append('file[' + index + ']', file);
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

 */