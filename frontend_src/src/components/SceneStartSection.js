import React from 'react';
import {Button, Caption, H3, ProgressBar} from 'hig-react'


class SceneStartSection extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            processStarted: false,
            processLabel: "Start scene processing",

        }
    }

    startProcessing = () => {
        this.setState({
            processStarted: true,
            processLabel: "Scene is being processed ..."
        });
        this.props.startProcess();
    };





    render() {

        let display;
        const {progress} = this.props;
        const {processStarted, processLabel} = this.state;

        if (processStarted) {
            display = <ProgressBar percentComplete={progress}/>
        } else {
            display = <div>
                <Caption>Everything is set:</Caption>
                <div style={{textAlign: "center"}}>
                    <Button title="Start scene processing" onClick={this.startProcessing}/>
                </div>
            </div>

        }

        return (
            <div>
                <H3>{processLabel}</H3>
                {display}
            </div>
        )
    }
}


export default SceneStartSection;